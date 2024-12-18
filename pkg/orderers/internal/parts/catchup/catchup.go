package catchup

import (
	"bytes"

	"github.com/matejpavlovic/mir/pkg/dsl"
	"github.com/matejpavlovic/mir/pkg/logging"
	common2 "github.com/matejpavlovic/mir/pkg/orderers/common"
	"github.com/matejpavlovic/mir/pkg/orderers/internal/common"
	hasherpbdsl "github.com/matejpavlovic/mir/pkg/pb/hasherpb/dsl"
	isspbdsl "github.com/matejpavlovic/mir/pkg/pb/isspb/dsl"
	pbftpbdsl "github.com/matejpavlovic/mir/pkg/pb/pbftpb/dsl"
	pbftpbmsgs "github.com/matejpavlovic/mir/pkg/pb/pbftpb/msgs"
	pbftpbtypes "github.com/matejpavlovic/mir/pkg/pb/pbftpb/types"
	transportpbdsl "github.com/matejpavlovic/mir/pkg/pb/transportpb/dsl"
	transportpbevents "github.com/matejpavlovic/mir/pkg/pb/transportpb/events"
	tt "github.com/matejpavlovic/mir/pkg/trantor/types"
	"github.com/matejpavlovic/mir/pkg/util/maputil"
	"github.com/matejpavlovic/mir/pkg/util/sliceutil"
	"github.com/matejpavlovic/mir/stdevents"
	stddsl "github.com/matejpavlovic/mir/stdevents/dsl"
	"github.com/matejpavlovic/mir/stdtypes"
)

func IncludeSegmentCheckpoint(
	m dsl.Module,
	state *common.State,
	params *common.ModuleParams,
	moduleConfig common2.ModuleConfig,
	logger logging.Logger,
) {

	// UponResultOne with a CatchUpResponse as context,
	// process the missing committed certificate as its hash became available.
	// This is the final step of catching up with an instance-level checkpoint.
	hasherpbdsl.UponResultOne(m, func(digest []byte, context *pbftpbtypes.CatchUpResponse) error {
		// Convenience variables
		preprepare := context.Resp
		sn := preprepare.Sn
		slot := state.Slots[state.View][sn]

		// Ignore preprepare if slot is already committed.
		if state.Slots[state.View][preprepare.Sn].Committed {
			return nil
		}

		// Check whether the received certificate was actually requested (a faulty node might have sent it on its own).
		digests := state.SegmentCheckpoint.Digests()
		if digests == nil {
			logger.Log(logging.LevelWarn, "Ignoring unsolicited CatchUpResponse.", "sn", sn)
			return nil
		}

		// Check whether the digest of the received message matches the requested one.
		if !bytes.Equal(digests[sn], digest) {
			logger.Log(logging.LevelWarn, "Ignoring CatchUpResponse with invalid digest.", "sn", sn)
			return nil
		}

		logger.Log(logging.LevelDebug, "Catching up with segment-level checkpoint.", "sn", sn)

		// Add the missing certificate, updating the corresponding Preprepare's view.
		// Note that copying a Preprepare with an updated view preserves its hash.
		preprepareCopy := *preprepare
		preprepareCopy.View = state.View
		slot.CatchUp(&preprepareCopy, digest)

		// If all certificates have been committed (i.e. this is the last certificate to commit),
		// send a Done message to all other nodes.
		// This is required for liveness, see comments for PbftSegmentChkp.
		if state.AllCommitted() {
			state.SegmentCheckpoint.SetDone()
			SendDoneMessages(m, state, params, moduleConfig, logger)
		}

		// Deliver certificate.
		isspbdsl.SBDeliver(
			m,
			moduleConfig.Ord,
			sn,
			slot.Preprepare.Data,
			slot.Preprepare.Aborted,
			state.Segment.Leader,
			moduleConfig.Self,
		)
		return nil
	})

	pbftpbdsl.UponDoneReceived(m, func(from stdtypes.NodeID, digests [][]byte) error {
		if !sliceutil.Contains(params.Config.Membership, from) {
			logger.Log(logging.LevelWarn, "sender %s is not a member.\n", from)
			return nil
		}
		applyMsgDone(m, state, params, moduleConfig, digests, from)
		return nil
	})

	pbftpbdsl.UponCatchUpRequestReceived(m, func(from stdtypes.NodeID, digest []uint8, sn tt.SeqNr) error {
		if !sliceutil.Contains(params.Config.Membership, from) {
			logger.Log(logging.LevelWarn, "sender %s is not a member.\n", from)
			return nil
		}

		if digest == nil {
			// Not necessary as this is checked where relevant later on, but an optimization
			return nil
		}

		applyMsgCatchUpRequest(m, state, moduleConfig, digest, sn, from)
		return nil
	})

	pbftpbdsl.UponCatchUpResponseReceived(m, func(from stdtypes.NodeID, resp *pbftpbtypes.Preprepare) error {
		if !sliceutil.Contains(params.Config.Membership, from) {
			logger.Log(logging.LevelWarn, "sender %s is not a member.\n", from)
			return nil
		}
		applyMsgCatchUpResponse(m, state, moduleConfig, resp, from)
		return nil
	})

}

// applyMsgDone applies a received Done message.
// Once enough Done messages have been applied, makes the protocol
// - stop participating in view changes and
// - set up a timer for fetching missing certificates.
func applyMsgDone(
	m dsl.Module,
	state *common.State,
	params *common.ModuleParams,
	moduleConfig common2.ModuleConfig,
	doneDigests [][]byte, from stdtypes.NodeID) {

	// Register Done message.
	state.SegmentCheckpoint.NodeDone(from, doneDigests, state.Segment)

	// If more Done messages still need to be received or retransmission has already been requested, do nothing.
	doneNodes := state.SegmentCheckpoint.DoneNodes()
	if doneNodes == nil || state.SegmentCheckpoint.CatchingUp {
		return
	}

	// If this was the last Done message required for a quorum,
	// set up a timer to ask for the missing committed certificates.
	// In case the requests get lost, they need to be periodically retransmitted.
	// Also, we still want to give the node some time to deliver the requests naturally before trying to catch up.
	// Thus, we pack a TimerRepeat OrdererEvent (that triggers the first "repetition" immediately) inside a TimerDelay.
	// We also set the catchingUp flag to prevent this code from executing more than once per PBFT instance.
	state.SegmentCheckpoint.CatchingUp = true

	stddsl.TimerDelay(
		m,
		moduleConfig.Timer,
		params.Config.CatchUpDelay,
		stdevents.NewTimerRepeat(
			moduleConfig.Timer,
			params.Config.CatchUpDelay,
			stdtypes.RetentionIndex(params.Config.EpochNr),
			catchUpRequests(state, moduleConfig, doneNodes, state.SegmentCheckpoint.Digests())...,
		),
	)

	// TODO: Requesting all missing certificates from all the nodes known to have them right away is quite an overkill,
	//       resulting in a huge waste of resources. Be smarter about it by, for example, only asking a few nodes first.
}

// catchUpRequests assembles and returns a list of Events
// representing requests for retransmission of committed certificates.
// The list contains one request for each slot of the segment that has not yet been committed.
func catchUpRequests(
	state *common.State,
	moduleConfig common2.ModuleConfig,
	nodes []stdtypes.NodeID,
	digests map[tt.SeqNr][]byte,
) []stdtypes.Event {

	catchUpRequests := make([]stdtypes.Event, 0)

	// Deterministically iterate through all the (sequence number, certificate) pairs
	// received in a quorum of Done messages.
	maputil.IterateSorted(digests, func(sn tt.SeqNr, digest []byte) bool {

		// If no certificate has been committed for the sequence number, create a retransmission request.
		if !state.Slots[state.View][sn].Committed {
			catchUpRequests = append(catchUpRequests, transportpbevents.SendMessage(
				moduleConfig.Net,
				pbftpbmsgs.CatchUpRequest(moduleConfig.Self, digest, sn),
				nodes,
			).Pb())
		}
		return true
	})

	return catchUpRequests
}

// applyMsgCatchUpRequest applies a request for retransmitting a missing committed entry.
// It looks up the requested entry (more precisely, the corresponding Preprepare message)
// by its sequence number and digest and sends it to the originator of the request inside a CatchUpResponse message.
// If no matching Preprepare is found, does nothing.
func applyMsgCatchUpRequest(
	m dsl.Module,
	state *common.State,
	moduleConfig common2.ModuleConfig,
	digest []byte,
	sn tt.SeqNr,
	from stdtypes.NodeID,
) {
	if preprepare := state.LookUpPreprepare(sn, digest); preprepare != nil {

		// If the requested Preprepare message is available, send it to the originator of the request.
		// No need for periodic re-transmission. The requester will re-transmit the request if needed.
		transportpbdsl.SendMessage(
			m,
			moduleConfig.Net,
			pbftpbmsgs.CatchUpResponse(moduleConfig.Self, preprepare),
			[]stdtypes.NodeID{from})

	}

	// If the requested Preprepare message is not available, ignore the request.
}

// applyMsgCatchUpResponse applies a retransmitted missing committed certificate.
// It only requests hashing of the response,
// the actual handling of it being performed only when the hash result is available.
func applyMsgCatchUpResponse(
	m dsl.Module,
	state *common.State,
	moduleConfig common2.ModuleConfig,
	preprepare *pbftpbtypes.Preprepare,
	_ stdtypes.NodeID,
) {

	if preprepare == nil {
		return
	}

	if slot, ok := state.Slots[state.View][preprepare.Sn]; !ok {
		return
	} else if slot.Committed {
		// Ignore preprepare if received in the meantime.
		// This check is technically redundant, as it is (and must be) performed also after the Preprepare is hashed.
		// However, it might prevent some unnecessary hash computation if performed here as well.
		return
	}

	hasherpbdsl.RequestOne(
		m,
		moduleConfig.Hasher,
		common.SerializePreprepareForHashing(preprepare),
		&pbftpbtypes.CatchUpResponse{Resp: preprepare},
	)
}

// SendDoneMessages sends a Done message to all other nodes as part of the instance-level checkpoint subprotocol.
// This method is called when all slots have been committed.
func SendDoneMessages(
	m dsl.Module,
	state *common.State,
	params *common.ModuleParams,
	moduleConfig common2.ModuleConfig,
	logger logging.Logger,
) {

	logger.Log(logging.LevelInfo, "Done with segment.")

	// Collect the preprepare digests of all committed certificates.
	digests := make([][]byte, 0, state.Segment.Len())
	maputil.IterateSorted(state.Slots[state.View], func(_ tt.SeqNr, slot *common.PbftSlot) bool {
		digests = append(digests, slot.Digest)
		return true
	})

	// Periodically send a Done message with the digests to all other nodes.
	stddsl.TimerRepeat(
		m,
		moduleConfig.Timer,
		params.Config.DoneResendPeriod,
		stdtypes.RetentionIndex(params.Config.EpochNr),
		transportpbevents.SendMessage(
			moduleConfig.Net,
			pbftpbmsgs.Done(moduleConfig.Self, digests),
			state.Segment.NodeIDs(),
		).Pb(),
	)
}

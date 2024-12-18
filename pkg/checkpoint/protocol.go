// TODO: Finish writing proper comments in this file.

package checkpoint

import (
	"bytes"
	"time"

	es "github.com/go-errors/errors"

	"github.com/matejpavlovic/mir/pkg/dsl"
	"github.com/matejpavlovic/mir/pkg/logging"
	"github.com/matejpavlovic/mir/pkg/modules"
	apppbdsl "github.com/matejpavlovic/mir/pkg/pb/apppb/dsl"
	checkpointpbdsl "github.com/matejpavlovic/mir/pkg/pb/checkpointpb/dsl"
	checkpointpbmsgs "github.com/matejpavlovic/mir/pkg/pb/checkpointpb/msgs"
	checkpointpbtypes "github.com/matejpavlovic/mir/pkg/pb/checkpointpb/types"
	cryptopbdsl "github.com/matejpavlovic/mir/pkg/pb/cryptopb/dsl"
	hasherpbdsl "github.com/matejpavlovic/mir/pkg/pb/hasherpb/dsl"
	transportpbevents "github.com/matejpavlovic/mir/pkg/pb/transportpb/events"
	trantorpbdsl "github.com/matejpavlovic/mir/pkg/pb/trantorpb/dsl"
	trantorpbtypes "github.com/matejpavlovic/mir/pkg/pb/trantorpb/types"
	tt "github.com/matejpavlovic/mir/pkg/trantor/types"
	"github.com/matejpavlovic/mir/pkg/util/maputil"
	"github.com/matejpavlovic/mir/pkg/util/membutil"
	stddsl "github.com/matejpavlovic/mir/stdevents/dsl"
	"github.com/matejpavlovic/mir/stdtypes"
)

const (
	DefaultResendPeriod = time.Second
)

type State struct {
	// State snapshot associated with this checkpoint.
	StateSnapshot *trantorpbtypes.StateSnapshot

	// Hash of the state snapshot data associated with this checkpoint.
	StateSnapshotHash []byte

	// Set of nodes' valid checkpoint Signatures (these will make up a checkpoint certificate).
	Signatures map[stdtypes.NodeID][]byte

	// Set of nodes from which a (potentially invalid) Checkpoint messages has been received
	// (used to ignore duplicate messages).
	SigReceived map[stdtypes.NodeID]struct{}

	// Set of Checkpoint messages that were received ahead of time.
	PendingMessages map[stdtypes.NodeID]*checkpointpbtypes.Checkpoint

	// Flag ensuring that the stable checkpoint is only Announced once.
	// Set to true when announcing a stable checkpoint for the first time.
	// When true, stable checkpoints are not Announced anymore.
	Announced bool
}

// NewModule allocates and returns a new instance of the ModuleParams associated with sequence number sn.
func NewModule(
	moduleConfig ModuleConfig,
	params *ModuleParams,
	logger logging.Logger) modules.PassiveModule {

	state := &State{
		StateSnapshot: &trantorpbtypes.StateSnapshot{
			AppData: nil,
			EpochData: &trantorpbtypes.EpochData{
				EpochConfig:        params.EpochConfig,
				ClientProgress:     nil, // This will be filled by a separate event.
				LeaderPolicy:       params.LeaderPolicyData,
				PreviousMembership: params.Membership,
			},
		},
		Announced:       false,
		Signatures:      make(map[stdtypes.NodeID][]byte),
		SigReceived:     make(map[stdtypes.NodeID]struct{}),
		PendingMessages: make(map[stdtypes.NodeID]*checkpointpbtypes.Checkpoint),
	}

	m := dsl.NewModule(moduleConfig.Self)

	apppbdsl.UponSnapshot(m, func(appData []uint8) error {
		// Treat nil data as an empty byte slice.
		if appData == nil {
			appData = []byte{}
		}

		// Save the received app snapshot if there is none yet.
		if state.StateSnapshot.AppData == nil {
			state.StateSnapshot.AppData = appData
			if state.SnapshotReady() {
				if err := processStateSnapshot(m, state, moduleConfig); err != nil {
					return err
				}
			}
		}
		return nil
	})

	hasherpbdsl.UponResultOne(m, func(digest []uint8, _ *struct{}) error {
		// Save the received snapshot hash
		state.StateSnapshotHash = digest

		// Request signature
		sigData := serializeCheckpointForSig(params.EpochConfig.EpochNr, params.EpochConfig.FirstSn, state.StateSnapshotHash)

		cryptopbdsl.SignRequest(m, moduleConfig.Crypto, sigData, &struct{}{})

		return nil
	})

	cryptopbdsl.UponSignResult(m, func(sig []uint8, _ *struct{}) error {

		// Save received own checkpoint signature
		state.Signatures[params.OwnID] = sig
		state.SigReceived[params.OwnID] = struct{}{}

		// In case the node's own signature is enough to reach quorum, announce the stable checkpoint.
		// This can happen in a small system where no failures are tolerated.
		if state.Stable(params) {
			announceStable(m, params, state, moduleConfig)
		}

		// Send a checkpoint message to all nodes.
		chkpMessage := checkpointpbmsgs.Checkpoint(moduleConfig.Self, params.EpochConfig.EpochNr, params.EpochConfig.FirstSn, state.StateSnapshotHash, sig)
		sortedMembership := maputil.GetSortedKeys(params.Membership.Nodes)
		stddsl.TimerRepeat(m,
			"timer", // TODO: Put this value into moduleConfig.
			params.ResendPeriod,
			stdtypes.RetentionIndex(params.EpochConfig.EpochNr),
			transportpbevents.SendMessage(moduleConfig.Net, chkpMessage, sortedMembership).Pb(),
		)

		logger.Log(logging.LevelDebug, "Sending checkpoint message",
			"epoch", params.EpochConfig.EpochNr,
			"dataLen", len(state.StateSnapshot.AppData),
			"memberships", len(state.StateSnapshot.EpochData.EpochConfig.Memberships),
		)

		// Apply pending Checkpoint messages
		for s, msg := range state.PendingMessages {
			if err := applyCheckpointReceived(m, params, state, moduleConfig, s, msg.Epoch, msg.Sn, msg.SnapshotHash, msg.Signature, logger); err != nil {
				logger.Log(logging.LevelWarn, "Error applying pending Checkpoint message", "error", err, "msg", msg)
				return err
			}

		}
		state.PendingMessages = nil

		return nil
	})

	cryptopbdsl.UponSigVerified(m, func(nodeId stdtypes.NodeID, err error, c *verificationContext) error {
		if err != nil {
			logger.Log(logging.LevelWarn, "Ignoring Checkpoint message. Invalid signature.", "source", nodeId, "error", err)
			return nil
		}

		// Note the reception of a valid Checkpoint message from node `source`.
		state.Signatures[nodeId] = c.signature

		// If, after having applied this message, the checkpoint became stable, produce the necessary events.
		if state.Stable(params) {
			announceStable(m, params, state, moduleConfig)
		}

		return nil
	})

	trantorpbdsl.UponClientProgress(m, func(progress map[tt.ClientID]*trantorpbtypes.DeliveredTXs) error {
		// Save the received client progress if there is none yet.
		if state.StateSnapshot.EpochData.ClientProgress == nil {
			state.StateSnapshot.EpochData.ClientProgress = &trantorpbtypes.ClientProgress{
				Progress: progress,
			}
			if state.SnapshotReady() {
				if err := processStateSnapshot(m, state, moduleConfig); err != nil {
					return err
				}
			}
		}
		return nil
	})

	checkpointpbdsl.UponCheckpointReceived(m, func(from stdtypes.NodeID, epoch tt.EpochNr, sn tt.SeqNr, snapshotHash []uint8, signature []uint8) error {
		return applyCheckpointReceived(m, params, state, moduleConfig, from, epoch, sn, snapshotHash, signature, logger)
	})

	return m
}

func processStateSnapshot(m dsl.Module, state *State, mc ModuleConfig) error {

	// Serialize the snapshot.
	snapshotData, err := serializeSnapshotForHash(state.StateSnapshot)
	if err != nil {
		return es.Errorf("failed serializing state snapshot: %w", err)
	}

	// Initiate computing the hash of the snapshot.
	hasherpbdsl.RequestOne(m,
		mc.Hasher,
		snapshotData,
		&struct{}{},
	)

	return nil
}

func announceStable(m dsl.Module, p *ModuleParams, state *State, mc ModuleConfig) {

	// Only announce the stable checkpoint once.
	if state.Announced {
		return
	}
	state.Announced = true

	// Assemble a multisig certificate from the received valid signatures.
	cert := make(map[stdtypes.NodeID][]byte)
	for node, sig := range state.Signatures {
		cert[node] = sig
	}

	// Announce the stable checkpoint to the ordering protocol.
	checkpointpbdsl.StableCheckpoint(m, mc.Ord, p.EpochConfig.FirstSn, state.StateSnapshot, cert)
}

func applyCheckpointReceived(m dsl.Module,
	p *ModuleParams,
	state *State,
	moduleConfig ModuleConfig,
	from stdtypes.NodeID,
	epoch tt.EpochNr,
	sn tt.SeqNr,
	snapshotHash []uint8,
	signature []uint8,
	logger logging.Logger) error {

	if sn != p.EpochConfig.FirstSn {
		logger.Log(logging.LevelWarn, "invalid sequence number %v, expected %v as first sequence number.\n", sn, p.EpochConfig.FirstSn)
		return nil
	}

	if epoch != p.EpochConfig.EpochNr {
		logger.Log(logging.LevelWarn, "invalid epoch number %v, expected %v.\n", epoch, p.EpochConfig.EpochNr)
		return nil
	}

	// check if from is part of the membership
	if _, ok := p.Membership.Nodes[from]; !ok {
		logger.Log(logging.LevelWarn, "sender %s is not a member.\n", from)
		return nil
	}

	// Notify the protocol about the progress of the from node.
	// If no progress is made for a configured number of epochs,
	// the node is considered to be a straggler and is sent a stable checkpoint to catch uparams.
	checkpointpbdsl.EpochProgress(m, moduleConfig.Ord, from, epoch)

	// If checkpoint is already stable, ignore message.
	if state.Stable(p) {
		return nil
	}

	// Check snapshot hash
	if state.StateSnapshotHash == nil {
		// The message is received too early, put it aside
		state.PendingMessages[from] = &checkpointpbtypes.Checkpoint{
			Epoch:        epoch,
			Sn:           sn,
			SnapshotHash: snapshotHash,
			Signature:    signature,
		}
		return nil
	} else if !bytes.Equal(state.StateSnapshotHash, snapshotHash) {
		// Snapshot hash mismatch
		logger.Log(logging.LevelWarn, "Ignoring Checkpoint message. Mismatching state snapshot hash.", "from", from)
		return nil
	}

	// Ignore duplicate messages.
	if _, ok := state.SigReceived[from]; ok {
		return nil
	}
	state.SigReceived[from] = struct{}{}

	// Verify signature of the sender.
	sigData := serializeCheckpointForSig(p.EpochConfig.EpochNr, p.EpochConfig.FirstSn, state.StateSnapshotHash)
	cryptopbdsl.VerifySig(m,
		moduleConfig.Crypto,
		sigData,
		signature,
		from,
		&verificationContext{signature: signature},
	)

	return nil
}

func (state *State) SnapshotReady() bool {
	return state.StateSnapshot.AppData != nil &&
		state.StateSnapshot.EpochData.ClientProgress != nil
}

func (state *State) Stable(p *ModuleParams) bool {
	return state.SnapshotReady() && membutil.HaveStrongQuorum(p.Membership, maputil.GetKeys(state.Signatures))
}

type verificationContext struct {
	signature []uint8
}

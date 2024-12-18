package chkpvalidator

import (
	"math"

	es "github.com/go-errors/errors"

	"github.com/matejpavlovic/mir/pkg/checkpoint"
	"github.com/matejpavlovic/mir/pkg/crypto"
	checkpointpbtypes "github.com/matejpavlovic/mir/pkg/pb/checkpointpb/types"
	trantorpbtypes "github.com/matejpavlovic/mir/pkg/pb/trantorpb/types"
	tt "github.com/matejpavlovic/mir/pkg/trantor/types"
	t "github.com/matejpavlovic/mir/stdtypes"
)

type ConservativeCV struct {
	configOffset int
	ownID        t.NodeID
	hashImpl     crypto.HashImpl
	chkpVerifier checkpoint.Verifier
}

// NewConservativeCV returns a new ConservativeCV. This checkpoint validity checker
// simply rejects checkpoints whose signatures cannot be verified because
// the node does not know the membership of the relevant epoch yet/anymore.
func NewConservativeCV(
	configOffset int,
	ownID t.NodeID,
	hashImpl crypto.HashImpl,
	chkpVerifier checkpoint.Verifier,
) *ConservativeCV {
	return &ConservativeCV{
		configOffset: configOffset,
		ownID:        ownID,
		hashImpl:     hashImpl,
		chkpVerifier: chkpVerifier,
	}
}

func (ccv *ConservativeCV) Verify(
	chkp *checkpointpbtypes.StableCheckpoint,
	epochNr tt.EpochNr,
	memberships []*trantorpbtypes.Membership,
) error {
	sc := checkpoint.StableCheckpointFromPb(chkp.Pb())

	// Check syntactic validity of the checkpoint.
	if err := sc.SyntacticCheck(ccv.configOffset); err != nil {
		return err
	}

	// We consider a checkpoint invalid if we are not part of its membership
	// (more precisely, membership of the epoch the checkpoint is at the start of).
	// Correct nodes should never send such checkpoints, but faulty ones could.
	if _, ok := sc.Memberships()[0].Nodes[ccv.ownID]; !ok {
		return es.Errorf("nodeID not in membership")
	}

	// Check if epoch is in integer bounds.
	if sc.Epoch() > math.MaxInt || epochNr > math.MaxInt {
		return es.Errorf("epoch number out of integer range")
	}

	// Check how far the received stable checkpoint is ahead of the local node's state.
	// Integer casting required here to prevent underflow.
	chkpMembershipOffset := int(sc.Epoch()) - 1 - int(epochNr) //nolint:gosec
	if chkpMembershipOffset <= 0 {
		// Ignore stable checkpoints that are not far enough
		// ahead of the current state of the local node.
		return es.Errorf("checkpoint not far ahead enough")
	}

	// Make sure ccv.configOffset is non-negative before conversion
	if ccv.configOffset < 0 {
		return es.Errorf("configOffset cannot be negative")
	}

	if chkpMembershipOffset > ccv.configOffset {
		// cannot verify checkpoint signatures, too far ahead
		return es.Errorf("checkpoint too far ahead")
	}

	chkpMembership := memberships[chkpMembershipOffset]

	return sc.Verify(ccv.configOffset, ccv.hashImpl, ccv.chkpVerifier, chkpMembership)
}

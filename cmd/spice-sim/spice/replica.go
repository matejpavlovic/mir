package spice

import (
	"fmt"

	"github.com/matejpavlovic/mir/cmd/spice-sim/events"
	"github.com/matejpavlovic/mir/cmd/spice-sim/types"
	"github.com/matejpavlovic/mir/pkg/dsl"
	"github.com/matejpavlovic/mir/pkg/modules"
	"github.com/matejpavlovic/mir/stdtypes"
)

func NewReplica(id stdtypes.ModuleID, shard int64, coreState CoreState) modules.Module {
	thisReplica := dsl.NewModule(id)

	executionHead := types.GenesisBlock().Id()
	var executionHeadPredecessor types.BlockID
	receiptsReceived := newReceiptBuffer()

	endorsements := NewEndorsementBuffer()

	canAdvanceState := func() bool {
		chunkID := types.ChunkID{executionHead, shard}
		return coreState.ChunkAvailable(chunkID) &&
			(executionHead.Height == 0 ||
				receiptsReceived.NumReceiptLists(executionHead.Height-1) == coreState.config.NumShards &&
					endorsements.StateCertified(executionHeadPredecessor, coreState.config.NumShards, coreState.config.StateCertQuorum))
	}

	advanceState := func(timestamp int64) bool {
		chunkID := types.ChunkID{executionHead, shard}
		fmt.Printf("(%v) %v: Executing Chunk: %s\n", timestamp, id, chunkID)
		for _, validatorID := range coreState.ValidatorsAt(chunkID.BlockID.Height, shard) {
			dsl.EmitEvent(thisReplica, events.NewStateWitnessEvent(
				id,
				validatorID,
				timestamp+coreState.config.ExecutionDelay+coreState.config.StateWitnessTransmissionDelay,
				chunkID,
			))
		}

		for targetShard := int64(0); targetShard < coreState.config.NumShards; targetShard++ {
			for _, replicaID := range coreState.ReplicaIDs(targetShard) {
				dsl.EmitEvent(thisReplica, events.NewReceiptsEvent(
					id,
					replicaID,
					timestamp+coreState.config.ExecutionDelay+coreState.config.ReceiptTransmissionDelay,
					chunkID,
				))
			}
		}

		executionHeadPredecessor = executionHead
		executionHead = coreState.NextCanonicalBlock(executionHead).Id()
		fmt.Printf("(%v) %v: Moved execution head: %s\n", timestamp, id, executionHead)

		return true
	}

	dsl.UponEvent(thisReplica, func(ev *events.InitEvent) error {
		return nil
	})

	dsl.UponEvent(thisReplica, func(ev *events.NewBlockEvent) error {
		fmt.Printf("(%v) %v: Received block: %s\n", ev.Timestamp(), id, ev.Block.Id())
		coreState.ApplyBlock(ev.Block)
		for _, cert := range ev.Block.AvailabilityCerts {
			fmt.Printf("(%v) %v: Av cert: %s\n", ev.Timestamp(), id, cert.ID())
		}
		for _, cert := range ev.Block.StateCerts {
			fmt.Printf("(%v) %v: State cert: %s\n", ev.Timestamp(), id, cert.ID())
		}

		timestamp := ev.Timestamp()
		for canAdvanceState() {
			advanceState(timestamp)
			timestamp += coreState.config.ExecutionDelay
		}

		return nil
	})

	dsl.UponEvent(thisReplica, func(ev *events.ReceiptsEvent) error {
		fmt.Printf("(%v) %v: Received receipts: %s\n", ev.Timestamp(), id, ev.ChunkID)
		receiptsReceived.Add(ev.ChunkID.BlockID.Height, ev.ChunkID.Shard)
		timestamp := ev.Timestamp()
		for canAdvanceState() {
			advanceState(timestamp)
			timestamp += coreState.config.ExecutionDelay
		}
		return nil
	})

	dsl.UponEvent(thisReplica, func(ev *events.EndorseExecResultEvent) error {

		endorsements.Add(ev)

		if int64(endorsements.NumEndorsements(ev.ChunkID)) == coreState.config.StateCertQuorum {
			timestamp := ev.Timestamp()
			for canAdvanceState() {
				advanceState(timestamp)
				timestamp += coreState.config.ExecutionDelay
			}
		}

		return nil
	})

	return thisReplica
}

// Usage: receiptBuffer[height][shard]
type receiptBuffer map[int64]map[int64]struct{}

func newReceiptBuffer() receiptBuffer {
	return make(map[int64]map[int64]struct{})
}

func (rb receiptBuffer) Add(height int64, shard int64) {
	if _, ok := rb[height]; !ok {
		rb[height] = make(map[int64]struct{})
	}
	rb[height][shard] = struct{}{}
}

func (rb receiptBuffer) NumReceiptLists(height int64) int64 {
	return int64(len(rb[height]))
}

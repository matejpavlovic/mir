package spice

import (
	"fmt"

	"github.com/matejpavlovic/mir/cmd/spice-sim/events"
	"github.com/matejpavlovic/mir/cmd/spice-sim/types"
	"github.com/matejpavlovic/mir/pkg/dsl"
	"github.com/matejpavlovic/mir/pkg/modules"
	"github.com/matejpavlovic/mir/pkg/util/maputil"
	"github.com/matejpavlovic/mir/stdtypes"
)

type ChunkPartBuffer struct {
	chunkID types.ChunkID
	parts   map[int64]struct{}
	signers map[stdtypes.ModuleID]struct{}
}

func NewChunkPartBuffer(chunkID types.ChunkID) ChunkPartBuffer {
	return ChunkPartBuffer{
		chunkID: chunkID,
		parts:   make(map[int64]struct{}),
		signers: make(map[stdtypes.ModuleID]struct{}),
	}
}

func (cpb *ChunkPartBuffer) Len() int {
	return len(cpb.parts)
}

func (cpb *ChunkPartBuffer) Add(signer stdtypes.ModuleID, index int64) {
	cpb.signers[signer] = struct{}{}
	cpb.parts[index] = struct{}{}
}

func (cpb *ChunkPartBuffer) Certificate() types.AvailabilityCert {
	return types.NewAvailabilityCert(cpb.chunkID, maputil.GetSortedKeys(cpb.signers))
}

func NewBlockProducer(id stdtypes.ModuleID, coreState CoreState) modules.Module {
	thisBlockProducer := dsl.NewModule(id)

	confirmedChunkParts := map[string]ChunkPartBuffer{}
	pendingAvailabilityCerts := map[string]types.AvailabilityCert{}

	endorsements := NewEndorsementBuffer()
	pendingStateCerts := map[string]types.StateCert{}

	broadcastBlock := func(timestamp int64, block *types.Block) {
		for _, blockProducerID := range coreState.BlockProducerIDs() {
			blockEvent := events.NewNewBlockEvent(
				id,
				blockProducerID,
				timestamp,
				block,
			)
			dsl.EmitEvent(thisBlockProducer, blockEvent)
		}
		for _, chunkProducerID := range coreState.ChunkProducerIDsAllShards() {
			blockEvent := events.NewNewBlockEvent(
				id,
				chunkProducerID,
				timestamp,
				block,
			)
			dsl.EmitEvent(thisBlockProducer, blockEvent)
		}
		for _, replicaID := range coreState.ReplicaIDsAllShards() {
			blockEvent := events.NewNewBlockEvent(
				id,
				replicaID,
				timestamp,
				block,
			)
			dsl.EmitEvent(thisBlockProducer, blockEvent)
		}
	}

	dsl.UponEvent(thisBlockProducer, func(ev *events.InitEvent) error {
		fmt.Printf("Initializing block producer: %v\n", id)
		if coreState.BlockProducerAt(0) == id {

			genesis := types.GenesisBlock()

			fmt.Printf("(%v) %v: Producing genesis block: %s (%d av certs) (%d state certs)\n",
				ev.Timestamp(),
				id,
				genesis.Id(),
				len(pendingAvailabilityCerts),
				len(pendingStateCerts),
			)

			broadcastBlock(
				ev.Timestamp()+coreState.config.BlockTime, genesis,
			)
		}
		return nil
	})

	dsl.UponEvent(thisBlockProducer, func(ev *events.NewBlockEvent) error {
		coreState.ApplyBlock(ev.Block)

		for _, cert := range ev.Block.AvailabilityCerts {
			delete(pendingAvailabilityCerts, cert.ID().String())
		}
		for _, cert := range ev.Block.StateCerts {
			delete(pendingStateCerts, cert.ID().String())
		}

		nextHeight := ev.Block.Height + 1
		if coreState.BlockProducerAt(nextHeight) == id {
			newBlock := types.NewBlock(
				nextHeight,
				ev.Block.Hash(),
				maputil.GetValues(pendingAvailabilityCerts),
				maputil.GetValues(pendingStateCerts),
			)

			fmt.Printf("(%v) %v: Producing new block: %s (%d av certs) (%d state certs)\n", ev.Timestamp(), id, newBlock.Id(), len(pendingAvailabilityCerts), len(pendingStateCerts))

			broadcastBlock(
				ev.Timestamp()+coreState.config.BlockTime,
				newBlock,
			)
		}
		return nil
	})

	dsl.UponEvent(thisBlockProducer, func(ev *events.ChunkPartStoredEvent) error {
		// Save confirmed chunk part.
		chunkKey := ev.PartID.ChunkID.String()
		chunkParts, ok := confirmedChunkParts[chunkKey]
		if !ok {
			chunkParts = NewChunkPartBuffer(ev.PartID.ChunkID)
			confirmedChunkParts[chunkKey] = chunkParts
		}

		// If we want actual fault tolerance, we need to check here if the part owner is really assigned to this part.
		// Below, we then need to check if a sufficient number of distinct owners confirmed a sufficient number of
		// different chunk parts to guarantee reconstruction despite failures.
		// The analogous needs to be done for endorsements (except that they don't have parts).

		chunkParts.Add(ev.SrcModule, ev.PartID.Index)

		if int64(chunkParts.Len()) == coreState.config.ChunkDataStatementQuorum {
			fmt.Printf("(%v) %v: Chunk availability certified: %s\n", ev.Timestamp(), id, ev.PartID.ChunkID)
			cert := chunkParts.Certificate()
			pendingAvailabilityCerts[cert.ID().String()] = cert
		}

		return nil
	})

	dsl.UponEvent(thisBlockProducer, func(ev *events.EndorseExecResultEvent) error {

		endorsements.Add(ev)

		if int64(endorsements.NumEndorsements(ev.ChunkID)) == coreState.config.StateCertQuorum {
			fmt.Printf("(%v) %v: Execution result certified: %s\n", ev.Timestamp(), id, ev.ChunkID)
			cert := endorsements.Certificate(ev.ChunkID)
			pendingStateCerts[cert.ID().String()] = cert
		}

		return nil
	})

	return thisBlockProducer
}

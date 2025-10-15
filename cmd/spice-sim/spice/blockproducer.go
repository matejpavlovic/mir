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
	parts   map[int64]struct{}
	signers map[stdtypes.ModuleID]struct{}
}

func NewChunkPartBuffer() ChunkPartBuffer {
	return ChunkPartBuffer{
		parts:   make(map[int64]struct{}),
		signers: make(map[stdtypes.ModuleID]struct{}),
	}
}

func NewBlockProducer(id stdtypes.ModuleID, coreState CoreState) modules.Module {
	bp := dsl.NewModule(id)

	confirmedChunkParts := map[string]ChunkPartBuffer{}
	pendingAvailabilityCerts := map[string]types.AvailabilityCert{}

	broadcastBlock := func(timestamp int64, block *types.Block) {
		for _, blockProducerID := range coreState.BlockProducerIDs() {
			blockEvent := events.NewNewBlockEvent(
				id,
				blockProducerID,
				timestamp,
				block,
			)
			dsl.EmitEvent(bp, blockEvent)
		}
		for _, chunkProducerID := range coreState.ChunkProducerIDsAllShards() {
			blockEvent := events.NewNewBlockEvent(
				id,
				chunkProducerID,
				timestamp,
				block,
			)
			dsl.EmitEvent(bp, blockEvent)
		}
	}

	dsl.UponEvent(bp, func(ev *events.InitEvent) error {
		fmt.Printf("Initializing block producer: %v\n", id)
		if coreState.BlockProducerAt(0) == id {
			newBlock := types.NewBlock(
				types.NewBlockID(0, []byte{}),
				maputil.GetValues(pendingAvailabilityCerts),
			)

			fmt.Printf("(%v) %v: Producing new block: %s\n", ev.Timestamp(), id, newBlock.Id)

			broadcastBlock(
				ev.Timestamp()+int64(coreState.config.BlockTime), newBlock,
			)
		}
		return nil
	})

	dsl.UponEvent(bp, func(ev *events.NewBlockEvent) error {
		for _, cert := range ev.Block.AvailabilityCerts {
			delete(pendingAvailabilityCerts, cert.ID().String())
		}

		nextHeight := ev.Block.Id.Height + 1
		if coreState.BlockProducerAt(nextHeight) == id {
			newBlock := types.NewBlock(
				types.NewBlockID(nextHeight, ev.Block.Hash()),
				maputil.GetValues(pendingAvailabilityCerts),
			)

			fmt.Printf("(%v) %v: Producing new block :%s\n", ev.Timestamp(), id, newBlock.Id)

			broadcastBlock(
				ev.Timestamp()+int64(coreState.config.BlockTime),
				newBlock,
			)
		}
		return nil
	})

	dsl.UponEvent(bp, func(ev *events.ChunkPartStoredEvent) error {
		// Save confirmed chunk part.
		chunkKey := ev.PartID.ChunkID.String()
		chunkParts, ok := confirmedChunkParts[chunkKey]
		if !ok {
			chunkParts = NewChunkPartBuffer()
			confirmedChunkParts[chunkKey] = chunkParts
		}

		// If we want actual fault tolerance, we need to check here if the part owner is really assigned to this part.
		// Below, we then need to check if a sufficient number of distinct owners confirmed a sufficient number of
		// different chunk parts to guarantee reconstruction despite failures.

		chunkParts.parts[ev.PartID.Index] = struct{}{}
		chunkParts.signers[ev.SrcModule] = struct{}{}

		if len(chunkParts.parts) == coreState.config.ChunkDataStatementQuorum {
			fmt.Printf("(%v) %v: Chunk availability certified: %s\n", ev.Timestamp(), id, ev.PartID.ChunkID)
			cert := types.NewAvailabilityCert(ev.PartID.ChunkID, maputil.GetSortedKeys(chunkParts.signers))
			pendingAvailabilityCerts[cert.ID().String()] = cert
		}

		return nil
	})

	return bp
}

package spice

import (
	"fmt"
	"github.com/matejpavlovic/mir/cmd/spice-sim/events"
	"github.com/matejpavlovic/mir/cmd/spice-sim/types"
	"github.com/matejpavlovic/mir/pkg/dsl"
	"github.com/matejpavlovic/mir/pkg/modules"
	"github.com/matejpavlovic/mir/stdtypes"
)

func NewChunkProducer(id stdtypes.ModuleID, shard int64, coreState CoreState) modules.Module {
	bp := dsl.NewModule(id)

	disperseChunk := func(timestamp int64, chunk types.Chunk) {
		for i, dataOwnerID := range coreState.DataOwnersAt(chunk.Id.BlockID.Height, chunk.Id.Shard) {
			dataPart := events.NewChunkPartEvent(id, dataOwnerID, timestamp, types.ChunkPartID{
				ChunkID: chunk.Id,
				Index:   int64(i),
			})
			dsl.EmitEvent(bp, dataPart)
		}
	}

	dsl.UponEvent(bp, func(ev *events.InitEvent) error {
		fmt.Printf("Initializing chunk producer: %v\n", id)
		if coreState.ChunkProducerAt(0, shard) == id {
			newChunk := types.NewChunk(types.ChunkID{
				BlockID: types.NewBlockID(0, nil),
				Shard:   shard,
			})
			fmt.Printf("(%v) %v: Producing new chunk: %s\n", ev.Timestamp(), id, newChunk.Id)
			disperseChunk(ev.Timestamp()+int64(coreState.config.ChunkDispersalDelay), newChunk)
		}
		return nil
	})

	dsl.UponEvent(bp, func(ev *events.NewBlockEvent) error {
		nextHeight := ev.Block.Id.Height
		if coreState.ChunkProducerAt(nextHeight, shard) == id {
			newChunk := types.NewChunk(types.ChunkID{
				BlockID: types.NewBlockID(nextHeight, ev.Block.Hash()),
				Shard:   shard,
			})

			fmt.Printf("(%v) %v: Producing new chunk: %s\n", ev.Timestamp(), id, newChunk.Id)
			disperseChunk(ev.Timestamp()+int64(coreState.config.ChunkDispersalDelay), newChunk)
		}
		return nil
	})

	return bp
}

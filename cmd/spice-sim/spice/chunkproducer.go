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
	thisChunkProducer := dsl.NewModule(id)

	disperseChunk := func(timestamp int64, chunk types.Chunk) {
		for i, dataOwnerID := range coreState.DataOwnersAt(chunk.Id.BlockID.Height, chunk.Id.Shard) {
			dataPart := events.NewChunkPartEvent(id, dataOwnerID, timestamp, types.ChunkPartID{
				ChunkID: chunk.Id,
				Index:   int64(i),
			})
			dsl.EmitEvent(thisChunkProducer, dataPart)
		}
	}

	dsl.UponEvent(thisChunkProducer, func(ev *events.InitEvent) error {
		return nil
	})

	dsl.UponEvent(thisChunkProducer, func(ev *events.NewBlockEvent) error {
		coreState.ApplyBlock(ev.Block)

		if coreState.ChunkProducerAt(ev.Block.Height, shard) == id {
			newChunk := types.NewChunk(types.ChunkID{
				BlockID: ev.Block.Id(),
				Shard:   shard,
			})

			fmt.Printf("(%v) %v: Producing new chunk: %s\n", ev.Timestamp(), id, newChunk.Id)
			disperseChunk(ev.Timestamp()+coreState.config.ChunkDispersalDelay, newChunk)
		}
		return nil
	})

	return thisChunkProducer
}

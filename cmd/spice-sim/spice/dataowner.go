package spice

import (
	"fmt"
	"github.com/matejpavlovic/mir/cmd/spice-sim/events"
	"github.com/matejpavlovic/mir/cmd/spice-sim/types"
	"github.com/matejpavlovic/mir/pkg/dsl"
	"github.com/matejpavlovic/mir/pkg/modules"
	"github.com/matejpavlovic/mir/stdtypes"
)

func NewDataOwner(id stdtypes.ModuleID, coreState CoreState) modules.Module {
	thisDataOwner := dsl.NewModule(id)

	storedParts := map[string]map[int64]struct{}{}

	submitStatementChunkPartStored := func(timestamp int64, partID types.ChunkPartID) {
		for _, blockProducerID := range coreState.BlockProducerIDs() {
			statement := events.NewChunkPartStoredEvent(id, blockProducerID, timestamp, partID)
			dsl.EmitEvent(thisDataOwner, statement)
		}
	}

	dsl.UponEvent(thisDataOwner, func(ev *events.InitEvent) error {
		return nil
	})

	dsl.UponEvent(thisDataOwner, func(ev *events.ChunkPartEvent) error {
		// Save received chunk part.
		chunkKey := ev.PartID.String()
		chunkParts, ok := storedParts[chunkKey]
		if !ok {
			chunkParts = map[int64]struct{}{}
			storedParts[chunkKey] = chunkParts
		}
		chunkParts[ev.PartID.Index] = struct{}{}

		// Confirm storing of chunk to block producers
		fmt.Printf("(%v) %v Chunk part stored: %s\n", ev.Timestamp(), id, ev.PartID)
		submitStatementChunkPartStored(
			ev.Timestamp()+coreState.config.StatementSubmissionDelay,
			ev.PartID,
		)

		return nil
	})

	return thisDataOwner
}

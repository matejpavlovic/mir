package spice

import (
	"fmt"

	"github.com/matejpavlovic/mir/cmd/spice-sim/events"
	"github.com/matejpavlovic/mir/cmd/spice-sim/types"
	"github.com/matejpavlovic/mir/pkg/dsl"
	"github.com/matejpavlovic/mir/pkg/modules"
	"github.com/matejpavlovic/mir/stdtypes"
)

func NewValidator(id stdtypes.ModuleID, coreState CoreState) modules.Module {
	thisValidator := dsl.NewModule(id)

	submitStatementEndorseExecResult := func(timestamp int64, chunkID types.ChunkID) {
		for _, blockProducerID := range coreState.BlockProducerIDs() {
			statement := events.NewEndorseExecResultEvent(id, blockProducerID, timestamp, chunkID)
			dsl.EmitEvent(thisValidator, statement)
		}
		for _, replicaID := range coreState.ReplicaIDsAllShards() {
			statement := events.NewEndorseExecResultEvent(id, replicaID, timestamp, chunkID)
			dsl.EmitEvent(thisValidator, statement)
		}
	}

	dsl.UponEvent(thisValidator, func(ev *events.InitEvent) error {
		return nil
	})

	dsl.UponEvent(thisValidator, func(ev *events.StateWitnessEvent) error {
		// Endorse execution result.
		fmt.Printf("(%v) %v Validating state witness: %s\n", ev.Timestamp(), id, ev.ChunkID)
		submitStatementEndorseExecResult(
			ev.Timestamp()+coreState.config.StateWitnessValidationDelay+coreState.config.StatementSubmissionDelay,
			ev.ChunkID,
		)

		return nil
	})

	return thisValidator
}

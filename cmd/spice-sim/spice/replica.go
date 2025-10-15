package spice

import (
	"github.com/matejpavlovic/mir/cmd/spice-sim/events"
	"github.com/matejpavlovic/mir/pkg/dsl"
	"github.com/matejpavlovic/mir/pkg/modules"
	"github.com/matejpavlovic/mir/stdtypes"
)

func NewReplica(id stdtypes.ModuleID, shard int64, coreState CoreState) modules.Module {
	bp := dsl.NewModule(id)

	dsl.UponEvent(bp, func(ev *events.InitEvent) error {
		return nil
	})

	dsl.UponEvent(bp, func(ev *events.NewBlockEvent) error {
		// TODO: Continue here.
		return nil
	})

	return bp
}

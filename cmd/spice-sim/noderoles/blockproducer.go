package noderoles

import (
	"fmt"
	"github.com/matejpavlovic/mir/cmd/spice-sim/events"
	"github.com/matejpavlovic/mir/pkg/dsl"
	"github.com/matejpavlovic/mir/pkg/modules"
	"github.com/matejpavlovic/mir/stdevents"
	"github.com/matejpavlovic/mir/stdtypes"
)

func NewBlockProducer(id stdtypes.ModuleID) modules.Module {
	bp := dsl.NewModule(id)

	dsl.UponEvent(bp, func(ev *stdevents.Init) error {
		if id.String() == "block-producer-0" {
			dsl.EmitEvent(bp, events.NewNewBlockEvent(id, id, 0, 0))
		}
		return nil
	})

	dsl.UponEvent(bp, func(ev *events.NewBlockEvent) error {
		fmt.Println(ev.ToString())
		if ev.Timestamp < 1000 {
			dsl.EmitEvent(bp, events.NewNewBlockEvent(id, id, ev.Timestamp+100, ev.Height+1))
		}
		return nil
	})

	return bp
}

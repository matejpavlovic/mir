// Code generated by Mir codegen. DO NOT EDIT.

package ordererpbdsl

import (
	dsl "github.com/matejpavlovic/mir/pkg/dsl"
	types1 "github.com/matejpavlovic/mir/pkg/pb/eventpb/types"
	types "github.com/matejpavlovic/mir/pkg/pb/ordererpb/types"
)

// Module-specific dsl functions for processing events.

func UponEvent[W types.Event_TypeWrapper[Ev], Ev any](m dsl.Module, handler func(ev *Ev) error) {
	dsl.UponMirEvent[*types1.Event_Orderer](m, func(ev *types.Event) error {
		w, ok := ev.Type.(W)
		if !ok {
			return nil
		}

		return handler(w.Unwrap())
	})
}

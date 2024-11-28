// Code generated by Mir codegen. DO NOT EDIT.

package transportpbdsl

import (
	dsl "github.com/matejpavlovic/mir/pkg/dsl"
	types1 "github.com/matejpavlovic/mir/pkg/pb/eventpb/types"
	types2 "github.com/matejpavlovic/mir/pkg/pb/messagepb/types"
	types "github.com/matejpavlovic/mir/pkg/pb/transportpb/types"
	stdtypes "github.com/matejpavlovic/mir/stdtypes"
)

// Module-specific dsl functions for processing events.

func UponEvent[W types.Event_TypeWrapper[Ev], Ev any](m dsl.Module, handler func(ev *Ev) error) {
	dsl.UponMirEvent[*types1.Event_Transport](m, func(ev *types.Event) error {
		w, ok := ev.Type.(W)
		if !ok {
			return nil
		}

		return handler(w.Unwrap())
	})
}

func UponSendMessage(m dsl.Module, handler func(msg *types2.Message, destinations []stdtypes.NodeID) error) {
	UponEvent[*types.Event_SendMessage](m, func(ev *types.SendMessage) error {
		return handler(ev.Msg, ev.Destinations)
	})
}

func UponMessageReceived(m dsl.Module, handler func(from stdtypes.NodeID, msg *types2.Message) error) {
	UponEvent[*types.Event_MessageReceived](m, func(ev *types.MessageReceived) error {
		return handler(ev.From, ev.Msg)
	})
}

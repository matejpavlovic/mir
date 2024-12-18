// Code generated by Mir codegen. DO NOT EDIT.

package transportpbdsl

import (
	dsl "github.com/matejpavlovic/mir/pkg/dsl"
	types "github.com/matejpavlovic/mir/pkg/pb/messagepb/types"
	events "github.com/matejpavlovic/mir/pkg/pb/transportpb/events"
	stdtypes "github.com/matejpavlovic/mir/stdtypes"
)

// Module-specific dsl functions for emitting events.

func SendMessage(m dsl.Module, destModule stdtypes.ModuleID, msg *types.Message, destinations []stdtypes.NodeID) {
	dsl.EmitMirEvent(m, events.SendMessage(destModule, msg, destinations))
}

func MessageReceived(m dsl.Module, destModule stdtypes.ModuleID, from stdtypes.NodeID, msg *types.Message) {
	dsl.EmitMirEvent(m, events.MessageReceived(destModule, from, msg))
}

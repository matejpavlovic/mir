// Code generated by Mir codegen. DO NOT EDIT.

package pbftpbdsl

import (
	dsl "github.com/matejpavlovic/mir/pkg/dsl"
	types "github.com/matejpavlovic/mir/pkg/orderers/types"
	events "github.com/matejpavlovic/mir/pkg/pb/pbftpb/events"
	stdtypes "github.com/matejpavlovic/mir/stdtypes"
)

// Module-specific dsl functions for emitting events.

func ProposeTimeout(m dsl.Module, destModule stdtypes.ModuleID, proposeTimeout uint64) {
	dsl.EmitMirEvent(m, events.ProposeTimeout(destModule, proposeTimeout))
}

func ViewChangeSNTimeout(m dsl.Module, destModule stdtypes.ModuleID, view types.ViewNr, numCommitted uint64) {
	dsl.EmitMirEvent(m, events.ViewChangeSNTimeout(destModule, view, numCommitted))
}

func ViewChangeSegTimeout(m dsl.Module, destModule stdtypes.ModuleID, viewChangeSegTimeout uint64) {
	dsl.EmitMirEvent(m, events.ViewChangeSegTimeout(destModule, viewChangeSegTimeout))
}

// Code generated by Mir codegen. DO NOT EDIT.

package testerpbdsl

import (
	dsl "github.com/matejpavlovic/mir/pkg/dsl"
	events "github.com/matejpavlovic/mir/pkg/pb/testerpb/events"
	stdtypes "github.com/matejpavlovic/mir/stdtypes"
)

// Module-specific dsl functions for emitting events.

func Tester(m dsl.Module, destModule stdtypes.ModuleID) {
	dsl.EmitMirEvent(m, events.Tester(destModule))
}

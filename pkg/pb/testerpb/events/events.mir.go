// Code generated by Mir codegen. DO NOT EDIT.

package testerpbevents

import (
	types "github.com/matejpavlovic/mir/pkg/pb/eventpb/types"
	types1 "github.com/matejpavlovic/mir/pkg/pb/testerpb/types"
	stdtypes "github.com/matejpavlovic/mir/stdtypes"
)

func Tester(destModule stdtypes.ModuleID) *types.Event {
	return &types.Event{
		DestModule: destModule,
		Type: &types.Event_Tester{
			Tester: &types1.Tester{},
		},
	}
}

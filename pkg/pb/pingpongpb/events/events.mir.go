// Code generated by Mir codegen. DO NOT EDIT.

package pingpongpbevents

import (
	types "github.com/matejpavlovic/mir/pkg/pb/eventpb/types"
	types1 "github.com/matejpavlovic/mir/pkg/pb/pingpongpb/types"
	stdtypes "github.com/matejpavlovic/mir/stdtypes"
)

func PingTime(destModule stdtypes.ModuleID) *types.Event {
	return &types.Event{
		DestModule: destModule,
		Type: &types.Event_PingPong{
			PingPong: &types1.Event{
				Type: &types1.Event_PingTime{
					PingTime: &types1.PingTime{},
				},
			},
		},
	}
}

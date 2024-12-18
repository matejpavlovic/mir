// Code generated by Mir codegen. DO NOT EDIT.

package bcbpbevents

import (
	types1 "github.com/matejpavlovic/mir/pkg/pb/bcbpb/types"
	types "github.com/matejpavlovic/mir/pkg/pb/eventpb/types"
	stdtypes "github.com/matejpavlovic/mir/stdtypes"
)

func BroadcastRequest(destModule stdtypes.ModuleID, data []uint8) *types.Event {
	return &types.Event{
		DestModule: destModule,
		Type: &types.Event_Bcb{
			Bcb: &types1.Event{
				Type: &types1.Event_Request{
					Request: &types1.BroadcastRequest{
						Data: data,
					},
				},
			},
		},
	}
}

func Deliver(destModule stdtypes.ModuleID, data []uint8) *types.Event {
	return &types.Event{
		DestModule: destModule,
		Type: &types.Event_Bcb{
			Bcb: &types1.Event{
				Type: &types1.Event_Deliver{
					Deliver: &types1.Deliver{
						Data: data,
					},
				},
			},
		},
	}
}

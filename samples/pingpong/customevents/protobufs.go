package customevents

import (
	"github.com/matejpavlovic/mir/pkg/pb/eventpb"
	"github.com/matejpavlovic/mir/pkg/pb/pingpongpb"
	t "github.com/matejpavlovic/mir/stdtypes"
)

func Event(destModule t.ModuleID, ppEvent *pingpongpb.Event) *eventpb.Event {
	return &eventpb.Event{
		DestModule: destModule.String(),
		Type: &eventpb.Event_PingPong{
			PingPong: ppEvent,
		},
	}
}

func PingTimeEvent(destModule t.ModuleID) *eventpb.Event {
	return Event(
		destModule,
		&pingpongpb.Event{Type: &pingpongpb.Event_PingTime{
			PingTime: &pingpongpb.PingTime{},
		}},
	)
}

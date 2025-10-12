package events

import (
	"encoding/json"
	"fmt"
	"github.com/matejpavlovic/mir/stdtypes"
)

type SimEvent struct {
	SrcModule  stdtypes.ModuleID
	DestModule stdtypes.ModuleID

	Timestamp int64
}

func NewSimEvent(srcModule stdtypes.ModuleID, destModule stdtypes.ModuleID, timestamp int64) *SimEvent {
	return &SimEvent{
		SrcModule:  srcModule,
		DestModule: destModule,
		Timestamp:  timestamp,
	}
}

func (e *SimEvent) Src() stdtypes.ModuleID {
	return e.SrcModule
}

func (e *SimEvent) NewSrc(newSrc stdtypes.ModuleID) stdtypes.Event {
	newE := *e
	e.SrcModule = newSrc
	return &newE
}

func (e *SimEvent) Dest() stdtypes.ModuleID {
	return e.DestModule
}

func (e *SimEvent) NewDest(newDest stdtypes.ModuleID) stdtypes.Event {
	newE := *e
	e.DestModule = newDest
	return &newE
}

func (e *SimEvent) ToBytes() ([]byte, error) {
	panic("Serialization not implemented for SimEvent")
}

func (e *SimEvent) ToString() string {
	data, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("unmarshalableEvent(%+v)", e)
	}
	return string(data)
}

type NewBlockEvent struct {
	SimEvent
	Height int64
}

func NewNewBlockEvent(
	srcModule stdtypes.ModuleID,
	destModule stdtypes.ModuleID,
	timestamp int64,
	height int64,
) *NewBlockEvent {
	return &NewBlockEvent{
		SimEvent: *NewSimEvent(srcModule, destModule, timestamp),
		Height:   height,
	}
}

package events

import (
	"github.com/matejpavlovic/mir/pkg/util/priorityqueue"
	"github.com/matejpavlovic/mir/stdtypes"
)

type SimEvent interface {
	stdtypes.Event
	priorityqueue.WithPriority
	Timestamp() int64
}

type SimEventImpl struct {
	SrcModule  stdtypes.ModuleID
	DestModule stdtypes.ModuleID

	Ts int64
}

func NewSimEvent(srcModule stdtypes.ModuleID, destModule stdtypes.ModuleID, timestamp int64) *SimEventImpl {
	return &SimEventImpl{
		SrcModule:  srcModule,
		DestModule: destModule,
		Ts:         timestamp,
	}
}

func (e *SimEventImpl) Src() stdtypes.ModuleID {
	return e.SrcModule
}

func (e *SimEventImpl) NewSrc(newSrc stdtypes.ModuleID) stdtypes.Event {
	newE := *e
	e.SrcModule = newSrc
	return &newE
}

func (e *SimEventImpl) Dest() stdtypes.ModuleID {
	return e.DestModule
}

func (e *SimEventImpl) NewDest(newDest stdtypes.ModuleID) stdtypes.Event {
	newE := *e
	e.DestModule = newDest
	return &newE
}

func (e *SimEventImpl) ToBytes() ([]byte, error) {
	// We do not implement this method on purpose so an actual event does not accidentally "inherit" it,
	// forgetting to serialize other fields that extend SimEventImpl.
	panic("SimEventImpl cannot be serialized on its own.")
}

func (e *SimEventImpl) ToString() string {
	// We do not implement this method on purpose so an actual event does not accidentally "inherit" it,
	// forgetting to include other fields that extend SimEventImpl.
	panic("SimEventImpl cannot be converted to a string on its own.")
}

func (e *SimEventImpl) Priority() int64 {
	return e.Ts
}

func (e *SimEventImpl) Hash() []byte {
	// We do not implement this method on purpose so an actual event does not accidentally "inherit" it,
	// forgetting to hash other fields that extend SimEventImpl.
	panic("SimEventImpl cannot be hashed on its own.")
}

func (e *SimEventImpl) Timestamp() int64 {
	return e.Ts
}

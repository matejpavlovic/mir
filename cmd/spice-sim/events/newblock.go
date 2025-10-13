package events

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/fxamacker/cbor/v2"
	"github.com/matejpavlovic/mir/stdtypes"
)

type NewBlockEvent struct {
	SimEventImpl
	Height int64
}

func NewNewBlockEvent(
	srcModule stdtypes.ModuleID,
	destModule stdtypes.ModuleID,
	timestamp int64,
	height int64,
) *NewBlockEvent {
	return &NewBlockEvent{
		SimEventImpl: *NewSimEvent(srcModule, destModule, timestamp),
		Height:       height,
	}
}

func (e *NewBlockEvent) ToBytes() ([]byte, error) {
	encMode, err := cbor.CoreDetEncOptions().EncMode()
	if err != nil {
		return nil, err
	}
	data, err := encMode.Marshal(*e)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (e *NewBlockEvent) ToString() string {
	data, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("unmarshalableEvent(%+v)", e)
	}
	return string(data)
}

func (e *NewBlockEvent) Hash() []byte {
	data, err := e.ToBytes()
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(data)
	return hash[:]
}

package events

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/fxamacker/cbor/v2"
	"github.com/matejpavlovic/mir/cmd/spice-sim/types"
	"github.com/matejpavlovic/mir/stdtypes"
)

type StateWitnessEvent struct {
	SimEventImpl
	ChunkID types.ChunkID
}

func NewStateWitnessEvent(
	srcModule stdtypes.ModuleID,
	destModule stdtypes.ModuleID,
	timestamp int64,
	chunkID types.ChunkID,
) *StateWitnessEvent {
	return &StateWitnessEvent{
		SimEventImpl: *NewSimEvent(srcModule, destModule, timestamp),
		ChunkID:      chunkID,
	}
}

func (e *StateWitnessEvent) ToBytes() ([]byte, error) {
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

func (e *StateWitnessEvent) ToString() string {
	data, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("unmarshalableEvent(%+v)", e)
	}
	return string(data)
}

func (e *StateWitnessEvent) Hash() []byte {
	data, err := e.ToBytes()
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(data)
	return hash[:]
}

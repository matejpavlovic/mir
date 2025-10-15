package events

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/fxamacker/cbor/v2"
	"github.com/matejpavlovic/mir/cmd/spice-sim/types"
	"github.com/matejpavlovic/mir/stdtypes"
)

type ChunkPartEvent struct {
	SimEventImpl
	PartID types.ChunkPartID
}

func NewChunkPartEvent(
	srcModule stdtypes.ModuleID,
	destModule stdtypes.ModuleID,
	timestamp int64,
	partID types.ChunkPartID,
) *ChunkPartEvent {
	return &ChunkPartEvent{
		SimEventImpl: *NewSimEvent(srcModule, destModule, timestamp),
		PartID:       partID,
	}
}

func (e *ChunkPartEvent) ToBytes() ([]byte, error) {
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

func (e *ChunkPartEvent) ToString() string {
	data, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("unmarshalableEvent(%+v)", e)
	}
	return string(data)
}

func (e *ChunkPartEvent) Hash() []byte {
	data, err := e.ToBytes()
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(data)
	return hash[:]
}

package types

import (
	"github.com/matejpavlovic/mir/stdtypes"
)

type StateCert struct {
	ChunkID ChunkID
	Signers []stdtypes.ModuleID
}

func NewStateCert(chunkID ChunkID, signers []stdtypes.ModuleID) StateCert {
	return StateCert{
		ChunkID: chunkID,
		Signers: signers,
	}
}

func (ac *StateCert) ID() ChunkID {
	return ac.ChunkID
}

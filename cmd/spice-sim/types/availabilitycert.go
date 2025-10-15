package types

import (
	"github.com/matejpavlovic/mir/stdtypes"
)

type AvailabilityCert struct {
	ChunkID ChunkID
	Signers []stdtypes.ModuleID
}

func NewAvailabilityCert(chunkID ChunkID, signers []stdtypes.ModuleID) AvailabilityCert {
	return AvailabilityCert{
		ChunkID: chunkID,
		Signers: signers,
	}
}

func (ac *AvailabilityCert) ID() ChunkID {
	return ac.ChunkID
}

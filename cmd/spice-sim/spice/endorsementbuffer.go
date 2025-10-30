package spice

import (
	"fmt"

	"github.com/matejpavlovic/mir/cmd/spice-sim/events"
	"github.com/matejpavlovic/mir/cmd/spice-sim/types"
	"github.com/matejpavlovic/mir/pkg/util/maputil"
	"github.com/matejpavlovic/mir/stdtypes"
)

type EndorsementBuffer map[string]chunkEndorsements

func NewEndorsementBuffer() EndorsementBuffer {
	return make(EndorsementBuffer)
}

func (eb *EndorsementBuffer) Add(e *events.EndorseExecResultEvent) {
	chunkKey := e.ChunkID.String()
	endorsements, ok := (*eb)[chunkKey]
	if !ok {
		endorsements = newChunkEndorsements(e.ChunkID)
		(*eb)[chunkKey] = endorsements
	}
	endorsements.Add(e.SrcModule)
}

func (eb *EndorsementBuffer) NumEndorsements(chunkID types.ChunkID) int {
	endorsements, ok := (*eb)[chunkID.String()]
	if !ok {
		return 0
	}
	return endorsements.Len()
}

// StateCertified returns true if all state associated with the given blockID (all shards)
// have been certified through enough endorsements.
func (eb *EndorsementBuffer) StateCertified(blockID types.BlockID, numShards int64, quorum int64) bool {
	for s := int64(0); s < numShards; s++ {
		chunkID := types.ChunkID{BlockID: blockID, Shard: s}
		if int64(eb.NumEndorsements(chunkID)) < quorum {
			return false
		}
	}
	return true
}

func (eb *EndorsementBuffer) Certificate(chunkID types.ChunkID) types.StateCert {
	endorsements, ok := (*eb)[chunkID.String()]
	if !ok {
		panic(fmt.Sprintf("no endorsements to certify %s", chunkID.String()))
	}
	return endorsements.Certificate()
}

type chunkEndorsements struct {
	chunkID types.ChunkID
	signers map[stdtypes.ModuleID]struct{}
}

func newChunkEndorsements(chunkID types.ChunkID) chunkEndorsements {
	return chunkEndorsements{
		chunkID: chunkID,
		signers: make(map[stdtypes.ModuleID]struct{}),
	}
}

func (eb *chunkEndorsements) Len() int {
	return len(eb.signers)
}

func (eb *chunkEndorsements) Add(signerID stdtypes.ModuleID) {
	eb.signers[signerID] = struct{}{}
}

func (eb *chunkEndorsements) Certificate() types.StateCert {
	return types.NewStateCert(eb.chunkID, maputil.GetSortedKeys(eb.signers))
}

package types

import "fmt"

type ChunkPartID struct {
	ChunkID ChunkID
	Index   int64
}

func (cpid ChunkPartID) String() string {
	return fmt.Sprintf("%s.%d", cpid.ChunkID, cpid.Index)
}

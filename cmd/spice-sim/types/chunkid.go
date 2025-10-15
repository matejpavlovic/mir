package types

import "fmt"

type ChunkID struct {
	BlockID BlockID
	Shard   int64
}

func (cid ChunkID) String() string {
	return fmt.Sprintf("%sc%d", cid.BlockID, cid.Shard)
}

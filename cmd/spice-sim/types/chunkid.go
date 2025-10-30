package types

import "fmt"

type ChunkID struct {
	// ID of the block with which the chunk is associated.
	// The block needs to exist before the chunk can be created.
	BlockID BlockID

	Shard int64
}

func (cid ChunkID) String() string {
	return fmt.Sprintf("%sc%d", cid.BlockID, cid.Shard)
}

package types

import (
	"bytes"
	"fmt"
)

type BlockID struct {
	Height int64
	Hash   []byte
}

func NewBlockID(height int64, hash []byte) BlockID {
	newHash := make([]byte, len(hash))
	copy(newHash, hash)
	return BlockID{
		Height: height,
		Hash:   newHash,
	}
}

func (bid BlockID) String() string {
	if len(bid.Hash) >= 2 {
		return fmt.Sprintf("b%d(%x)", bid.Height, bid.Hash[:2])
	} else {
		return fmt.Sprintf("b%d(____)", bid.Height)
	}
}

func (bid BlockID) Equals(other BlockID) bool {
	return bid.Height == other.Height && bytes.Equal(bid.Hash, other.Hash)
}

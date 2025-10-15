package types

import "fmt"

type BlockID struct {
	Height     int64
	ParentHash []byte
}

func NewBlockID(height int64, parentHash []byte) BlockID {
	newParentHash := make([]byte, len(parentHash))
	copy(newParentHash, parentHash)
	return BlockID{
		Height:     height,
		ParentHash: newParentHash,
	}
}

func (bid BlockID) String() string {
	if len(bid.ParentHash) >= 2 {
		return fmt.Sprintf("b%d(%x)", bid.Height, bid.ParentHash[:2])
	} else {
		return fmt.Sprintf("b%d(____)", bid.Height)
	}
}

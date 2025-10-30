package types

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Height            int64
	ParentHash        []byte
	AvailabilityCerts []AvailabilityCert
	StateCerts        []StateCert
}

func NewBlock(height int64, parentHash []byte, availabilityCerts []AvailabilityCert, stateCerts []StateCert) *Block {
	return &Block{
		Height:            height,
		ParentHash:        parentHash,
		AvailabilityCerts: availabilityCerts,
		StateCerts:        stateCerts,
	}
}

func GenesisBlock() *Block {
	return NewBlock(0, []byte{}, []AvailabilityCert{}, []StateCert{})
}

func (b *Block) Hash() []byte {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v|", b.Height)))
	h.Write(b.ParentHash)
	for _, cert := range b.AvailabilityCerts {
		h.Write([]byte(cert.ID().String()))
		for _, signer := range cert.Signers {
			h.Write([]byte(signer.String()))
		}
	}
	for _, cert := range b.StateCerts {
		h.Write([]byte(cert.ID().String()))
		for _, signer := range cert.Signers {
			h.Write([]byte(signer.String()))
		}
	}
	return h.Sum(nil)
}

func (b *Block) Id() BlockID {
	return NewBlockID(b.Height, b.Hash())
}

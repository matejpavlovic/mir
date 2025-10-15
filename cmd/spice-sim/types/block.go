package types

import "crypto/sha256"

type Block struct {
	Id                BlockID
	AvailabilityCerts []AvailabilityCert
}

func NewBlock(id BlockID, availabilityCerts []AvailabilityCert) *Block {
	return &Block{
		Id:                id,
		AvailabilityCerts: availabilityCerts,
	}
}

func (b *Block) Hash() []byte {
	h := sha256.New()
	h.Write([]byte(b.Id.String()))
	for _, cert := range b.AvailabilityCerts {
		h.Write([]byte(cert.ID().String()))
		for _, signer := range cert.Signers {
			h.Write([]byte(signer.String()))
		}
	}
	return h.Sum(nil)
}

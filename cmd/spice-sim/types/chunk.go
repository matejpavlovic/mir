package types

type Chunk struct {
	Id ChunkID
}

func NewChunk(id ChunkID) Chunk {
	return Chunk{id}
}

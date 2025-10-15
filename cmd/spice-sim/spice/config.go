package spice

type Config struct {
	BlockTime                 int
	NumBlockProducers         int
	NumShards                 int
	NumChunkProducersPerShard int
	ChunkDispersalDelay       int
	NumDataOwners             int
	NumChunkDataParts         int
	ChunkDataStatementQuorum  int
	StatementSubmissionDelay  int
	// TODO: Complete this.
}

func DefaultConfig() Config {
	return Config{
		BlockTime:                 400,
		NumBlockProducers:         1,
		NumShards:                 1,
		NumChunkProducersPerShard: 2,
		ChunkDispersalDelay:       300,
		NumDataOwners:             8,
		NumChunkDataParts:         4,
		ChunkDataStatementQuorum:  3,
		StatementSubmissionDelay:  200,
	}
}

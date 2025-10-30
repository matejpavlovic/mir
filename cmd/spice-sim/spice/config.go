package spice

type Config struct {
	BlockTime                     int64
	NumBlockProducers             int64
	NumShards                     int64
	NumChunkProducersPerShard     int64
	ChunkDispersalDelay           int64
	NumDataOwners                 int64
	NumChunkDataParts             int64
	ChunkDataStatementQuorum      int64
	StatementSubmissionDelay      int64
	NumReplicasPerShard           int64
	NumValidators                 int64
	ValidatorSampleSize           int64
	StateCertQuorum               int64
	ExecutionDelay                int64
	ReceiptTransmissionDelay      int64
	StateWitnessTransmissionDelay int64
	StateWitnessValidationDelay   int64

	// TODO: Complete this.
}

func DefaultConfig() Config {
	numShards := int64(2)

	return Config{

		BlockTime: 400,

		NumBlockProducers: 1,

		NumShards: numShards,

		NumChunkProducersPerShard: 2,

		ChunkDispersalDelay: 300,

		NumDataOwners: 8,

		NumChunkDataParts: 4,

		ChunkDataStatementQuorum: 3,

		StatementSubmissionDelay: 200,

		NumReplicasPerShard: 1,

		NumValidators: 4 * numShards,

		ValidatorSampleSize: 4,

		// Number of state endorsements required for a certificate.
		// Must not be more than ValidatorSampleSize.
		StateCertQuorum: 3,

		ExecutionDelay: 50,

		ReceiptTransmissionDelay: 100,

		StateWitnessTransmissionDelay: 400,

		StateWitnessValidationDelay: 100,
	}
}

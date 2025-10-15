package spice

import (
	"fmt"
	"github.com/matejpavlovic/mir/stdtypes"
)

type CoreState struct {
	config            Config
	availabilityCerts map[string]struct{}
}

func NewCoreState(config Config) CoreState {
	return CoreState{
		config:            config,
		availabilityCerts: make(map[string]struct{}),
	}
}

func (cs *CoreState) BlockProducerIDs() []stdtypes.ModuleID {
	ids := make([]stdtypes.ModuleID, 0, cs.config.NumBlockProducers)
	for i := 0; i < cs.config.NumBlockProducers; i++ {
		ids = append(ids, stdtypes.ModuleID(fmt.Sprintf("block-producer-%d", i)))
	}
	return ids
}

func (cs *CoreState) ChunkProducerIDs(shard int64) []stdtypes.ModuleID {
	ids := make([]stdtypes.ModuleID, 0, cs.config.NumChunkProducersPerShard)
	for i := 0; i < cs.config.NumChunkProducersPerShard; i++ {
		ids = append(ids, stdtypes.ModuleID(fmt.Sprintf("chunk-producer-%d-%d", shard, i)))
	}
	return ids
}

func (cs *CoreState) ChunkProducerIDsAllShards() []stdtypes.ModuleID {
	ids := make([]stdtypes.ModuleID, 0, cs.config.NumChunkProducersPerShard*cs.config.NumShards)
	for shard := 0; shard < cs.config.NumShards; shard++ {
		for i := 0; i < cs.config.NumChunkProducersPerShard; i++ {
			ids = append(ids, stdtypes.ModuleID(fmt.Sprintf("chunk-producer-%d-%d", shard, i)))
		}
	}
	return ids
}

func (cs *CoreState) DataOwnerIDs() []stdtypes.ModuleID {
	ids := make([]stdtypes.ModuleID, 0, cs.config.NumDataOwners)
	for i := 0; i < cs.config.NumDataOwners; i++ {
		ids = append(ids, stdtypes.ModuleID(fmt.Sprintf("data-owner-%d", i)))
	}
	return ids
}

func (cs *CoreState) BlockProducerAt(height int64) stdtypes.ModuleID {
	return stdtypes.ModuleID(fmt.Sprintf("block-producer-%d", height%int64(cs.config.NumBlockProducers)))
}

func (cs *CoreState) ChunkProducerAt(height int64, shard int64) stdtypes.ModuleID {
	return stdtypes.ModuleID(fmt.Sprintf(
		"chunk-producer-%d-%d",
		shard,
		height%int64(cs.config.NumChunkProducersPerShard),
	))
}

func (cs *CoreState) DataOwnersAt(height int64, shard int64) []stdtypes.ModuleID {
	// We use a round-robin assignment of data parts to data owners, irrespectively of the shards.
	// For each chunk, we assign the next cs.config.NumChunkDataParts data owners in line.
	//         < ---------------------- data parts produced until now ---------------------- >
	//         < -- shards in previous and this block -- >
	//         < - shards in previous blocks - >
	index := ((height*int64(cs.config.NumShards) + shard) * int64(cs.config.NumChunkDataParts)) % int64(cs.config.NumDataOwners)
	dataOwnerIDs := make([]stdtypes.ModuleID, 0, cs.config.NumChunkDataParts)
	for i := 0; i < cs.config.NumChunkDataParts; i++ {
		dataOwnerIDs = append(dataOwnerIDs, stdtypes.ModuleID(fmt.Sprintf("data-owner-%d", index)))
		index = (index + 1) % int64(cs.config.NumDataOwners)
	}
	return dataOwnerIDs
}

func (cs *CoreState) AddAvailabilityCert(cert string) {
	cs.availabilityCerts[cert] = struct{}{}
}

func (cs *CoreState) ChunkAvailable(chunkID string) bool {
	_, ok := cs.availabilityCerts[chunkID]
	return ok
}

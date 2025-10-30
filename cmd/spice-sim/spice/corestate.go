package spice

import (
	"bytes"
	"fmt"

	"github.com/matejpavlovic/mir/cmd/spice-sim/types"
	"github.com/matejpavlovic/mir/stdtypes"
)

type CoreState struct {
	config            Config
	availabilityCerts map[string]struct{}
	stateCerts        map[string]struct{}
	canonicalChain    []*types.Block
}

func NewCoreState(config Config) CoreState {
	return CoreState{
		config:            config,
		availabilityCerts: make(map[string]struct{}),
		stateCerts:        make(map[string]struct{}),
		canonicalChain:    make([]*types.Block, 0),
	}
}

func (cs *CoreState) BlockProducerIDs() []stdtypes.ModuleID {
	ids := make([]stdtypes.ModuleID, 0, cs.config.NumBlockProducers)
	for i := int64(0); i < cs.config.NumBlockProducers; i++ {
		ids = append(ids, stdtypes.ModuleID(fmt.Sprintf("block-producer-%d", i)))
	}
	return ids
}

func (cs *CoreState) ChunkProducerIDs(shard int64) []stdtypes.ModuleID {
	ids := make([]stdtypes.ModuleID, 0, cs.config.NumChunkProducersPerShard)
	for i := int64(0); i < cs.config.NumChunkProducersPerShard; i++ {
		ids = append(ids, stdtypes.ModuleID(fmt.Sprintf("chunk-producer-%d-%d", shard, i)))
	}
	return ids
}

func (cs *CoreState) ChunkProducerIDsAllShards() []stdtypes.ModuleID {
	ids := make([]stdtypes.ModuleID, 0, cs.config.NumChunkProducersPerShard*cs.config.NumShards)
	for shard := int64(0); shard < cs.config.NumShards; shard++ {
		ids = append(ids, cs.ChunkProducerIDs(int64(shard))...)
	}
	return ids
}

func (cs *CoreState) ReplicaIDs(shard int64) []stdtypes.ModuleID {
	ids := make([]stdtypes.ModuleID, 0, cs.config.NumReplicasPerShard)
	for i := int64(0); i < cs.config.NumReplicasPerShard; i++ {
		ids = append(ids, stdtypes.ModuleID(fmt.Sprintf("replica-%d-%d", shard, i)))
	}
	return ids
}

func (cs *CoreState) ReplicaIDsAllShards() []stdtypes.ModuleID {
	ids := make([]stdtypes.ModuleID, 0, cs.config.NumReplicasPerShard*cs.config.NumShards)
	for shard := int64(0); shard < cs.config.NumShards; shard++ {
		ids = append(ids, cs.ReplicaIDs(shard)...)
	}
	return ids
}

func (cs *CoreState) DataOwnerIDs() []stdtypes.ModuleID {
	ids := make([]stdtypes.ModuleID, 0, cs.config.NumDataOwners)
	for i := int64(0); i < cs.config.NumDataOwners; i++ {
		ids = append(ids, stdtypes.ModuleID(fmt.Sprintf("data-owner-%d", i)))
	}
	return ids
}

func (cs *CoreState) BlockProducerAt(height int64) stdtypes.ModuleID {
	return stdtypes.ModuleID(fmt.Sprintf("block-producer-%d", height%cs.config.NumBlockProducers))
}

func (cs *CoreState) ChunkProducerAt(height int64, shard int64) stdtypes.ModuleID {
	return stdtypes.ModuleID(fmt.Sprintf(
		"chunk-producer-%d-%d",
		shard,
		height%cs.config.NumChunkProducersPerShard,
	))
}

func (cs *CoreState) DataOwnersAt(height int64, shard int64) []stdtypes.ModuleID {
	// We use a round-robin assignment of data parts to data owners, irrespectively of the shards.
	// For each chunk, we assign the next cs.config.NumChunkDataParts data owners in line.
	//         < ---------------------- data parts produced until now ---------------------- >
	//         < -- shards in previous and this block -- >
	//         < - shards in previous blocks - >
	index := ((height*cs.config.NumShards + shard) * cs.config.NumChunkDataParts) % cs.config.NumDataOwners
	dataOwnerIDs := make([]stdtypes.ModuleID, 0, cs.config.NumChunkDataParts)
	for i := int64(0); i < cs.config.NumChunkDataParts; i++ {
		dataOwnerIDs = append(dataOwnerIDs, stdtypes.ModuleID(fmt.Sprintf("data-owner-%d", index)))
		index = (index + 1) % cs.config.NumDataOwners
	}
	return dataOwnerIDs
}

func (cs *CoreState) ValidatorIDs() []stdtypes.ModuleID {
	ids := make([]stdtypes.ModuleID, 0, cs.config.NumValidators)
	for i := int64(0); i < cs.config.NumValidators; i++ {
		ids = append(ids, stdtypes.ModuleID(fmt.Sprintf("validator-%d", i)))
	}
	return ids
}

func (cs *CoreState) ValidatorsAt(height int64, shard int64) []stdtypes.ModuleID {
	return roundRobin(height, shard, cs.config.NumShards, cs.config.ValidatorSampleSize, cs.config.NumValidators, "validator-")
}

func (cs *CoreState) ApplyBlock(block *types.Block) {

	// Don't look for previous head when applying genesis.
	if len(cs.canonicalChain) != 0 {
		oldHead := cs.canonicalChain[len(cs.canonicalChain)-1]
		if !bytes.Equal(block.ParentHash, oldHead.Hash()) {
			panic(fmt.Sprintf(
				"applying block out of order not implemented. last block (%s) has hash: %x, received block: %s",
				oldHead.Id,
				oldHead.Hash(),
				block.Id,
			))
		}
	}

	cs.canonicalChain = append(cs.canonicalChain, block)

	for _, cert := range block.AvailabilityCerts {
		cs.addAvailabilityCert(cert)
	}
	for _, cert := range block.StateCerts {
		cs.addStateCert(cert)
	}
}

func (cs *CoreState) NextCanonicalBlock(blockID types.BlockID) *types.Block {

	// Find the given block.
	var index int
	for index = len(cs.canonicalChain) - 1; index >= 0; index-- {
		if bytes.Equal(cs.canonicalChain[index].Hash(), blockID.Hash) {
			break
		}
	}

	// If there is no successor yet
	if index == len(cs.canonicalChain)-1 {
		return nil
	}

	// If block was not found
	if !bytes.Equal(cs.canonicalChain[index].Hash(), blockID.Hash) {
		panic("block not found in canonical chain")
	}

	return cs.canonicalChain[index+1]
}

func (cs *CoreState) addAvailabilityCert(cert types.AvailabilityCert) {
	cs.availabilityCerts[cert.ChunkID.String()] = struct{}{}
}

func (cs *CoreState) addStateCert(cert types.StateCert) {
	cs.stateCerts[cert.ChunkID.String()] = struct{}{}
}

func (cs *CoreState) ChunkAvailable(chunkID types.ChunkID) bool {
	_, ok := cs.availabilityCerts[chunkID.String()]
	return ok
}

func (cs *CoreState) StateCertified(chunkID types.ChunkID) bool {
	_, ok := cs.stateCerts[chunkID.String()]
	return ok
}

func roundRobin(height int64, shard int64, numShards int64, numItems int64, poolSize int64, prefix string) []stdtypes.ModuleID {
	//         < ---------------------- data parts produced until now ---------------------- >
	//         < -------- items in previous and this block -------- >
	//         < - batches in previous and this block - >
	//         < - batches in previous blocks - >
	index := ((height /*   */ * /*   */ numShards + shard) * numItems) % poolSize
	moduleIDs := make([]stdtypes.ModuleID, 0, numItems)
	for i := int64(0); i < numItems; i++ {
		moduleIDs = append(moduleIDs, stdtypes.ModuleID(fmt.Sprintf("%s%d", prefix, index)))
		index = (index + 1) % poolSize
	}
	return moduleIDs
}

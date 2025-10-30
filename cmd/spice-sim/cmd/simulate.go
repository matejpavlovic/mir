package cmd

import (
	"context"
	"fmt"

	"github.com/matejpavlovic/mir"
	"github.com/matejpavlovic/mir/cmd/spice-sim/spice"
	"github.com/matejpavlovic/mir/pkg/modules"
	"github.com/matejpavlovic/mir/stdtypes"
	"github.com/spf13/cobra"
)

var (
	nodeCmd = &cobra.Command{
		Use:   "simulate",
		Short: "Runa a simulation",
		RunE: func(_ *cobra.Command, _ []string) error {
			return runSimulation()
		},
	}
)

func init() {
	rootCmd.AddCommand(nodeCmd)
}

func runSimulation() error {
	fmt.Println("Running simulation...")

	config := spice.DefaultConfig()
	coreState := spice.NewCoreState(config)

	spiceModules := map[stdtypes.ModuleID]modules.Module{}

	// Create block producers.
	for _, blockProducerID := range coreState.BlockProducerIDs() {
		spiceModules[blockProducerID] = spice.NewBlockProducer(blockProducerID, coreState)
	}

	// Create chunk producers.
	for shard := int64(0); shard < int64(config.NumShards); shard++ {
		for _, chunkProducerID := range coreState.ChunkProducerIDs(shard) {
			spiceModules[chunkProducerID] = spice.NewChunkProducer(chunkProducerID, shard, coreState)
		}
	}

	// Create data owners.
	for _, dataOwnerID := range coreState.DataOwnerIDs() {
		spiceModules[dataOwnerID] = spice.NewDataOwner(dataOwnerID, coreState)
	}

	// Create replicas
	for shard := int64(0); shard < int64(config.NumShards); shard++ {
		for _, replicaID := range coreState.ReplicaIDs(shard) {
			spiceModules[replicaID] = spice.NewReplica(replicaID, shard, coreState)
		}
	}

	// Create Validators
	for _, validatorID := range coreState.ValidatorIDs() {
		spiceModules[validatorID] = spice.NewValidator(validatorID, coreState)
	}

	mirNode, err := mir.NewNode("0", mir.DefaultNodeConfig(), spiceModules, nil)
	if err != nil {
		return err
	}

	return mirNode.RunFor(context.Background(), 10000)
}

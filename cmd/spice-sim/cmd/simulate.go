package cmd

import (
	"context"
	"fmt"
	"github.com/matejpavlovic/mir"
	"github.com/matejpavlovic/mir/cmd/spice-sim/noderoles"
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

	mirNode, err := mir.NewNode("0", mir.DefaultNodeConfig(), map[stdtypes.ModuleID]modules.Module{
		"block-producer-0": noderoles.NewBlockProducer("block-producer-0"),
	}, nil)
	if err != nil {
		return err
	}

	return mirNode.Run(context.Background())
}

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	verbose bool

	rootCmd = &cobra.Command{
		Use:   "spice-sim",
		Short: "Simulation tool for SPICE",
		Long: "SPICE simulation tool for simulating high-level interactions between" +
			"different kinds of nodes in SPICE. It can be used to understand the interactions" +
			"and estimate performance of SPICE.",
	}
)

func Execute() {
	start := time.Now()
	err := rootCmd.Execute()
	fmt.Printf("Time elapsed: %v\n", time.Since(start))
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose mode")
}

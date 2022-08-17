package main

import (
	"context"
	"crypto"
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/filecoin-project/mir"
	"github.com/filecoin-project/mir/pkg/availability/multisigcollector"
	mirCrypto "github.com/filecoin-project/mir/pkg/crypto"
	"github.com/filecoin-project/mir/pkg/logging"
	"github.com/filecoin-project/mir/pkg/mempool/simplemempool"
	"github.com/filecoin-project/mir/pkg/modules"
	"github.com/filecoin-project/mir/pkg/net/grpc"
	t "github.com/filecoin-project/mir/pkg/types"
	grpctools "github.com/filecoin-project/mir/pkg/util/grpc"
)

const (

	// Base port number for the nodes to listen to messages from each other.
	// The nodes will listen on ports starting from nodeBasePort through nodeBasePort+3.
	nodeBasePort = 10000

	// The number of nodes participating in the chat.
	nodeNumber = 4
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	args := parseArgs(os.Args)

	// IDs of nodes that are part of the system.
	// This example uses a static configuration of nodeNumber nodes.
	nodeIDs := make([]t.NodeID, nodeNumber)
	for i := 0; i < nodeNumber; i++ {
		nodeIDs[i] = t.NewNodeIDFromInt(i)
	}

	nodeAddrs := make(map[t.NodeID]t.NodeAddress)
	for i := range nodeIDs {
		nodeAddrs[t.NewNodeIDFromInt(i)] = t.NodeAddress(grpctools.NewDummyMultiaddr(i + nodeBasePort))
	}

	transport, err := grpc.NewTransport(args.OwnID, nodeAddrs[args.OwnID], logging.NilLogger)
	if err != nil {
		return fmt.Errorf("failed to get network transport %w", err)
	}
	if err := transport.Start(); err != nil {
		return fmt.Errorf("could not start network transport: %w", err)
	}
	transport.Connect(context.Background(), nodeAddrs)

	mempool := simplemempool.NewModule(
		&simplemempool.ModuleConfig{
			Self:   "mempool",
			Hasher: "hasher",
		},
		&simplemempool.ModuleParams{
			MaxTransactionsInBatch: 10,
		},
	)

	availability, err := multisigcollector.NewModule(
		&multisigcollector.ModuleConfig{
			Self:    "availability",
			Mempool: "mempool",
			Net:     "net",
			Crypto:  "crypto",
		},
		&multisigcollector.ModuleParams{
			InstanceUID: []byte("testing instance"),
			AllNodes:    nodeIDs,
			F:           (len(nodeIDs) - 1) / 2,
		},
		args.OwnID,
	)
	if err != nil {
		return err
	}

	// control module reads the user input from the console and processes it.
	control := newControlModule()

	m := map[t.ModuleID]modules.Module{
		"net":          transport,
		"mempool":      mempool,
		"hasher":       mirCrypto.NewHasher(crypto.SHA256),
		"crypto":       mirCrypto.New(&mirCrypto.DummyCrypto{DummySig: []byte{0}}),
		"availability": availability,
		"control":      control,
	}

	// create a Mir node
	node, err := mir.NewNode("client", &mir.NodeConfig{Logger: logging.NilLogger}, m, nil, nil)
	if err != nil {
		return fmt.Errorf("error creating a Mir node: %w", err)
	}

	// run the node
	err = node.Run(context.Background())
	if err != nil {
		return fmt.Errorf("error running node: %w", err)
	}

	return nil
}

// Parses the command-line arguments and returns them in a params struct.
func parseArgs(args []string) *parsedArgs {
	app := kingpin.New("chat-demo", "Small chat application to demonstrate the usage of the Mir library.")
	verbose := app.Flag("verbose", "Verbose mode.").Short('v').Bool()
	ownID := app.Arg("id", "ID of this node").Required().String()

	if _, err := app.Parse(args[1:]); err != nil { // Skip args[0], which is the name of the program, not an argument.
		app.FatalUsage("could not parse arguments: %v\n", err)
	}

	return &parsedArgs{
		OwnID:   t.NodeID(*ownID),
		Verbose: *verbose,
	}
}

// parsedArgs represents parsed command-line parameters passed to the program.
type parsedArgs struct {

	// ID of this node.
	// The package github.com/hyperledger-labs/mir/pkg/types defines this and other types used by the library.
	OwnID t.NodeID

	// If set, print verbose output to stdout.
	Verbose bool
}
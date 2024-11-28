package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/matejpavlovic/mir"
	"github.com/matejpavlovic/mir/pkg/logging"
	"github.com/matejpavlovic/mir/pkg/modules"
	"github.com/matejpavlovic/mir/pkg/net/grpc"
	trantorpbtypes "github.com/matejpavlovic/mir/pkg/pb/trantorpb/types"
	"github.com/matejpavlovic/mir/samples/pingpong/customevents"
	"github.com/matejpavlovic/mir/stdmodules/timer"
	"github.com/matejpavlovic/mir/stdtypes"
)

func main() {
	fmt.Println("Starting ping-pong.")

	// Manually create system membership with just 2 nodes.
	membership := &trantorpbtypes.Membership{map[stdtypes.NodeID]*trantorpbtypes.NodeIdentity{ // nolint:govet
		"0": {"0", "/ip4/127.0.0.1/tcp/10000", nil, "1"}, // nolint:govet
		"1": {"1", "/ip4/127.0.0.1/tcp/10001", nil, "1"}, // nolint:govet
	}}

	// Get own ID from command line.
	ownID := stdtypes.NodeID(os.Args[1])

	// Instantiate network trnasport module and establish connections.
	transport, err := grpc.NewTransport(grpc.DefaultParams(), ownID, membership.Nodes[ownID].Addr, logging.ConsoleWarnLogger, nil)
	if err != nil {
		panic(err)
	}
	if err := transport.Start(); err != nil {
		panic(err)
	}
	transport.Connect(membership)

	// Instantiate Mir node.
	node, err := mir.NewNode(
		ownID,
		mir.DefaultNodeConfig(),
		map[stdtypes.ModuleID]modules.Module{
			"transport": transport,
			//"pingpong":  NewPingPong(ownID),
			//"pingpong": lowlevel.NewPingPong(ownID),
			"pingpong": customevents.NewPingPong(ownID),
			"timer":    timer.New(),
		},
		nil,
	)
	if err != nil {
		panic(err)
	}

	// Run the node for 5 seconds.
	nodeError := make(chan error)
	go func() {
		nodeError <- node.Run(context.Background())
	}()
	fmt.Println("Mir node running.")
	time.Sleep(5 * time.Second)

	// Stop the node.
	node.Stop()
	transport.Stop()
	fmt.Printf("Mir node stopped: %v\n", <-nodeError)
}

package deploytest

import (
	"time"

	"github.com/matejpavlovic/mir/pkg/trantor/types"
	t "github.com/matejpavlovic/mir/stdtypes"

	es "github.com/go-errors/errors"

	"github.com/matejpavlovic/mir/pkg/logging"
	"github.com/matejpavlovic/mir/pkg/net"
	trantorpbtypes "github.com/matejpavlovic/mir/pkg/pb/trantorpb/types"
	"github.com/matejpavlovic/mir/pkg/testsim"
)

type LocalTransportLayer interface {
	Link(source t.NodeID) (net.Transport, error)
	Membership() *trantorpbtypes.Membership
	Close()
}

// NewLocalTransportLayer creates an instance of LocalTransportLayer suitable for tests.
// transportType is one of: "sim", "fake", or "grpc".
func NewLocalTransportLayer(sim *Simulation, transportType string, nodeIDsWeight map[t.NodeID]types.VoteWeight, logger logging.Logger) (LocalTransportLayer, error) {
	switch transportType {
	case "sim":
		messageDelayFn := func(_, _ t.NodeID) time.Duration {
			// TODO: Make min and max message delay configurable
			return testsim.RandDuration(sim.Rand, 0, 10*time.Millisecond)
		}
		return NewSimTransport(sim, nodeIDsWeight, messageDelayFn), nil
	case "fake":
		return NewFakeTransport(nodeIDsWeight), nil
	case "grpc":
		return NewLocalGrpcTransport(nodeIDsWeight, logger)
	default:
		return nil, es.Errorf("unexpected transport type: %v", transportType)
	}
}

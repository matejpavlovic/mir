package trantor

import (
	"time"

	"github.com/matejpavlovic/mir/pkg/availability/multisigcollector"
	issconfig "github.com/matejpavlovic/mir/pkg/iss/config"
	"github.com/matejpavlovic/mir/pkg/mempool/simplemempool"
	"github.com/matejpavlovic/mir/pkg/net/grpc"
	trantorpbtypes "github.com/matejpavlovic/mir/pkg/pb/trantorpb/types"
)

type Params struct {
	Mempool      *simplemempool.ModuleParams
	Iss          *issconfig.ModuleParams
	Net          grpc.Params
	Availability multisigcollector.ModuleParams
}

func DefaultParams(initialMembership *trantorpbtypes.Membership) Params {
	return Params{
		Mempool:      simplemempool.DefaultModuleParams(),
		Iss:          issconfig.DefaultParams(initialMembership),
		Net:          grpc.DefaultParams(),
		Availability: multisigcollector.DefaultParamsTemplate(),
	}
}

func (p *Params) AdjustSpeed(maxProposeDelay time.Duration) *Params {
	p.Iss.AdjustSpeed(maxProposeDelay)
	return p
}

package orderers

import (
	"fmt"

	issconfig "github.com/matejpavlovic/mir/pkg/iss/config"
	"github.com/matejpavlovic/mir/pkg/logging"
	"github.com/matejpavlovic/mir/pkg/modules"
	"github.com/matejpavlovic/mir/pkg/orderers/common"
	ordererpbtypes "github.com/matejpavlovic/mir/pkg/pb/ordererpb/types"
	tt "github.com/matejpavlovic/mir/pkg/trantor/types"
	"github.com/matejpavlovic/mir/stdmodules/factory"
	"github.com/matejpavlovic/mir/stdtypes"
)

func Factory(
	mc common.ModuleConfig,
	issParams *issconfig.ModuleParams,
	ownID stdtypes.NodeID,
	logger logging.Logger,
) modules.PassiveModule {
	if logger == nil {
		logger = logging.ConsoleErrorLogger
	}
	return factory.New(
		mc.Self,
		factory.DefaultParams(

			// This function will be called whenever the factory module
			// is asked to create a new instance of the Ordering protocol.
			func(submoduleID stdtypes.ModuleID, params any) (modules.PassiveModule, error) {

				// Crate a copy of basic module config with an adapted ID for the submodule.
				submc := mc
				submc.Self = submoduleID

				// Load parameters from received protobuf
				// TODO: Use a switch statement and check for a serialized form of the parameters.
				p := (*ordererpbtypes.PBFTModule)(params.(*InstanceParams))
				availabilityID := stdtypes.ModuleID(p.AvailabilityId)
				ppvID := stdtypes.ModuleID(p.PpvModuleId)
				submc.Ava = availabilityID
				submc.PPrepValidator = ppvID
				epoch := tt.EpochNr(p.Epoch)
				segment := (*common.Segment)(p.Segment)

				// Create new configuration for this particular orderer instance.
				ordererConfig := newOrdererConfig(issParams, segment.NodeIDs(), epoch)

				//TODO better logging here
				logger := logging.Decorate(logger, "", "submoduleID", submoduleID)

				// Select validity checker

				if ppvID == stdtypes.ModuleID("pprepvalidatorchkp").Then(stdtypes.ModuleID(fmt.Sprintf("%v", epoch))) { // this must change but that's the scope of a different PR
					// TODO: This is a dirty hack! Put (at least the relevant parts of) the configuration in params.
					// Make the agreement on a checkpoint start immediately.
					ordererConfig.MaxProposeDelay = 0
				}

				// Instantiate new protocol instance.
				protocol := NewOrdererModule(
					submc,
					ownID,
					segment,
					ordererConfig,
					logger,
				)

				return protocol, nil
			},
		),
		logger,
	)
}

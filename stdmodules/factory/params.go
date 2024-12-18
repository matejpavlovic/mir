package factory

import (
	"github.com/matejpavlovic/mir/pkg/modules"
	t "github.com/matejpavlovic/mir/stdtypes"
)

const (
	DefaultMsgBufSize = 1024 * 1024 // 1 MB
)

// ModuleGenerator is a function that generates a passive module from a module ID and the parameters.
// TODO: Change the params type from any to stdtypes.Serializable when legacy proto events are not used with factory.
type ModuleGenerator func(id t.ModuleID, params any) (modules.PassiveModule, error)

// ModuleParams contains the parameters for creating a factory module.
type ModuleParams struct {
	Generator  ModuleGenerator
	MsgBufSize int
}

// DefaultParams creates a valid ModuleParams with default values.
func DefaultParams(generator ModuleGenerator) ModuleParams {
	return ModuleParams{
		Generator:  generator,
		MsgBufSize: DefaultMsgBufSize,
	}
}

package externalmodule

import (
	"context"

	"github.com/matejpavlovic/mir/pkg/dsl"
	"github.com/matejpavlovic/mir/pkg/modules"
	"github.com/matejpavlovic/mir/stdevents"
	"github.com/matejpavlovic/mir/stdtypes"
)

// NewProxyModule returns a new module that serves as a local proxy to an external module hosted on a module server.
// The addr parameter specifies the full URL (address and path) of the module at the server.
// The connection between the proxy and the module server is established when the module receives stdevents.Init,
// At which time the server must be running and accepting new connections.
func NewProxyModule(moduleID stdtypes.ModuleID, addr string) modules.PassiveModule {
	m := dsl.NewModule(moduleID)
	var connection *Connection
	ctx := context.Background()
	// TODO: Using a local context here might make the whole Mir node stuck if the connection gets stuck.
	//   There is no way to stop the module's operation from the outside - it can only stop by itself.
	//   This is a more general problem of passive modules. There is no way to force them to stop processing when the
	//   Mir node is shutting down. For most of the passive modules it is not an issue though, as they only locally
	//   process events and are guaranteed to eventually finish. For the proxy module, this is only the case if it can
	//   communicate with its corresponding server.

	// Upon Init, connect to the remote module and relay the Init event to it.
	dsl.UponEvent(m, func(ev *stdevents.Init) error {
		var err error

		// Create connection to module server.
		connection, err = Connect(ctx, addr)
		if err != nil {
			return err
		}

		// Relay Init event to remote module.
		eventsOut, err := connection.Submit(ctx, stdtypes.ListOf(ev))
		if err != nil {
			return err
		}
		dsl.EmitEvents(m, eventsOut)
		return nil
	})

	// Simply relay all events (except for Init, which is handled separately) to the remote module.
	dsl.UponOtherEvent(m, func(ev stdtypes.Event) error {
		eventsOut, err := connection.Submit(ctx, stdtypes.ListOf(ev))
		if err != nil {
			return err
		}
		dsl.EmitEvents(m, eventsOut)
		return nil
	})

	return m
}

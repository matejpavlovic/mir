package externalmodule

import (
	"context"
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/coder/websocket"

	"github.com/filecoin-project/mir/pkg/modules"
	"github.com/filecoin-project/mir/stdevents"
	"github.com/filecoin-project/mir/stdtypes"
)

const (
	ConnActive = iota
	ConnPending
)

// ModuleHandler implements a handler function for an incoming connection at the module server for a PassiveModule.
type ModuleHandler struct {

	// URL path (after the domain) under which this module will be accessible
	path string

	// The module logic
	module modules.PassiveModule

	// Used to make sure there is only a single client talking to the handler.
	// This is needed to prevent concurrent access to the module.
	// The int32 type is somewhat arbitrary - it only needs to be supported by the CompareAndSwap
	// family of functions in the atomic package.
	connectionStatus int32
}

// NewHandler allocates and returns a pointer to a new ModuleHandler.
func NewHandler(path string, module modules.PassiveModule) *ModuleHandler {
	return &ModuleHandler{
		path:             path,
		module:           module,
		connectionStatus: ConnPending,
	}
}

// handleConnection is the function that will be invoked by the HTTP server this handler is part of
// each time a connection to this handler's path is created.
// It reads websocket messages from the connection, passes them to the module logic,
// and writes back the generated events.
func (mh *ModuleHandler) handleConnection(writer http.ResponseWriter, request *http.Request) {

	// Only accept the first connection.
	if !atomic.CompareAndSwapInt32(&mh.connectionStatus, ConnPending, ConnActive) {
		writer.WriteHeader(http.StatusForbidden)
		return
	}

	// Only accept a websocket connection.
	conn, err := websocket.Accept(writer, request, nil)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	//TODO: Figure out a better way to deal with the context.
	ctx := context.Background()

	// Main loop for reading incoming websocket messages.
	var msgType websocket.MessageType
	var msgData []byte
	for msgType, msgData, err = conn.Read(ctx); err == nil; msgType, msgData, err = conn.Read(ctx) {

		// Only accept binary type messages.
		if msgType != websocket.MessageBinary {
			err = fmt.Errorf("only binary message type is accepted for control messages")
			break
		}

		// The first message must always be a control message, followed by a variable number of event messages.
		command, loadingErr := controlMessageFromBytes(msgData)
		if loadingErr != nil {
			err = fmt.Errorf("could not load control message: %w", loadingErr)
			break
		}

		if command.MsgType == MsgEvents {
			// If the control message announces the number of events that follow,
			// process them all (mh.processEvents reads them from conn).
			err = mh.processEvents(ctx, conn, command.NumEvents)
			if err != nil {
				break
			}
		} else if command.MsgType == MsgClose {
			// If we received a closing message, stop processing.
			break
		} else {
			err = fmt.Errorf("unknown control msg type: %v", command.MsgType)
			break
		}

	}
	if err != nil {
		fmt.Printf("Error processing incoming websocket message: %v\n", err)
	}

	err = conn.Close(websocket.StatusNormalClosure, "")
	if err != nil {
		_ = conn.CloseNow()
	}
}

// processEvents reads messages from the connection, passes them to the module logic,
// and sends back the generated events.
func (mh *ModuleHandler) processEvents(ctx context.Context, conn *websocket.Conn, numEvents int) error {
	resultEvents := stdtypes.EmptyList()
	for ; numEvents > 0; numEvents-- {
		newEvents, err := mh.processNextEvent(ctx, conn)
		if err != nil {
			return err
		}
		resultEvents.PushBackList(newEvents)
	}

	return sendEvents(ctx, conn, resultEvents)
}

// processNextEvent reads a single message from the given websocket connection, applies it to the module logic,
// and returns the resulting events.
func (mh *ModuleHandler) processNextEvent(ctx context.Context, conn *websocket.Conn) (*stdtypes.EventList, error) {
	msgType, msgData, err := conn.Read(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not read control message: %w", err)
	}

	if msgType != websocket.MessageBinary {
		return nil, fmt.Errorf("only binary message type is accepted for events")
	}

	// We can afford using stdevents.Deserialize because sendEvents (used on the other side of the websocket)
	// only ever sends serialized events of type stdevent.RawEvent
	event, err := stdevents.Deserialize(msgData)
	if err != nil {
		return nil, fmt.Errorf("could not deserialize event: %w", err)
	}

	return mh.module.ApplyEvents(stdtypes.ListOf(event))
}

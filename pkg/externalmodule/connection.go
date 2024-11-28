package externalmodule

import (
	"context"
	"fmt"
	"sync"

	"github.com/coder/websocket"

	"github.com/matejpavlovic/mir/stdevents"
	"github.com/matejpavlovic/mir/stdtypes"
)

// Connection represents a connection to a particular module at a particular module server.
// It is used to send events to and receive events from it.
type Connection websocket.Conn

// Connect establishes and returns a new connection
// to a module server at address addr (in the form of "ws://server:port/path").
// The path component of the address is used to specify which module at the module server to connect to.
// When ctx is canceled before the connection is established, connecting aborts.
func Connect(ctx context.Context, addr string) (*Connection, error) {

	conn, _, err := websocket.Dial(ctx, addr, nil)
	if err != nil {
		return nil, err
	}

	return (*Connection)(conn), nil
}

// Submit sends the given events to the remote module, waits until the remote module processes them, and returns
// the resulting events produced by the remote module.
// One can see it as the proxy for the remote module's ApplyEvents method.
func (c *Connection) Submit(ctx context.Context, events *stdtypes.EventList) (*stdtypes.EventList, error) {
	conn := (*websocket.Conn)(c)
	ctx, cancel := context.WithCancel(ctx)
	wg := sync.WaitGroup{}
	wg.Add(1)
	var sendErr error

	// We need to run sendEvents concurrently with receiveResponse to avoid a deadlock.
	// If we first tried to send all the events and only then started receiving the response, the sending could be
	// blocked by the server side blocked by the processing blocked by the sending of response events blocked by the
	// client not having started receiving them.
	go func() {
		sendErr = sendEvents(ctx, conn, events)
		if sendErr != nil {
			cancel() // If sending fails, receiving of the response also must be aborted.
		}
		wg.Done()
	}()

	response, err := receiveResponse(ctx, conn)

	// When reaching this line, sending events must have finished, as receiveResponse would otherwise not have returned.
	// Waiting on the wait group is not necessary in this sense.
	// Nevertheless, we still need to synchronize access to sendErr
	// (and it's good practice to collect spawned goroutines before returning).
	wg.Wait()
	if sendErr != nil {
		return nil, sendErr
	}

	return response, err
}

// Close closes the connection to the remote module.
func (c *Connection) Close(ctx context.Context) error {
	conn := (*websocket.Conn)(c)
	defer func() { _ = conn.CloseNow() }()

	err := conn.Write(ctx, websocket.MessageBinary, controlMessageClose().Bytes())
	if err != nil {
		return err
	}

	return conn.Close(websocket.StatusNormalClosure, "")
}

// sendEvents writes a list of events to the raw websocket connection.
// All sent events are serialized and wrapped in a stdevents.Raw event.
// Thus, on the other side of the connection, only control messages and events of type stdevents.Raw can be expected.
func sendEvents(ctx context.Context, conn *websocket.Conn, events *stdtypes.EventList) error {
	// Announce the number of events that will be sent.
	err := conn.Write(ctx, websocket.MessageBinary, controlMessageEvents(events.Len()).Bytes())
	if err != nil {
		return err
	}

	// Send all the events, using one websocket message per event.
	iter := events.Iterator()
	for event := iter.Next(); event != nil; event = iter.Next() {

		rawEvent, err := stdevents.WrapInRaw(event)
		if err != nil {
			return err
		}

		data, err := rawEvent.ToBytes()
		if err != nil {
			return err
		}

		err = conn.Write(ctx, websocket.MessageBinary, data)
		if err != nil {
			return err
		}
	}

	return nil
}

// receiveResponse reads the events the module server sends over the websocket and returns them in an EventList.
func receiveResponse(ctx context.Context, conn *websocket.Conn) (*stdtypes.EventList, error) {
	// Read the number of resulting events returned from the remote module.
	msgType, msgData, err := conn.Read(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not read response data: %w", err)
	}

	if msgType != websocket.MessageBinary {
		return nil, fmt.Errorf("only binary message type is accepted for control messages")
	}
	command, err := controlMessageFromBytes(msgData)
	if err != nil {
		return nil, fmt.Errorf("could not load control message: %w", err)
	}
	if command.MsgType != MsgEvents {
		return nil, fmt.Errorf("expected MSG_EVENTS control message type but got %v", command.MsgType)
	}

	// Receive the resulting events.
	resultEvents := stdtypes.EmptyList()
	for i := 0; i < command.NumEvents; i++ {

		msgType, msgData, err := conn.Read(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not read response data: %w", err)
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

		resultEvents.PushBack(event)
	}

	return resultEvents, nil
}

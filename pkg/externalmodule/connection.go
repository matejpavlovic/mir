package externalmodule

import (
	"context"
	"fmt"
	"sync"

	"github.com/coder/websocket"

	"github.com/filecoin-project/mir/stdevents"
	"github.com/filecoin-project/mir/stdtypes"
)

type Connection websocket.Conn

func Connect(ctx context.Context, addr string) (*Connection, error) {

	conn, _, err := websocket.Dial(ctx, addr, nil)
	if err != nil {
		return nil, err
	}

	return (*Connection)(conn), nil
}

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

func (c *Connection) Close(ctx context.Context) error {
	conn := (*websocket.Conn)(c)
	defer func() { _ = conn.CloseNow() }()

	err := conn.Write(ctx, websocket.MessageBinary, controlMessageClose().Bytes())
	if err != nil {
		return err
	}

	return conn.Close(websocket.StatusNormalClosure, "")
}

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

		msgType, msgData, err := conn.Read(context.Background())
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

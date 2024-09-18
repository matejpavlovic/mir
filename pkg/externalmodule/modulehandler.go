package externalmodule

import (
	"context"
	"fmt"
	"github.com/coder/websocket"
	"github.com/filecoin-project/mir/pkg/modules"
	"github.com/filecoin-project/mir/stdevents"
	"github.com/filecoin-project/mir/stdtypes"
	"net/http"
	"sync/atomic"
)

const (
	CONN_ACTIVE = iota
	CONN_PENDING
)

type ModuleHandler struct {
	path   string
	module modules.PassiveModule

	// Used to make sure there is only a single client talking to the server.
	// This is needed to prevent concurrent access to the module.
	connectionStatus int32
}

func NewHandler(path string, module modules.PassiveModule) *ModuleHandler {
	return &ModuleHandler{
		path:             path,
		module:           module,
		connectionStatus: CONN_PENDING,
	}
}

func (mh *ModuleHandler) handleConnection(writer http.ResponseWriter, request *http.Request) {

	if !atomic.CompareAndSwapInt32(&mh.connectionStatus, CONN_PENDING, CONN_ACTIVE) {
		writer.WriteHeader(http.StatusForbidden)
		return
	}

	conn, err := websocket.Accept(writer, request, nil)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	//TODO: Figure out a better way to deal with the context.
	ctx := context.Background()

	var msgType websocket.MessageType
	var msgData []byte
	for msgType, msgData, err = conn.Read(ctx); err == nil; msgType, msgData, err = conn.Read(ctx) {

		if msgType != websocket.MessageBinary {
			err = fmt.Errorf("only binary message type is accepted for control messages")
			break
		}

		command, loadingErr := controlMessageFromBytes(msgData)
		if loadingErr != nil {
			err = fmt.Errorf("could not load control message: %w", loadingErr)
			break
		}

		if command.MsgType == MSG_EVENTS {
			err = mh.processEvents(ctx, conn, command.NumEvents)
		} else if command.MsgType == MSG_CLOSE {
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

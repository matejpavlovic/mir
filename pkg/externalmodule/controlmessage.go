package externalmodule

import "encoding/json"

type controlMessageType int

const (
	MsgEvents = iota
	MsgClose
)

type ControlMessage struct {
	MsgType   controlMessageType
	NumEvents int // Only used for EVENT_LIST type.
}

func (cm *ControlMessage) Bytes() []byte {
	data, err := json.Marshal(cm)
	if err != nil {
		panic(err)
	}
	return data
}

func controlMessageFromBytes(data []byte) (*ControlMessage, error) {
	var msg ControlMessage
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

func controlMessageEvents(numEvents int) *ControlMessage {
	return &ControlMessage{MsgEvents, numEvents}
}

func controlMessageClose() *ControlMessage {
	return &ControlMessage{MsgClose, 0}
}

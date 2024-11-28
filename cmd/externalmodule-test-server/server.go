package main

import (
	"fmt"
	"time"

	"github.com/filecoin-project/mir/pkg/externalmodule"
	"github.com/filecoin-project/mir/stdevents"
	"github.com/filecoin-project/mir/stdtypes"
)

type EmptyModule struct {
	prefix string
}

func (e EmptyModule) ImplementsModule() {}

func (e EmptyModule) ApplyEvents(events *stdtypes.EventList) (*stdtypes.EventList, error) {
	fmt.Printf("%s: Received %d event(s).\n", e.prefix, events.Len())
	return stdtypes.ListOf(stdevents.NewTestString("anonymous-module", "Pong")), nil
}

func main() {
	s := externalmodule.NewServer(
		externalmodule.NewHandler("conn1", EmptyModule{"conn1"}),
		externalmodule.NewHandler("conn2", EmptyModule{"conn2"}),
	)

	time.AfterFunc(10*time.Second, func() {
		err := s.Stop()
		if err != nil {
			fmt.Printf("Error stopping server: %v\n", err)
		}
	})

	err := s.Serve("0.0.0.0:8080")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server stopped cleanly.")
	}
}

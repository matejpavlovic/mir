package main

import (
	"context"
	"fmt"
	"time"

	"github.com/filecoin-project/mir/pkg/externalmodule"
	"github.com/filecoin-project/mir/stdevents"
	"github.com/filecoin-project/mir/stdtypes"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn1, err := externalmodule.Connect(ctx, "ws://localhost:8080/conn1")
	if err != nil {
		panic(err)
	}

	response, err := conn1.Submit(ctx, stdtypes.ListOf(
		stdevents.NewTestString("remote", "Ping"),
		stdevents.NewRaw("remote", []byte{0, 1, 2, 3}),
	))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Conn1 received %d events in response.\n", response.Len())

	conn2, err := externalmodule.Connect(ctx, "ws://localhost:8080/conn2")
	if err != nil {
		panic(err)
	}
	response, err = conn2.Submit(ctx, stdtypes.ListOf(
		stdevents.NewTestString("remote", "Ping"),
		stdevents.NewRaw("remote", []byte{0, 1, 2, 3}),
	))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Conn2 received %d events in response.\n", response.Len())

	err = conn1.Close(ctx)
	if err != nil {
		panic(err)
	}
	err = conn2.Close(ctx)
	if err != nil {
		panic(err)
	}
}

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

	beroConn, err := externalmodule.Connect(ctx, "ws://localhost:8080/bero")
	if err != nil {
		panic(err)
	}

	response, err := beroConn.Submit(ctx, stdtypes.ListOf(
		stdevents.NewTestString("remote", "Ping"),
		stdevents.NewRaw("remote", []byte{0, 1, 2, 3}),
	))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Bero received %d events in response.\n", response.Len())

	cecoConn, err := externalmodule.Connect(ctx, "ws://localhost:8080/ceco")
	if err != nil {
		panic(err)
	}
	response, err = cecoConn.Submit(ctx, stdtypes.ListOf(
		stdevents.NewTestString("remote", "Ping"),
		stdevents.NewRaw("remote", []byte{0, 1, 2, 3}),
	))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Ceco received %d events in response.\n", response.Len())

	err = beroConn.Close(ctx)
	if err != nil {
		panic(err)
	}
	err = cecoConn.Close(ctx)
	if err != nil {
		panic(err)
	}
}

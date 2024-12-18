package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	es "github.com/go-errors/errors"
	"google.golang.org/protobuf/proto"

	"github.com/matejpavlovic/mir/pkg/modules"
	"github.com/matejpavlovic/mir/pkg/pb/availabilitypb"
	apbevents "github.com/matejpavlovic/mir/pkg/pb/availabilitypb/events"
	apbtypes "github.com/matejpavlovic/mir/pkg/pb/availabilitypb/types"
	"github.com/matejpavlovic/mir/pkg/pb/eventpb"
	mempoolpbevents "github.com/matejpavlovic/mir/pkg/pb/mempoolpb/events"
	trantorpbtypes "github.com/matejpavlovic/mir/pkg/pb/trantorpb/types"
	"github.com/matejpavlovic/mir/stdevents"
	"github.com/matejpavlovic/mir/stdtypes"
)

type controlModule struct {
	eventsOut           chan *stdtypes.EventList
	readyForNextCommand chan struct{}
}

func newControlModule() modules.ActiveModule {
	return &controlModule{
		eventsOut: make(chan *stdtypes.EventList),
	}
}

func (m *controlModule) ImplementsModule() {}

func (m *controlModule) ApplyEvents(_ context.Context, events *stdtypes.EventList) error {
	iter := events.Iterator()
	for event := iter.Next(); event != nil; event = iter.Next() {

		switch event := event.(type) {
		case *stdevents.Init:
			go func() {
				err := m.readConsole()
				if err != nil {
					panic(err)
				}
			}()
		case *eventpb.Event:
			switch event := event.Type.(type) {

			case *eventpb.Event_Availability:
				switch event := event.Availability.Type.(type) {

				case *availabilitypb.Event_NewCert:
					certBytes, err := proto.Marshal(event.NewCert.Cert)
					if err != nil {
						return es.Errorf("error marshalling certificate: %w", err)
					}

					fmt.Println(base64.StdEncoding.EncodeToString(certBytes))
					close(m.readyForNextCommand)

				case *availabilitypb.Event_ProvideTransactions:
					for _, tx := range event.ProvideTransactions.Txs {
						fmt.Println(string(tx.Data))
					}
					close(m.readyForNextCommand)
				default:
					return es.Errorf("unknown availability event type: %T", event)
				}
			default:
				return es.Errorf("unknown proto event type: %T", event)
			}
		default:
			return es.Errorf("unknown event type: %T", event)
		}
	}

	return nil
}

func (m *controlModule) EventsOut() <-chan *stdtypes.EventList {
	return m.eventsOut
}

func (m *controlModule) readConsole() error {
	// Read the user input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Type in the command ('createBatch', 'readBatch', 'exit')")
		scanner.Scan()
		if scanner.Err() != nil {
			return es.Errorf("error reading from console: %w", scanner.Err())
		}

		text := scanner.Text()

		switch cmd := strings.TrimSpace(text); cmd {
		case "createBatch":
			m.readyForNextCommand = make(chan struct{})
			err := m.createBatch(scanner)
			if err != nil {
				return err
			}
			<-m.readyForNextCommand

		case "readBatch":
			m.readyForNextCommand = make(chan struct{})
			err := m.readBatch(scanner)
			if err != nil {
				return err
			}
			<-m.readyForNextCommand

		case "exit":
			return nil

		default:
			fmt.Println("Unknown command: ", cmd)
		}
	}
}

func (m *controlModule) createBatch(scanner *bufio.Scanner) error {
	fmt.Println("Type in 1 transaction per line, then type 'send!' and press Enter")

	for {
		scanner.Scan()
		if scanner.Err() != nil {
			return es.Errorf("error reading user data: %w", scanner.Err())
		}

		text := scanner.Text()
		if strings.TrimSpace(text) == "send!" {
			break
		}

		tx := &trantorpbtypes.Transaction{Data: []byte(text)}
		m.eventsOut <- stdtypes.ListOf(mempoolpbevents.NewTransactions("mempool", []*trantorpbtypes.Transaction{tx}).Pb())
	}

	m.eventsOut <- stdtypes.ListOf(apbevents.RequestCert("availability", &apbtypes.RequestCertOrigin{
		Module: "control",
		Type:   &apbtypes.RequestCertOrigin_ContextStore{},
	}).Pb())

	return nil
}

func (m *controlModule) readBatch(scanner *bufio.Scanner) error {
	fmt.Println("type in the availability certificate and press Enter")

	scanner.Scan()
	if scanner.Err() != nil {
		return es.Errorf("error reading batch id: %w", scanner.Err())
	}

	certBase64 := strings.TrimSpace(scanner.Text())
	certBytes, err := base64.StdEncoding.DecodeString(certBase64)
	if err != nil {
		return es.Errorf("error decoding certificate: %w", err)
	}

	_cert := new(availabilitypb.Cert)
	err = proto.Unmarshal(certBytes, _cert)
	if err != nil {
		return es.Errorf("error unmarshalling certificate: %w", err)
	}
	cert := apbtypes.CertFromPb(_cert)

	m.eventsOut <- stdtypes.ListOf(apbevents.RequestTransactions("availability", cert,
		&apbtypes.RequestTransactionsOrigin{
			Module: "control",
			Type:   &apbtypes.RequestTransactionsOrigin_ContextStore{},
		}).Pb())

	return nil
}

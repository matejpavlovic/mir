package eventlog

import (
	"database/sql"
	"fmt"

	es "github.com/go-errors/errors"
	_ "github.com/mattn/go-sqlite3" // Driver for the sql database
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/matejpavlovic/mir/pkg/logging"
	"github.com/matejpavlovic/mir/pkg/pb/eventpb"
	t "github.com/matejpavlovic/mir/stdtypes"
)

const (
	newTable string = `CREATE TABLE IF NOT EXISTS events (
    ts INTEGER,
    nodeid INTEGER NOT NULL,
    evtype TEXT,
    evdata TEXT
);`
	insert string = `INSERT INTO events VALUES (?,?,?,?);`
)

type sqliteWriter struct {
	db     *sql.DB
	nodeID t.NodeID
	logger logging.Logger
}

func NewSqliteWriter(filename string, nodeID t.NodeID, logger logging.Logger) (EventWriter, error) {
	db, err := sql.Open("sqlite3", filename+".db")
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(newTable); err != nil {
		return nil, err
	}

	return &sqliteWriter{
		db:     db,
		nodeID: nodeID,
		logger: logger,
	}, nil
}

func (w sqliteWriter) Write(record EventRecord) error {
	// For each incoming event
	iter := record.Events.Iterator()
	for event := iter.Next(); event != nil; event = iter.Next() {

		pbevent, ok := event.(*eventpb.Event)
		if !ok {
			return es.Errorf("SQLite event writer only supports proto events, received %T", event)
		}

		jsonData, err := protojson.Marshal(pbevent)
		if err != nil {
			return err
		}

		_, err = w.db.Exec(
			insert,
			record.Time,
			w.nodeID,
			fmt.Sprintf("%T", pbevent.Type)[len("*eventpb.Event_"):],
			jsonData,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w sqliteWriter) Flush() error {
	return nil
}

func (w sqliteWriter) Close() error {
	return w.db.Close()
}

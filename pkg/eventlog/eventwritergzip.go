package eventlog

import (
	"compress/gzip"
	"encoding/binary"
	"io"
	"os"

	es "github.com/go-errors/errors"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"

	"github.com/matejpavlovic/mir/pkg/logging"
	"github.com/matejpavlovic/mir/pkg/pb/eventpb"
	"github.com/matejpavlovic/mir/pkg/pb/recordingpb"
	"github.com/matejpavlovic/mir/stdtypes"
)

type gzipWriter struct {
	dest             *os.File
	compressionLevel int
	nodeID           stdtypes.NodeID
	logger           logging.Logger
}

func NewGzipWriter(filename string, compressionLevel int, nodeID stdtypes.NodeID, logger logging.Logger) (EventWriter, error) {
	dest, err := os.Create(filename + ".gz")
	if err != nil {
		return nil, es.Errorf("error creating event log file: %w", err)
	}
	return &gzipWriter{
		dest:             dest,
		compressionLevel: compressionLevel,
		nodeID:           nodeID,
		logger:           logger,
	}, nil
}

func (w *gzipWriter) Write(record EventRecord) error {
	gzWriter, err := gzip.NewWriterLevel(w.dest, w.compressionLevel)
	if err != nil {
		return err
	}
	defer func() {
		if err := gzWriter.Close(); err != nil {
			w.logger.Log(logging.LevelError, "Error closing gzWriter.", "err", err)
		}
	}()

	pbEvents, err := pbEventSlice(record.Events)
	if err != nil {
		return err
	}

	return writeRecordedEvent(gzWriter, &recordingpb.Entry{
		NodeId: string(w.nodeID.Bytes()),
		Time:   record.Time,
		Events: pbEvents,
	})
}

func (w *gzipWriter) Flush() error {
	return w.dest.Sync()
}

func (w *gzipWriter) Close() error {
	return w.dest.Close()
}

func writeRecordedEvent(writer io.Writer, entry *recordingpb.Entry) error {
	return writeSizePrefixedProto(writer, entry)
}

func writeSizePrefixedProto(dest io.Writer, msg proto.Message) error {
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return errors.WithMessage(err, "could not marshal")
	}

	lenBuf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(lenBuf, int64(len(msgBytes)))
	if _, err = dest.Write(lenBuf[:n]); err != nil {
		return errors.WithMessage(err, "could not write length prefix")
	}

	if _, err = dest.Write(msgBytes); err != nil {
		return errors.WithMessage(err, "could not write message")
	}

	return nil
}

func pbEventSlice(list *stdtypes.EventList) ([]*eventpb.Event, error) {
	// Create empty result slice.
	result := make([]*eventpb.Event, 0, list.Len())

	// Populate result slice by appending events one by one.
	iter := list.Iterator()
	for event := iter.Next(); event != nil; event = iter.Next() {
		pbevent, ok := event.(*eventpb.Event)
		if !ok {
			data, err := event.ToBytes()
			if err != nil {
				return nil, err
			}
			pbevent = &eventpb.Event{
				DestModule: event.Dest().String(),
				Type:       &eventpb.Event_Serialized{Serialized: &eventpb.SerializedEvent{Data: data}},
			}
		}
		result = append(result, pbevent)
	}

	return result, nil
}

package internal

import (
	"github.com/matejpavlovic/mir/stdtypes"
)

type ModuleImpl interface {
	Event(ev stdtypes.Event) (*stdtypes.EventList, error)
}

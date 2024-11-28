package batchfetcher

import (
	t "github.com/matejpavlovic/mir/stdtypes"
)

// ModuleConfig determines the IDs of the modules the batch fetcher interacts with.
type ModuleConfig struct {
	Self t.ModuleID // Own ID.

	Availability t.ModuleID // ID of the factory module containing the availability modules.
	Checkpoint   t.ModuleID // ID of the checkpoint factory module to which to submit client progress.
	Destination  t.ModuleID // ID of the module to deliver the produced event stream to (usually the application).
	Mempool      t.ModuleID // ID of the mempool module to send updates on client progress to.
}

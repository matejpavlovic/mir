package simplemempool

import (
	"time"

	"github.com/matejpavlovic/mir/pkg/dsl"
	"github.com/matejpavlovic/mir/pkg/logging"
	"github.com/matejpavlovic/mir/pkg/mempool/simplemempool/common"
	"github.com/matejpavlovic/mir/pkg/mempool/simplemempool/internal/parts/computeids"
	"github.com/matejpavlovic/mir/pkg/mempool/simplemempool/internal/parts/formbatchesext"
	"github.com/matejpavlovic/mir/pkg/mempool/simplemempool/internal/parts/formbatchesint"
	"github.com/matejpavlovic/mir/pkg/mempool/simplemempool/internal/parts/lookuptxs"
	"github.com/matejpavlovic/mir/pkg/modules"
	trantorpbtypes "github.com/matejpavlovic/mir/pkg/pb/trantorpb/types"
	tt "github.com/matejpavlovic/mir/pkg/trantor/types"
	"github.com/matejpavlovic/mir/pkg/util/indexedlist"
)

// ModuleConfig sets the module ids. All replicas are expected to use identical module configurations.
type ModuleConfig = common.ModuleConfig

// ModuleParams sets the values for the parameters of an instance of the protocol.
// All replicas are expected to use identical module parameters.
type ModuleParams = common.ModuleParams

func DefaultModuleParams() *ModuleParams {
	return &ModuleParams{
		MaxTransactionsInBatch: 1024,
		MaxPayloadInBatch:      1024 * 1024, // 1 MiB
		BatchTimeout:           100 * time.Millisecond,
		TxFetcher:              nil,
	}
}

// NewModule creates a new instance of a simple mempool module implementation. It passively waits for
// mempoolpb.NewTransactions events and stores them in a local map.
//
// On a batch request, this implementation creates a batch that consists of as many transactions received since the
// previous batch request as possible with respect to params.MaxTransactionsInBatch.
//
// This implementation uses the hash function provided by the mc.Hasher module to compute transaction IDs and batch IDs.
func NewModule(mc ModuleConfig, params *ModuleParams, logger logging.Logger) modules.Module {
	m := dsl.NewModule(mc.Self)

	commonState := &common.State{
		Transactions: indexedlist.New[tt.TxID, *trantorpbtypes.Transaction](),
	}

	computeids.IncludeComputationOfTransactionAndBatchIDs(m, mc, params, logger, commonState)
	lookuptxs.IncludeTransactionLookupByID(m, mc, params, commonState)

	if params.TxFetcher != nil {
		formbatchesext.IncludeBatchCreation(m, mc, params.TxFetcher)
	} else {
		formbatchesint.IncludeBatchCreation(m, mc, params, commonState, logger)
	}

	return m
}

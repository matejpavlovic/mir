package formbatchesext

import (
	"github.com/matejpavlovic/mir/pkg/dsl"
	"github.com/matejpavlovic/mir/pkg/mempool/simplemempool/common"
	mpdsl "github.com/matejpavlovic/mir/pkg/pb/mempoolpb/dsl"
	mppbtypes "github.com/matejpavlovic/mir/pkg/pb/mempoolpb/types"
	trantorpbtypes "github.com/matejpavlovic/mir/pkg/pb/trantorpb/types"
	tt "github.com/matejpavlovic/mir/pkg/trantor/types"
)

// IncludeBatchCreation registers event handlers for processing NewTransactions and RequestBatch events.
func IncludeBatchCreation(
	m dsl.Module,
	mc common.ModuleConfig,
	fetchTransactions func() []*trantorpbtypes.Transaction,
) {

	mpdsl.UponTransactionIDsResponse(m, func(txIDs []tt.TxID, context *requestTxIDsContext) error {
		mpdsl.NewBatch(m, context.origin.Module, txIDs, context.txs, context.origin)
		return nil
	})

	mpdsl.UponRequestBatch(m, func(_ tt.EpochNr, origin *mppbtypes.RequestBatchOrigin) error {
		txs := fetchTransactions()
		mpdsl.RequestTransactionIDs(m, mc.Self, txs, &requestTxIDsContext{
			txs:    txs,
			origin: origin,
		})
		return nil
	})
}

// Context data structures

type requestTxIDsContext struct {
	txs    []*trantorpbtypes.Transaction
	origin *mppbtypes.RequestBatchOrigin
}

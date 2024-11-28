package lookuptxs

import (
	"github.com/matejpavlovic/mir/pkg/dsl"
	"github.com/matejpavlovic/mir/pkg/mempool/simplemempool/common"
	mpdsl "github.com/matejpavlovic/mir/pkg/pb/mempoolpb/dsl"
	mppb "github.com/matejpavlovic/mir/pkg/pb/mempoolpb/types"
	tt "github.com/matejpavlovic/mir/pkg/trantor/types"
)

// IncludeTransactionLookupByID registers event handlers for processing RequestTransactions events.
func IncludeTransactionLookupByID(
	m dsl.Module,
	_ common.ModuleConfig,
	_ *common.ModuleParams,
	state *common.State,
) {
	mpdsl.UponRequestTransactions(m, func(txIDs []tt.TxID, origin *mppb.RequestTransactionsOrigin) error {

		foundIDs, foundTXs, missingIDs := state.Transactions.LookUp(txIDs)

		mpdsl.TransactionsResponse(m, origin.Module, foundIDs, foundTXs, missingIDs, origin)
		return nil
	})
}

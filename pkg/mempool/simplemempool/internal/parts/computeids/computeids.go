package computeids

import (
	"github.com/matejpavlovic/mir/pkg/dsl"
	"github.com/matejpavlovic/mir/pkg/logging"
	"github.com/matejpavlovic/mir/pkg/mempool/simplemempool/common"
	hasherpbdsl "github.com/matejpavlovic/mir/pkg/pb/hasherpb/dsl"
	hasherpbtypes "github.com/matejpavlovic/mir/pkg/pb/hasherpb/types"
	mppbdsl "github.com/matejpavlovic/mir/pkg/pb/mempoolpb/dsl"
	mppbtypes "github.com/matejpavlovic/mir/pkg/pb/mempoolpb/types"
	"github.com/matejpavlovic/mir/pkg/pb/trantorpb"
	trantorpbtypes "github.com/matejpavlovic/mir/pkg/pb/trantorpb/types"
	"github.com/matejpavlovic/mir/pkg/serializing"
	tt "github.com/matejpavlovic/mir/pkg/trantor/types"
	"github.com/matejpavlovic/mir/pkg/util/sliceutil"
)

// IncludeComputationOfTransactionAndBatchIDs registers event handler for processing RequestTransactionIDs and
// RequestBatchID events.
func IncludeComputationOfTransactionAndBatchIDs(
	m dsl.Module,
	mc common.ModuleConfig,
	params *common.ModuleParams,
	logger logging.Logger,
	_ *common.State,
) {
	mppbdsl.UponRequestTransactionIDs(m, func(txs []*trantorpbtypes.Transaction, origin *mppbtypes.RequestTransactionIDsOrigin) error {
		if len(txs) > params.MaxTransactionsInBatch {
			// Invalid request, ignore
			logger.Log(logging.LevelWarn, "Ignoring invalid request: too big for mempool", "numTXs", len(txs))
			return nil
		}

		txMsgs := make([]*hasherpbtypes.HashData, 0, len(txs))
		for i, tx := range txs {
			serializedTx := serializeTXForHash(tx.Pb())
			if serializedTx == nil {
				logger.Log(logging.LevelWarn, "Ignoring invalid request: contains nil transaction", "offset", i)
				return nil
			}
			txMsgs = append(txMsgs, &hasherpbtypes.HashData{Data: serializedTx})
		}

		hasherpbdsl.Request(m, mc.Hasher, txMsgs, &computeHashForTransactionIDsContext{origin})
		return nil
	})

	hasherpbdsl.UponResult(m, func(hashes [][]uint8, context *computeHashForTransactionIDsContext) error {
		mppbdsl.TransactionIDsResponse(
			m,
			context.origin.Module,
			sliceutil.Transform(hashes, func(_ int, hash []uint8) tt.TxID {
				return tt.TxID(hash)
			}),
			context.origin,
		)
		return nil
	})

	mppbdsl.UponRequestBatchID(m, func(txIDs []tt.TxID, origin *mppbtypes.RequestBatchIDOrigin) error {
		hasherpbdsl.RequestOne(
			m,
			mc.Hasher,
			&hasherpbtypes.HashData{Data: sliceutil.Transform(txIDs, func(_ int, txId tt.TxID) []byte {
				return []byte(txId)
			})},
			&computeHashForBatchIDContext{origin},
		)
		return nil
	})

	hasherpbdsl.UponResultOne(m, func(hash []byte, context *computeHashForBatchIDContext) error {
		mppbdsl.BatchIDResponse(m, context.origin.Module, tt.TxID(hash), context.origin)
		return nil
	})
}

// Context data structures

type computeHashForTransactionIDsContext struct {
	origin *mppbtypes.RequestTransactionIDsOrigin
}

type computeHashForBatchIDContext struct {
	origin *mppbtypes.RequestBatchIDOrigin
}

// Auxiliary functions

func serializeTXForHash(tx *trantorpb.Transaction) [][]byte {
	if tx == nil {
		return nil
	}

	// Encode integer fields.
	clientIDBuf := []byte(tx.ClientId)

	// Return serialized integers along with the request data itself.
	return [][]byte{clientIDBuf, serializing.Uint64ToBytes(tx.TxNo), tx.Data}
}

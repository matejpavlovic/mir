// Code generated by Mir codegen. DO NOT EDIT.

package mempoolpbevents

import (
	types4 "github.com/matejpavlovic/mir/pkg/availability/multisigcollector/types"
	types2 "github.com/matejpavlovic/mir/pkg/pb/eventpb/types"
	types1 "github.com/matejpavlovic/mir/pkg/pb/mempoolpb/types"
	types3 "github.com/matejpavlovic/mir/pkg/pb/trantorpb/types"
	types "github.com/matejpavlovic/mir/pkg/trantor/types"
	stdtypes "github.com/matejpavlovic/mir/stdtypes"
)

func RequestBatch(destModule stdtypes.ModuleID, epoch types.EpochNr, origin *types1.RequestBatchOrigin) *types2.Event {
	return &types2.Event{
		DestModule: destModule,
		Type: &types2.Event_Mempool{
			Mempool: &types1.Event{
				Type: &types1.Event_RequestBatch{
					RequestBatch: &types1.RequestBatch{
						Epoch:  epoch,
						Origin: origin,
					},
				},
			},
		},
	}
}

func NewBatch(destModule stdtypes.ModuleID, txIds []types.TxID, txs []*types3.Transaction, origin *types1.RequestBatchOrigin) *types2.Event {
	return &types2.Event{
		DestModule: destModule,
		Type: &types2.Event_Mempool{
			Mempool: &types1.Event{
				Type: &types1.Event_NewBatch{
					NewBatch: &types1.NewBatch{
						TxIds:  txIds,
						Txs:    txs,
						Origin: origin,
					},
				},
			},
		},
	}
}

func RequestTransactions(destModule stdtypes.ModuleID, txIds []types.TxID, origin *types1.RequestTransactionsOrigin) *types2.Event {
	return &types2.Event{
		DestModule: destModule,
		Type: &types2.Event_Mempool{
			Mempool: &types1.Event{
				Type: &types1.Event_RequestTransactions{
					RequestTransactions: &types1.RequestTransactions{
						TxIds:  txIds,
						Origin: origin,
					},
				},
			},
		},
	}
}

func TransactionsResponse(destModule stdtypes.ModuleID, foundIds []types.TxID, foundTxs []*types3.Transaction, missingIds []types.TxID, origin *types1.RequestTransactionsOrigin) *types2.Event {
	return &types2.Event{
		DestModule: destModule,
		Type: &types2.Event_Mempool{
			Mempool: &types1.Event{
				Type: &types1.Event_TransactionsResponse{
					TransactionsResponse: &types1.TransactionsResponse{
						FoundIds:   foundIds,
						FoundTxs:   foundTxs,
						MissingIds: missingIds,
						Origin:     origin,
					},
				},
			},
		},
	}
}

func RequestTransactionIDs(destModule stdtypes.ModuleID, txs []*types3.Transaction, origin *types1.RequestTransactionIDsOrigin) *types2.Event {
	return &types2.Event{
		DestModule: destModule,
		Type: &types2.Event_Mempool{
			Mempool: &types1.Event{
				Type: &types1.Event_RequestTransactionIds{
					RequestTransactionIds: &types1.RequestTransactionIDs{
						Txs:    txs,
						Origin: origin,
					},
				},
			},
		},
	}
}

func TransactionIDsResponse(destModule stdtypes.ModuleID, txIds []types.TxID, origin *types1.RequestTransactionIDsOrigin) *types2.Event {
	return &types2.Event{
		DestModule: destModule,
		Type: &types2.Event_Mempool{
			Mempool: &types1.Event{
				Type: &types1.Event_TransactionIdsResponse{
					TransactionIdsResponse: &types1.TransactionIDsResponse{
						TxIds:  txIds,
						Origin: origin,
					},
				},
			},
		},
	}
}

func RequestBatchID(destModule stdtypes.ModuleID, txIds []types.TxID, origin *types1.RequestBatchIDOrigin) *types2.Event {
	return &types2.Event{
		DestModule: destModule,
		Type: &types2.Event_Mempool{
			Mempool: &types1.Event{
				Type: &types1.Event_RequestBatchId{
					RequestBatchId: &types1.RequestBatchID{
						TxIds:  txIds,
						Origin: origin,
					},
				},
			},
		},
	}
}

func BatchIDResponse(destModule stdtypes.ModuleID, batchId types4.BatchID, origin *types1.RequestBatchIDOrigin) *types2.Event {
	return &types2.Event{
		DestModule: destModule,
		Type: &types2.Event_Mempool{
			Mempool: &types1.Event{
				Type: &types1.Event_BatchIdResponse{
					BatchIdResponse: &types1.BatchIDResponse{
						BatchId: batchId,
						Origin:  origin,
					},
				},
			},
		},
	}
}

func NewTransactions(destModule stdtypes.ModuleID, transactions []*types3.Transaction) *types2.Event {
	return &types2.Event{
		DestModule: destModule,
		Type: &types2.Event_Mempool{
			Mempool: &types1.Event{
				Type: &types1.Event_NewTransactions{
					NewTransactions: &types1.NewTransactions{
						Transactions: transactions,
					},
				},
			},
		},
	}
}

func BatchTimeout(destModule stdtypes.ModuleID, batchReqID uint64) *types2.Event {
	return &types2.Event{
		DestModule: destModule,
		Type: &types2.Event_Mempool{
			Mempool: &types1.Event{
				Type: &types1.Event_BatchTimeout{
					BatchTimeout: &types1.BatchTimeout{
						BatchReqID: batchReqID,
					},
				},
			},
		},
	}
}

func NewEpoch(destModule stdtypes.ModuleID, epochNr types.EpochNr, clientProgress *types3.ClientProgress) *types2.Event {
	return &types2.Event{
		DestModule: destModule,
		Type: &types2.Event_Mempool{
			Mempool: &types1.Event{
				Type: &types1.Event_NewEpoch{
					NewEpoch: &types1.NewEpoch{
						EpochNr:        epochNr,
						ClientProgress: clientProgress,
					},
				},
			},
		},
	}
}

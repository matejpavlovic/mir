// Code generated by Mir codegen. DO NOT EDIT.

package batchdbpbevents

import (
	types "github.com/matejpavlovic/mir/pkg/availability/multisigcollector/types"
	types1 "github.com/matejpavlovic/mir/pkg/pb/availabilitypb/batchdbpb/types"
	types2 "github.com/matejpavlovic/mir/pkg/pb/eventpb/types"
	types3 "github.com/matejpavlovic/mir/pkg/pb/trantorpb/types"
	types4 "github.com/matejpavlovic/mir/pkg/trantor/types"
	stdtypes "github.com/matejpavlovic/mir/stdtypes"
)

func LookupBatch(destModule stdtypes.ModuleID, batchId types.BatchID, origin *types1.LookupBatchOrigin) *types2.Event {
	return &types2.Event{
		DestModule: destModule,
		Type: &types2.Event_BatchDb{
			BatchDb: &types1.Event{
				Type: &types1.Event_Lookup{
					Lookup: &types1.LookupBatch{
						BatchId: batchId,
						Origin:  origin,
					},
				},
			},
		},
	}
}

func LookupBatchResponse(destModule stdtypes.ModuleID, found bool, txs []*types3.Transaction, origin *types1.LookupBatchOrigin) *types2.Event {
	return &types2.Event{
		DestModule: destModule,
		Type: &types2.Event_BatchDb{
			BatchDb: &types1.Event{
				Type: &types1.Event_LookupResponse{
					LookupResponse: &types1.LookupBatchResponse{
						Found:  found,
						Txs:    txs,
						Origin: origin,
					},
				},
			},
		},
	}
}

func StoreBatch(destModule stdtypes.ModuleID, batchId types.BatchID, txs []*types3.Transaction, retentionIndex types4.RetentionIndex, origin *types1.StoreBatchOrigin) *types2.Event {
	return &types2.Event{
		DestModule: destModule,
		Type: &types2.Event_BatchDb{
			BatchDb: &types1.Event{
				Type: &types1.Event_Store{
					Store: &types1.StoreBatch{
						BatchId:        batchId,
						Txs:            txs,
						RetentionIndex: retentionIndex,
						Origin:         origin,
					},
				},
			},
		},
	}
}

func BatchStored(destModule stdtypes.ModuleID, origin *types1.StoreBatchOrigin) *types2.Event {
	return &types2.Event{
		DestModule: destModule,
		Type: &types2.Event_BatchDb{
			BatchDb: &types1.Event{
				Type: &types1.Event_Stored{
					Stored: &types1.BatchStored{
						Origin: origin,
					},
				},
			},
		},
	}
}

func GarbageCollect(destModule stdtypes.ModuleID, retentionIndex types4.RetentionIndex) *types2.Event {
	return &types2.Event{
		DestModule: destModule,
		Type: &types2.Event_BatchDb{
			BatchDb: &types1.Event{
				Type: &types1.Event_GarbageCollect{
					GarbageCollect: &types1.GarbageCollect{
						RetentionIndex: retentionIndex,
					},
				},
			},
		},
	}
}

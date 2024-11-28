// Code generated by Mir codegen. DO NOT EDIT.

package mempoolpbtypes

import (
	mirreflect "github.com/matejpavlovic/mir/codegen/mirreflect"
	types2 "github.com/matejpavlovic/mir/codegen/model/types"
	types3 "github.com/matejpavlovic/mir/pkg/availability/multisigcollector/types"
	types4 "github.com/matejpavlovic/mir/pkg/pb/contextstorepb/types"
	types5 "github.com/matejpavlovic/mir/pkg/pb/dslpb/types"
	mempoolpb "github.com/matejpavlovic/mir/pkg/pb/mempoolpb"
	trantorpb "github.com/matejpavlovic/mir/pkg/pb/trantorpb"
	types1 "github.com/matejpavlovic/mir/pkg/pb/trantorpb/types"
	types "github.com/matejpavlovic/mir/pkg/trantor/types"
	reflectutil "github.com/matejpavlovic/mir/pkg/util/reflectutil"
	stdtypes "github.com/matejpavlovic/mir/stdtypes"
)

type Event struct {
	Type Event_Type
}

type Event_Type interface {
	mirreflect.GeneratedType
	isEvent_Type()
	Pb() mempoolpb.Event_Type
}

type Event_TypeWrapper[T any] interface {
	Event_Type
	Unwrap() *T
}

func Event_TypeFromPb(pb mempoolpb.Event_Type) Event_Type {
	if pb == nil {
		return nil
	}
	switch pb := pb.(type) {
	case *mempoolpb.Event_RequestBatch:
		return &Event_RequestBatch{RequestBatch: RequestBatchFromPb(pb.RequestBatch)}
	case *mempoolpb.Event_NewBatch:
		return &Event_NewBatch{NewBatch: NewBatchFromPb(pb.NewBatch)}
	case *mempoolpb.Event_RequestTransactions:
		return &Event_RequestTransactions{RequestTransactions: RequestTransactionsFromPb(pb.RequestTransactions)}
	case *mempoolpb.Event_TransactionsResponse:
		return &Event_TransactionsResponse{TransactionsResponse: TransactionsResponseFromPb(pb.TransactionsResponse)}
	case *mempoolpb.Event_RequestTransactionIds:
		return &Event_RequestTransactionIds{RequestTransactionIds: RequestTransactionIDsFromPb(pb.RequestTransactionIds)}
	case *mempoolpb.Event_TransactionIdsResponse:
		return &Event_TransactionIdsResponse{TransactionIdsResponse: TransactionIDsResponseFromPb(pb.TransactionIdsResponse)}
	case *mempoolpb.Event_RequestBatchId:
		return &Event_RequestBatchId{RequestBatchId: RequestBatchIDFromPb(pb.RequestBatchId)}
	case *mempoolpb.Event_BatchIdResponse:
		return &Event_BatchIdResponse{BatchIdResponse: BatchIDResponseFromPb(pb.BatchIdResponse)}
	case *mempoolpb.Event_NewTransactions:
		return &Event_NewTransactions{NewTransactions: NewTransactionsFromPb(pb.NewTransactions)}
	case *mempoolpb.Event_BatchTimeout:
		return &Event_BatchTimeout{BatchTimeout: BatchTimeoutFromPb(pb.BatchTimeout)}
	case *mempoolpb.Event_NewEpoch:
		return &Event_NewEpoch{NewEpoch: NewEpochFromPb(pb.NewEpoch)}
	}
	return nil
}

type Event_RequestBatch struct {
	RequestBatch *RequestBatch
}

func (*Event_RequestBatch) isEvent_Type() {}

func (w *Event_RequestBatch) Unwrap() *RequestBatch {
	return w.RequestBatch
}

func (w *Event_RequestBatch) Pb() mempoolpb.Event_Type {
	if w == nil {
		return nil
	}
	if w.RequestBatch == nil {
		return &mempoolpb.Event_RequestBatch{}
	}
	return &mempoolpb.Event_RequestBatch{RequestBatch: (w.RequestBatch).Pb()}
}

func (*Event_RequestBatch) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.Event_RequestBatch]()}
}

type Event_NewBatch struct {
	NewBatch *NewBatch
}

func (*Event_NewBatch) isEvent_Type() {}

func (w *Event_NewBatch) Unwrap() *NewBatch {
	return w.NewBatch
}

func (w *Event_NewBatch) Pb() mempoolpb.Event_Type {
	if w == nil {
		return nil
	}
	if w.NewBatch == nil {
		return &mempoolpb.Event_NewBatch{}
	}
	return &mempoolpb.Event_NewBatch{NewBatch: (w.NewBatch).Pb()}
}

func (*Event_NewBatch) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.Event_NewBatch]()}
}

type Event_RequestTransactions struct {
	RequestTransactions *RequestTransactions
}

func (*Event_RequestTransactions) isEvent_Type() {}

func (w *Event_RequestTransactions) Unwrap() *RequestTransactions {
	return w.RequestTransactions
}

func (w *Event_RequestTransactions) Pb() mempoolpb.Event_Type {
	if w == nil {
		return nil
	}
	if w.RequestTransactions == nil {
		return &mempoolpb.Event_RequestTransactions{}
	}
	return &mempoolpb.Event_RequestTransactions{RequestTransactions: (w.RequestTransactions).Pb()}
}

func (*Event_RequestTransactions) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.Event_RequestTransactions]()}
}

type Event_TransactionsResponse struct {
	TransactionsResponse *TransactionsResponse
}

func (*Event_TransactionsResponse) isEvent_Type() {}

func (w *Event_TransactionsResponse) Unwrap() *TransactionsResponse {
	return w.TransactionsResponse
}

func (w *Event_TransactionsResponse) Pb() mempoolpb.Event_Type {
	if w == nil {
		return nil
	}
	if w.TransactionsResponse == nil {
		return &mempoolpb.Event_TransactionsResponse{}
	}
	return &mempoolpb.Event_TransactionsResponse{TransactionsResponse: (w.TransactionsResponse).Pb()}
}

func (*Event_TransactionsResponse) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.Event_TransactionsResponse]()}
}

type Event_RequestTransactionIds struct {
	RequestTransactionIds *RequestTransactionIDs
}

func (*Event_RequestTransactionIds) isEvent_Type() {}

func (w *Event_RequestTransactionIds) Unwrap() *RequestTransactionIDs {
	return w.RequestTransactionIds
}

func (w *Event_RequestTransactionIds) Pb() mempoolpb.Event_Type {
	if w == nil {
		return nil
	}
	if w.RequestTransactionIds == nil {
		return &mempoolpb.Event_RequestTransactionIds{}
	}
	return &mempoolpb.Event_RequestTransactionIds{RequestTransactionIds: (w.RequestTransactionIds).Pb()}
}

func (*Event_RequestTransactionIds) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.Event_RequestTransactionIds]()}
}

type Event_TransactionIdsResponse struct {
	TransactionIdsResponse *TransactionIDsResponse
}

func (*Event_TransactionIdsResponse) isEvent_Type() {}

func (w *Event_TransactionIdsResponse) Unwrap() *TransactionIDsResponse {
	return w.TransactionIdsResponse
}

func (w *Event_TransactionIdsResponse) Pb() mempoolpb.Event_Type {
	if w == nil {
		return nil
	}
	if w.TransactionIdsResponse == nil {
		return &mempoolpb.Event_TransactionIdsResponse{}
	}
	return &mempoolpb.Event_TransactionIdsResponse{TransactionIdsResponse: (w.TransactionIdsResponse).Pb()}
}

func (*Event_TransactionIdsResponse) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.Event_TransactionIdsResponse]()}
}

type Event_RequestBatchId struct {
	RequestBatchId *RequestBatchID
}

func (*Event_RequestBatchId) isEvent_Type() {}

func (w *Event_RequestBatchId) Unwrap() *RequestBatchID {
	return w.RequestBatchId
}

func (w *Event_RequestBatchId) Pb() mempoolpb.Event_Type {
	if w == nil {
		return nil
	}
	if w.RequestBatchId == nil {
		return &mempoolpb.Event_RequestBatchId{}
	}
	return &mempoolpb.Event_RequestBatchId{RequestBatchId: (w.RequestBatchId).Pb()}
}

func (*Event_RequestBatchId) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.Event_RequestBatchId]()}
}

type Event_BatchIdResponse struct {
	BatchIdResponse *BatchIDResponse
}

func (*Event_BatchIdResponse) isEvent_Type() {}

func (w *Event_BatchIdResponse) Unwrap() *BatchIDResponse {
	return w.BatchIdResponse
}

func (w *Event_BatchIdResponse) Pb() mempoolpb.Event_Type {
	if w == nil {
		return nil
	}
	if w.BatchIdResponse == nil {
		return &mempoolpb.Event_BatchIdResponse{}
	}
	return &mempoolpb.Event_BatchIdResponse{BatchIdResponse: (w.BatchIdResponse).Pb()}
}

func (*Event_BatchIdResponse) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.Event_BatchIdResponse]()}
}

type Event_NewTransactions struct {
	NewTransactions *NewTransactions
}

func (*Event_NewTransactions) isEvent_Type() {}

func (w *Event_NewTransactions) Unwrap() *NewTransactions {
	return w.NewTransactions
}

func (w *Event_NewTransactions) Pb() mempoolpb.Event_Type {
	if w == nil {
		return nil
	}
	if w.NewTransactions == nil {
		return &mempoolpb.Event_NewTransactions{}
	}
	return &mempoolpb.Event_NewTransactions{NewTransactions: (w.NewTransactions).Pb()}
}

func (*Event_NewTransactions) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.Event_NewTransactions]()}
}

type Event_BatchTimeout struct {
	BatchTimeout *BatchTimeout
}

func (*Event_BatchTimeout) isEvent_Type() {}

func (w *Event_BatchTimeout) Unwrap() *BatchTimeout {
	return w.BatchTimeout
}

func (w *Event_BatchTimeout) Pb() mempoolpb.Event_Type {
	if w == nil {
		return nil
	}
	if w.BatchTimeout == nil {
		return &mempoolpb.Event_BatchTimeout{}
	}
	return &mempoolpb.Event_BatchTimeout{BatchTimeout: (w.BatchTimeout).Pb()}
}

func (*Event_BatchTimeout) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.Event_BatchTimeout]()}
}

type Event_NewEpoch struct {
	NewEpoch *NewEpoch
}

func (*Event_NewEpoch) isEvent_Type() {}

func (w *Event_NewEpoch) Unwrap() *NewEpoch {
	return w.NewEpoch
}

func (w *Event_NewEpoch) Pb() mempoolpb.Event_Type {
	if w == nil {
		return nil
	}
	if w.NewEpoch == nil {
		return &mempoolpb.Event_NewEpoch{}
	}
	return &mempoolpb.Event_NewEpoch{NewEpoch: (w.NewEpoch).Pb()}
}

func (*Event_NewEpoch) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.Event_NewEpoch]()}
}

func EventFromPb(pb *mempoolpb.Event) *Event {
	if pb == nil {
		return nil
	}
	return &Event{
		Type: Event_TypeFromPb(pb.Type),
	}
}

func (m *Event) Pb() *mempoolpb.Event {
	if m == nil {
		return nil
	}
	pbMessage := &mempoolpb.Event{}
	{
		if m.Type != nil {
			pbMessage.Type = (m.Type).Pb()
		}
	}

	return pbMessage
}

func (*Event) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.Event]()}
}

type NewEpoch struct {
	EpochNr        types.EpochNr
	ClientProgress *types1.ClientProgress
}

func NewEpochFromPb(pb *mempoolpb.NewEpoch) *NewEpoch {
	if pb == nil {
		return nil
	}
	return &NewEpoch{
		EpochNr:        (types.EpochNr)(pb.EpochNr),
		ClientProgress: types1.ClientProgressFromPb(pb.ClientProgress),
	}
}

func (m *NewEpoch) Pb() *mempoolpb.NewEpoch {
	if m == nil {
		return nil
	}
	pbMessage := &mempoolpb.NewEpoch{}
	{
		pbMessage.EpochNr = (uint64)(m.EpochNr)
		if m.ClientProgress != nil {
			pbMessage.ClientProgress = (m.ClientProgress).Pb()
		}
	}

	return pbMessage
}

func (*NewEpoch) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.NewEpoch]()}
}

type NewTransactions struct {
	Transactions []*types1.Transaction
}

func NewTransactionsFromPb(pb *mempoolpb.NewTransactions) *NewTransactions {
	if pb == nil {
		return nil
	}
	return &NewTransactions{
		Transactions: types2.ConvertSlice(pb.Transactions, func(t *trantorpb.Transaction) *types1.Transaction {
			return types1.TransactionFromPb(t)
		}),
	}
}

func (m *NewTransactions) Pb() *mempoolpb.NewTransactions {
	if m == nil {
		return nil
	}
	pbMessage := &mempoolpb.NewTransactions{}
	{
		pbMessage.Transactions = types2.ConvertSlice(m.Transactions, func(t *types1.Transaction) *trantorpb.Transaction {
			return (t).Pb()
		})
	}

	return pbMessage
}

func (*NewTransactions) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.NewTransactions]()}
}

type RequestBatch struct {
	Epoch  types.EpochNr
	Origin *RequestBatchOrigin
}

func RequestBatchFromPb(pb *mempoolpb.RequestBatch) *RequestBatch {
	if pb == nil {
		return nil
	}
	return &RequestBatch{
		Epoch:  (types.EpochNr)(pb.Epoch),
		Origin: RequestBatchOriginFromPb(pb.Origin),
	}
}

func (m *RequestBatch) Pb() *mempoolpb.RequestBatch {
	if m == nil {
		return nil
	}
	pbMessage := &mempoolpb.RequestBatch{}
	{
		pbMessage.Epoch = (uint64)(m.Epoch)
		if m.Origin != nil {
			pbMessage.Origin = (m.Origin).Pb()
		}
	}

	return pbMessage
}

func (*RequestBatch) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.RequestBatch]()}
}

type NewBatch struct {
	TxIds  []types.TxID
	Txs    []*types1.Transaction
	Origin *RequestBatchOrigin
}

func NewBatchFromPb(pb *mempoolpb.NewBatch) *NewBatch {
	if pb == nil {
		return nil
	}
	return &NewBatch{
		TxIds: types2.ConvertSlice(pb.TxIds, func(t []uint8) types.TxID {
			return (types.TxID)(t)
		}),
		Txs: types2.ConvertSlice(pb.Txs, func(t *trantorpb.Transaction) *types1.Transaction {
			return types1.TransactionFromPb(t)
		}),
		Origin: RequestBatchOriginFromPb(pb.Origin),
	}
}

func (m *NewBatch) Pb() *mempoolpb.NewBatch {
	if m == nil {
		return nil
	}
	pbMessage := &mempoolpb.NewBatch{}
	{
		pbMessage.TxIds = types2.ConvertSlice(m.TxIds, func(t types.TxID) []uint8 {
			return ([]uint8)(t)
		})
		pbMessage.Txs = types2.ConvertSlice(m.Txs, func(t *types1.Transaction) *trantorpb.Transaction {
			return (t).Pb()
		})
		if m.Origin != nil {
			pbMessage.Origin = (m.Origin).Pb()
		}
	}

	return pbMessage
}

func (*NewBatch) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.NewBatch]()}
}

type RequestTransactions struct {
	TxIds  []types.TxID
	Origin *RequestTransactionsOrigin
}

func RequestTransactionsFromPb(pb *mempoolpb.RequestTransactions) *RequestTransactions {
	if pb == nil {
		return nil
	}
	return &RequestTransactions{
		TxIds: types2.ConvertSlice(pb.TxIds, func(t []uint8) types.TxID {
			return (types.TxID)(t)
		}),
		Origin: RequestTransactionsOriginFromPb(pb.Origin),
	}
}

func (m *RequestTransactions) Pb() *mempoolpb.RequestTransactions {
	if m == nil {
		return nil
	}
	pbMessage := &mempoolpb.RequestTransactions{}
	{
		pbMessage.TxIds = types2.ConvertSlice(m.TxIds, func(t types.TxID) []uint8 {
			return ([]uint8)(t)
		})
		if m.Origin != nil {
			pbMessage.Origin = (m.Origin).Pb()
		}
	}

	return pbMessage
}

func (*RequestTransactions) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.RequestTransactions]()}
}

type TransactionsResponse struct {
	FoundIds   []types.TxID
	FoundTxs   []*types1.Transaction
	MissingIds []types.TxID
	Origin     *RequestTransactionsOrigin
}

func TransactionsResponseFromPb(pb *mempoolpb.TransactionsResponse) *TransactionsResponse {
	if pb == nil {
		return nil
	}
	return &TransactionsResponse{
		FoundIds: types2.ConvertSlice(pb.FoundIds, func(t []uint8) types.TxID {
			return (types.TxID)(t)
		}),
		FoundTxs: types2.ConvertSlice(pb.FoundTxs, func(t *trantorpb.Transaction) *types1.Transaction {
			return types1.TransactionFromPb(t)
		}),
		MissingIds: types2.ConvertSlice(pb.MissingIds, func(t []uint8) types.TxID {
			return (types.TxID)(t)
		}),
		Origin: RequestTransactionsOriginFromPb(pb.Origin),
	}
}

func (m *TransactionsResponse) Pb() *mempoolpb.TransactionsResponse {
	if m == nil {
		return nil
	}
	pbMessage := &mempoolpb.TransactionsResponse{}
	{
		pbMessage.FoundIds = types2.ConvertSlice(m.FoundIds, func(t types.TxID) []uint8 {
			return ([]uint8)(t)
		})
		pbMessage.FoundTxs = types2.ConvertSlice(m.FoundTxs, func(t *types1.Transaction) *trantorpb.Transaction {
			return (t).Pb()
		})
		pbMessage.MissingIds = types2.ConvertSlice(m.MissingIds, func(t types.TxID) []uint8 {
			return ([]uint8)(t)
		})
		if m.Origin != nil {
			pbMessage.Origin = (m.Origin).Pb()
		}
	}

	return pbMessage
}

func (*TransactionsResponse) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.TransactionsResponse]()}
}

type RequestTransactionIDs struct {
	Txs    []*types1.Transaction
	Origin *RequestTransactionIDsOrigin
}

func RequestTransactionIDsFromPb(pb *mempoolpb.RequestTransactionIDs) *RequestTransactionIDs {
	if pb == nil {
		return nil
	}
	return &RequestTransactionIDs{
		Txs: types2.ConvertSlice(pb.Txs, func(t *trantorpb.Transaction) *types1.Transaction {
			return types1.TransactionFromPb(t)
		}),
		Origin: RequestTransactionIDsOriginFromPb(pb.Origin),
	}
}

func (m *RequestTransactionIDs) Pb() *mempoolpb.RequestTransactionIDs {
	if m == nil {
		return nil
	}
	pbMessage := &mempoolpb.RequestTransactionIDs{}
	{
		pbMessage.Txs = types2.ConvertSlice(m.Txs, func(t *types1.Transaction) *trantorpb.Transaction {
			return (t).Pb()
		})
		if m.Origin != nil {
			pbMessage.Origin = (m.Origin).Pb()
		}
	}

	return pbMessage
}

func (*RequestTransactionIDs) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.RequestTransactionIDs]()}
}

type TransactionIDsResponse struct {
	TxIds  []types.TxID
	Origin *RequestTransactionIDsOrigin
}

func TransactionIDsResponseFromPb(pb *mempoolpb.TransactionIDsResponse) *TransactionIDsResponse {
	if pb == nil {
		return nil
	}
	return &TransactionIDsResponse{
		TxIds: types2.ConvertSlice(pb.TxIds, func(t []uint8) types.TxID {
			return (types.TxID)(t)
		}),
		Origin: RequestTransactionIDsOriginFromPb(pb.Origin),
	}
}

func (m *TransactionIDsResponse) Pb() *mempoolpb.TransactionIDsResponse {
	if m == nil {
		return nil
	}
	pbMessage := &mempoolpb.TransactionIDsResponse{}
	{
		pbMessage.TxIds = types2.ConvertSlice(m.TxIds, func(t types.TxID) []uint8 {
			return ([]uint8)(t)
		})
		if m.Origin != nil {
			pbMessage.Origin = (m.Origin).Pb()
		}
	}

	return pbMessage
}

func (*TransactionIDsResponse) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.TransactionIDsResponse]()}
}

type RequestBatchID struct {
	TxIds  []types.TxID
	Origin *RequestBatchIDOrigin
}

func RequestBatchIDFromPb(pb *mempoolpb.RequestBatchID) *RequestBatchID {
	if pb == nil {
		return nil
	}
	return &RequestBatchID{
		TxIds: types2.ConvertSlice(pb.TxIds, func(t []uint8) types.TxID {
			return (types.TxID)(t)
		}),
		Origin: RequestBatchIDOriginFromPb(pb.Origin),
	}
}

func (m *RequestBatchID) Pb() *mempoolpb.RequestBatchID {
	if m == nil {
		return nil
	}
	pbMessage := &mempoolpb.RequestBatchID{}
	{
		pbMessage.TxIds = types2.ConvertSlice(m.TxIds, func(t types.TxID) []uint8 {
			return ([]uint8)(t)
		})
		if m.Origin != nil {
			pbMessage.Origin = (m.Origin).Pb()
		}
	}

	return pbMessage
}

func (*RequestBatchID) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.RequestBatchID]()}
}

type BatchIDResponse struct {
	BatchId types3.BatchID
	Origin  *RequestBatchIDOrigin
}

func BatchIDResponseFromPb(pb *mempoolpb.BatchIDResponse) *BatchIDResponse {
	if pb == nil {
		return nil
	}
	return &BatchIDResponse{
		BatchId: (types3.BatchID)(pb.BatchId),
		Origin:  RequestBatchIDOriginFromPb(pb.Origin),
	}
}

func (m *BatchIDResponse) Pb() *mempoolpb.BatchIDResponse {
	if m == nil {
		return nil
	}
	pbMessage := &mempoolpb.BatchIDResponse{}
	{
		pbMessage.BatchId = ([]uint8)(m.BatchId)
		if m.Origin != nil {
			pbMessage.Origin = (m.Origin).Pb()
		}
	}

	return pbMessage
}

func (*BatchIDResponse) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.BatchIDResponse]()}
}

type BatchTimeout struct {
	BatchReqID uint64
}

func BatchTimeoutFromPb(pb *mempoolpb.BatchTimeout) *BatchTimeout {
	if pb == nil {
		return nil
	}
	return &BatchTimeout{
		BatchReqID: pb.BatchReqID,
	}
}

func (m *BatchTimeout) Pb() *mempoolpb.BatchTimeout {
	if m == nil {
		return nil
	}
	pbMessage := &mempoolpb.BatchTimeout{}
	{
		pbMessage.BatchReqID = m.BatchReqID
	}

	return pbMessage
}

func (*BatchTimeout) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.BatchTimeout]()}
}

type RequestBatchOrigin struct {
	Module stdtypes.ModuleID
	Type   RequestBatchOrigin_Type
}

type RequestBatchOrigin_Type interface {
	mirreflect.GeneratedType
	isRequestBatchOrigin_Type()
	Pb() mempoolpb.RequestBatchOrigin_Type
}

type RequestBatchOrigin_TypeWrapper[T any] interface {
	RequestBatchOrigin_Type
	Unwrap() *T
}

func RequestBatchOrigin_TypeFromPb(pb mempoolpb.RequestBatchOrigin_Type) RequestBatchOrigin_Type {
	if pb == nil {
		return nil
	}
	switch pb := pb.(type) {
	case *mempoolpb.RequestBatchOrigin_ContextStore:
		return &RequestBatchOrigin_ContextStore{ContextStore: types4.OriginFromPb(pb.ContextStore)}
	case *mempoolpb.RequestBatchOrigin_Dsl:
		return &RequestBatchOrigin_Dsl{Dsl: types5.OriginFromPb(pb.Dsl)}
	}
	return nil
}

type RequestBatchOrigin_ContextStore struct {
	ContextStore *types4.Origin
}

func (*RequestBatchOrigin_ContextStore) isRequestBatchOrigin_Type() {}

func (w *RequestBatchOrigin_ContextStore) Unwrap() *types4.Origin {
	return w.ContextStore
}

func (w *RequestBatchOrigin_ContextStore) Pb() mempoolpb.RequestBatchOrigin_Type {
	if w == nil {
		return nil
	}
	if w.ContextStore == nil {
		return &mempoolpb.RequestBatchOrigin_ContextStore{}
	}
	return &mempoolpb.RequestBatchOrigin_ContextStore{ContextStore: (w.ContextStore).Pb()}
}

func (*RequestBatchOrigin_ContextStore) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.RequestBatchOrigin_ContextStore]()}
}

type RequestBatchOrigin_Dsl struct {
	Dsl *types5.Origin
}

func (*RequestBatchOrigin_Dsl) isRequestBatchOrigin_Type() {}

func (w *RequestBatchOrigin_Dsl) Unwrap() *types5.Origin {
	return w.Dsl
}

func (w *RequestBatchOrigin_Dsl) Pb() mempoolpb.RequestBatchOrigin_Type {
	if w == nil {
		return nil
	}
	if w.Dsl == nil {
		return &mempoolpb.RequestBatchOrigin_Dsl{}
	}
	return &mempoolpb.RequestBatchOrigin_Dsl{Dsl: (w.Dsl).Pb()}
}

func (*RequestBatchOrigin_Dsl) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.RequestBatchOrigin_Dsl]()}
}

func RequestBatchOriginFromPb(pb *mempoolpb.RequestBatchOrigin) *RequestBatchOrigin {
	if pb == nil {
		return nil
	}
	return &RequestBatchOrigin{
		Module: (stdtypes.ModuleID)(pb.Module),
		Type:   RequestBatchOrigin_TypeFromPb(pb.Type),
	}
}

func (m *RequestBatchOrigin) Pb() *mempoolpb.RequestBatchOrigin {
	if m == nil {
		return nil
	}
	pbMessage := &mempoolpb.RequestBatchOrigin{}
	{
		pbMessage.Module = (string)(m.Module)
		if m.Type != nil {
			pbMessage.Type = (m.Type).Pb()
		}
	}

	return pbMessage
}

func (*RequestBatchOrigin) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.RequestBatchOrigin]()}
}

type RequestTransactionsOrigin struct {
	Module stdtypes.ModuleID
	Type   RequestTransactionsOrigin_Type
}

type RequestTransactionsOrigin_Type interface {
	mirreflect.GeneratedType
	isRequestTransactionsOrigin_Type()
	Pb() mempoolpb.RequestTransactionsOrigin_Type
}

type RequestTransactionsOrigin_TypeWrapper[T any] interface {
	RequestTransactionsOrigin_Type
	Unwrap() *T
}

func RequestTransactionsOrigin_TypeFromPb(pb mempoolpb.RequestTransactionsOrigin_Type) RequestTransactionsOrigin_Type {
	if pb == nil {
		return nil
	}
	switch pb := pb.(type) {
	case *mempoolpb.RequestTransactionsOrigin_ContextStore:
		return &RequestTransactionsOrigin_ContextStore{ContextStore: types4.OriginFromPb(pb.ContextStore)}
	case *mempoolpb.RequestTransactionsOrigin_Dsl:
		return &RequestTransactionsOrigin_Dsl{Dsl: types5.OriginFromPb(pb.Dsl)}
	}
	return nil
}

type RequestTransactionsOrigin_ContextStore struct {
	ContextStore *types4.Origin
}

func (*RequestTransactionsOrigin_ContextStore) isRequestTransactionsOrigin_Type() {}

func (w *RequestTransactionsOrigin_ContextStore) Unwrap() *types4.Origin {
	return w.ContextStore
}

func (w *RequestTransactionsOrigin_ContextStore) Pb() mempoolpb.RequestTransactionsOrigin_Type {
	if w == nil {
		return nil
	}
	if w.ContextStore == nil {
		return &mempoolpb.RequestTransactionsOrigin_ContextStore{}
	}
	return &mempoolpb.RequestTransactionsOrigin_ContextStore{ContextStore: (w.ContextStore).Pb()}
}

func (*RequestTransactionsOrigin_ContextStore) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.RequestTransactionsOrigin_ContextStore]()}
}

type RequestTransactionsOrigin_Dsl struct {
	Dsl *types5.Origin
}

func (*RequestTransactionsOrigin_Dsl) isRequestTransactionsOrigin_Type() {}

func (w *RequestTransactionsOrigin_Dsl) Unwrap() *types5.Origin {
	return w.Dsl
}

func (w *RequestTransactionsOrigin_Dsl) Pb() mempoolpb.RequestTransactionsOrigin_Type {
	if w == nil {
		return nil
	}
	if w.Dsl == nil {
		return &mempoolpb.RequestTransactionsOrigin_Dsl{}
	}
	return &mempoolpb.RequestTransactionsOrigin_Dsl{Dsl: (w.Dsl).Pb()}
}

func (*RequestTransactionsOrigin_Dsl) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.RequestTransactionsOrigin_Dsl]()}
}

func RequestTransactionsOriginFromPb(pb *mempoolpb.RequestTransactionsOrigin) *RequestTransactionsOrigin {
	if pb == nil {
		return nil
	}
	return &RequestTransactionsOrigin{
		Module: (stdtypes.ModuleID)(pb.Module),
		Type:   RequestTransactionsOrigin_TypeFromPb(pb.Type),
	}
}

func (m *RequestTransactionsOrigin) Pb() *mempoolpb.RequestTransactionsOrigin {
	if m == nil {
		return nil
	}
	pbMessage := &mempoolpb.RequestTransactionsOrigin{}
	{
		pbMessage.Module = (string)(m.Module)
		if m.Type != nil {
			pbMessage.Type = (m.Type).Pb()
		}
	}

	return pbMessage
}

func (*RequestTransactionsOrigin) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.RequestTransactionsOrigin]()}
}

type RequestTransactionIDsOrigin struct {
	Module stdtypes.ModuleID
	Type   RequestTransactionIDsOrigin_Type
}

type RequestTransactionIDsOrigin_Type interface {
	mirreflect.GeneratedType
	isRequestTransactionIDsOrigin_Type()
	Pb() mempoolpb.RequestTransactionIDsOrigin_Type
}

type RequestTransactionIDsOrigin_TypeWrapper[T any] interface {
	RequestTransactionIDsOrigin_Type
	Unwrap() *T
}

func RequestTransactionIDsOrigin_TypeFromPb(pb mempoolpb.RequestTransactionIDsOrigin_Type) RequestTransactionIDsOrigin_Type {
	if pb == nil {
		return nil
	}
	switch pb := pb.(type) {
	case *mempoolpb.RequestTransactionIDsOrigin_ContextStore:
		return &RequestTransactionIDsOrigin_ContextStore{ContextStore: types4.OriginFromPb(pb.ContextStore)}
	case *mempoolpb.RequestTransactionIDsOrigin_Dsl:
		return &RequestTransactionIDsOrigin_Dsl{Dsl: types5.OriginFromPb(pb.Dsl)}
	}
	return nil
}

type RequestTransactionIDsOrigin_ContextStore struct {
	ContextStore *types4.Origin
}

func (*RequestTransactionIDsOrigin_ContextStore) isRequestTransactionIDsOrigin_Type() {}

func (w *RequestTransactionIDsOrigin_ContextStore) Unwrap() *types4.Origin {
	return w.ContextStore
}

func (w *RequestTransactionIDsOrigin_ContextStore) Pb() mempoolpb.RequestTransactionIDsOrigin_Type {
	if w == nil {
		return nil
	}
	if w.ContextStore == nil {
		return &mempoolpb.RequestTransactionIDsOrigin_ContextStore{}
	}
	return &mempoolpb.RequestTransactionIDsOrigin_ContextStore{ContextStore: (w.ContextStore).Pb()}
}

func (*RequestTransactionIDsOrigin_ContextStore) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.RequestTransactionIDsOrigin_ContextStore]()}
}

type RequestTransactionIDsOrigin_Dsl struct {
	Dsl *types5.Origin
}

func (*RequestTransactionIDsOrigin_Dsl) isRequestTransactionIDsOrigin_Type() {}

func (w *RequestTransactionIDsOrigin_Dsl) Unwrap() *types5.Origin {
	return w.Dsl
}

func (w *RequestTransactionIDsOrigin_Dsl) Pb() mempoolpb.RequestTransactionIDsOrigin_Type {
	if w == nil {
		return nil
	}
	if w.Dsl == nil {
		return &mempoolpb.RequestTransactionIDsOrigin_Dsl{}
	}
	return &mempoolpb.RequestTransactionIDsOrigin_Dsl{Dsl: (w.Dsl).Pb()}
}

func (*RequestTransactionIDsOrigin_Dsl) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.RequestTransactionIDsOrigin_Dsl]()}
}

func RequestTransactionIDsOriginFromPb(pb *mempoolpb.RequestTransactionIDsOrigin) *RequestTransactionIDsOrigin {
	if pb == nil {
		return nil
	}
	return &RequestTransactionIDsOrigin{
		Module: (stdtypes.ModuleID)(pb.Module),
		Type:   RequestTransactionIDsOrigin_TypeFromPb(pb.Type),
	}
}

func (m *RequestTransactionIDsOrigin) Pb() *mempoolpb.RequestTransactionIDsOrigin {
	if m == nil {
		return nil
	}
	pbMessage := &mempoolpb.RequestTransactionIDsOrigin{}
	{
		pbMessage.Module = (string)(m.Module)
		if m.Type != nil {
			pbMessage.Type = (m.Type).Pb()
		}
	}

	return pbMessage
}

func (*RequestTransactionIDsOrigin) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.RequestTransactionIDsOrigin]()}
}

type RequestBatchIDOrigin struct {
	Module stdtypes.ModuleID
	Type   RequestBatchIDOrigin_Type
}

type RequestBatchIDOrigin_Type interface {
	mirreflect.GeneratedType
	isRequestBatchIDOrigin_Type()
	Pb() mempoolpb.RequestBatchIDOrigin_Type
}

type RequestBatchIDOrigin_TypeWrapper[T any] interface {
	RequestBatchIDOrigin_Type
	Unwrap() *T
}

func RequestBatchIDOrigin_TypeFromPb(pb mempoolpb.RequestBatchIDOrigin_Type) RequestBatchIDOrigin_Type {
	if pb == nil {
		return nil
	}
	switch pb := pb.(type) {
	case *mempoolpb.RequestBatchIDOrigin_ContextStore:
		return &RequestBatchIDOrigin_ContextStore{ContextStore: types4.OriginFromPb(pb.ContextStore)}
	case *mempoolpb.RequestBatchIDOrigin_Dsl:
		return &RequestBatchIDOrigin_Dsl{Dsl: types5.OriginFromPb(pb.Dsl)}
	}
	return nil
}

type RequestBatchIDOrigin_ContextStore struct {
	ContextStore *types4.Origin
}

func (*RequestBatchIDOrigin_ContextStore) isRequestBatchIDOrigin_Type() {}

func (w *RequestBatchIDOrigin_ContextStore) Unwrap() *types4.Origin {
	return w.ContextStore
}

func (w *RequestBatchIDOrigin_ContextStore) Pb() mempoolpb.RequestBatchIDOrigin_Type {
	if w == nil {
		return nil
	}
	if w.ContextStore == nil {
		return &mempoolpb.RequestBatchIDOrigin_ContextStore{}
	}
	return &mempoolpb.RequestBatchIDOrigin_ContextStore{ContextStore: (w.ContextStore).Pb()}
}

func (*RequestBatchIDOrigin_ContextStore) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.RequestBatchIDOrigin_ContextStore]()}
}

type RequestBatchIDOrigin_Dsl struct {
	Dsl *types5.Origin
}

func (*RequestBatchIDOrigin_Dsl) isRequestBatchIDOrigin_Type() {}

func (w *RequestBatchIDOrigin_Dsl) Unwrap() *types5.Origin {
	return w.Dsl
}

func (w *RequestBatchIDOrigin_Dsl) Pb() mempoolpb.RequestBatchIDOrigin_Type {
	if w == nil {
		return nil
	}
	if w.Dsl == nil {
		return &mempoolpb.RequestBatchIDOrigin_Dsl{}
	}
	return &mempoolpb.RequestBatchIDOrigin_Dsl{Dsl: (w.Dsl).Pb()}
}

func (*RequestBatchIDOrigin_Dsl) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.RequestBatchIDOrigin_Dsl]()}
}

func RequestBatchIDOriginFromPb(pb *mempoolpb.RequestBatchIDOrigin) *RequestBatchIDOrigin {
	if pb == nil {
		return nil
	}
	return &RequestBatchIDOrigin{
		Module: (stdtypes.ModuleID)(pb.Module),
		Type:   RequestBatchIDOrigin_TypeFromPb(pb.Type),
	}
}

func (m *RequestBatchIDOrigin) Pb() *mempoolpb.RequestBatchIDOrigin {
	if m == nil {
		return nil
	}
	pbMessage := &mempoolpb.RequestBatchIDOrigin{}
	{
		pbMessage.Module = (string)(m.Module)
		if m.Type != nil {
			pbMessage.Type = (m.Type).Pb()
		}
	}

	return pbMessage
}

func (*RequestBatchIDOrigin) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mempoolpb.RequestBatchIDOrigin]()}
}

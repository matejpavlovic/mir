//
//Copyright IBM Corp. All Rights Reserved.
//
//SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: isspb/isspb.proto

package isspb

import (
	availabilitypb "github.com/matejpavlovic/mir/pkg/pb/availabilitypb"
	checkpointpb "github.com/matejpavlovic/mir/pkg/pb/checkpointpb"
	_ "github.com/matejpavlovic/mir/pkg/pb/mir"
	_ "github.com/matejpavlovic/mir/pkg/pb/net"
	trantorpb "github.com/matejpavlovic/mir/pkg/pb/trantorpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ISSMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*ISSMessage_StableCheckpoint
	Type isISSMessage_Type `protobuf_oneof:"type"`
}

func (x *ISSMessage) Reset() {
	*x = ISSMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isspb_isspb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ISSMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ISSMessage) ProtoMessage() {}

func (x *ISSMessage) ProtoReflect() protoreflect.Message {
	mi := &file_isspb_isspb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ISSMessage.ProtoReflect.Descriptor instead.
func (*ISSMessage) Descriptor() ([]byte, []int) {
	return file_isspb_isspb_proto_rawDescGZIP(), []int{0}
}

func (m *ISSMessage) GetType() isISSMessage_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *ISSMessage) GetStableCheckpoint() *checkpointpb.StableCheckpoint {
	if x, ok := x.GetType().(*ISSMessage_StableCheckpoint); ok {
		return x.StableCheckpoint
	}
	return nil
}

type isISSMessage_Type interface {
	isISSMessage_Type()
}

type ISSMessage_StableCheckpoint struct {
	StableCheckpoint *checkpointpb.StableCheckpoint `protobuf:"bytes,3,opt,name=stable_checkpoint,json=stableCheckpoint,proto3,oneof"`
}

func (*ISSMessage_StableCheckpoint) isISSMessage_Type() {}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*Event_PushCheckpoint
	//	*Event_SbDeliver
	//	*Event_DeliverCert
	//	*Event_NewConfig
	Type isEvent_Type `protobuf_oneof:"type"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isspb_isspb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_isspb_isspb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_isspb_isspb_proto_rawDescGZIP(), []int{1}
}

func (m *Event) GetType() isEvent_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Event) GetPushCheckpoint() *PushCheckpoint {
	if x, ok := x.GetType().(*Event_PushCheckpoint); ok {
		return x.PushCheckpoint
	}
	return nil
}

func (x *Event) GetSbDeliver() *SBDeliver {
	if x, ok := x.GetType().(*Event_SbDeliver); ok {
		return x.SbDeliver
	}
	return nil
}

func (x *Event) GetDeliverCert() *DeliverCert {
	if x, ok := x.GetType().(*Event_DeliverCert); ok {
		return x.DeliverCert
	}
	return nil
}

func (x *Event) GetNewConfig() *NewConfig {
	if x, ok := x.GetType().(*Event_NewConfig); ok {
		return x.NewConfig
	}
	return nil
}

type isEvent_Type interface {
	isEvent_Type()
}

type Event_PushCheckpoint struct {
	PushCheckpoint *PushCheckpoint `protobuf:"bytes,1,opt,name=push_checkpoint,json=pushCheckpoint,proto3,oneof"`
}

type Event_SbDeliver struct {
	SbDeliver *SBDeliver `protobuf:"bytes,2,opt,name=sb_deliver,json=sbDeliver,proto3,oneof"`
}

type Event_DeliverCert struct {
	DeliverCert *DeliverCert `protobuf:"bytes,3,opt,name=deliver_cert,json=deliverCert,proto3,oneof"`
}

type Event_NewConfig struct {
	NewConfig *NewConfig `protobuf:"bytes,4,opt,name=new_config,json=newConfig,proto3,oneof"`
}

func (*Event_PushCheckpoint) isEvent_Type() {}

func (*Event_SbDeliver) isEvent_Type() {}

func (*Event_DeliverCert) isEvent_Type() {}

func (*Event_NewConfig) isEvent_Type() {}

type PushCheckpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushCheckpoint) Reset() {
	*x = PushCheckpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isspb_isspb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushCheckpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushCheckpoint) ProtoMessage() {}

func (x *PushCheckpoint) ProtoReflect() protoreflect.Message {
	mi := &file_isspb_isspb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushCheckpoint.ProtoReflect.Descriptor instead.
func (*PushCheckpoint) Descriptor() ([]byte, []int) {
	return file_isspb_isspb_proto_rawDescGZIP(), []int{2}
}

type SBDeliver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sn         uint64 `protobuf:"varint,1,opt,name=sn,proto3" json:"sn,omitempty"`
	Data       []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Aborted    bool   `protobuf:"varint,3,opt,name=aborted,proto3" json:"aborted,omitempty"`
	Leader     string `protobuf:"bytes,4,opt,name=leader,proto3" json:"leader,omitempty"`
	InstanceId string `protobuf:"bytes,5,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
}

func (x *SBDeliver) Reset() {
	*x = SBDeliver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isspb_isspb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SBDeliver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SBDeliver) ProtoMessage() {}

func (x *SBDeliver) ProtoReflect() protoreflect.Message {
	mi := &file_isspb_isspb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SBDeliver.ProtoReflect.Descriptor instead.
func (*SBDeliver) Descriptor() ([]byte, []int) {
	return file_isspb_isspb_proto_rawDescGZIP(), []int{3}
}

func (x *SBDeliver) GetSn() uint64 {
	if x != nil {
		return x.Sn
	}
	return 0
}

func (x *SBDeliver) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SBDeliver) GetAborted() bool {
	if x != nil {
		return x.Aborted
	}
	return false
}

func (x *SBDeliver) GetLeader() string {
	if x != nil {
		return x.Leader
	}
	return ""
}

func (x *SBDeliver) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

type DeliverCert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sn    uint64               `protobuf:"varint,1,opt,name=sn,proto3" json:"sn,omitempty"`
	Cert  *availabilitypb.Cert `protobuf:"bytes,2,opt,name=cert,proto3" json:"cert,omitempty"`
	Empty bool                 `protobuf:"varint,3,opt,name=empty,proto3" json:"empty,omitempty"`
}

func (x *DeliverCert) Reset() {
	*x = DeliverCert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isspb_isspb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliverCert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliverCert) ProtoMessage() {}

func (x *DeliverCert) ProtoReflect() protoreflect.Message {
	mi := &file_isspb_isspb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliverCert.ProtoReflect.Descriptor instead.
func (*DeliverCert) Descriptor() ([]byte, []int) {
	return file_isspb_isspb_proto_rawDescGZIP(), []int{4}
}

func (x *DeliverCert) GetSn() uint64 {
	if x != nil {
		return x.Sn
	}
	return 0
}

func (x *DeliverCert) GetCert() *availabilitypb.Cert {
	if x != nil {
		return x.Cert
	}
	return nil
}

func (x *DeliverCert) GetEmpty() bool {
	if x != nil {
		return x.Empty
	}
	return false
}

type NewConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EpochNr    uint64                `protobuf:"varint,1,opt,name=epoch_nr,json=epochNr,proto3" json:"epoch_nr,omitempty"`
	Membership *trantorpb.Membership `protobuf:"bytes,2,opt,name=membership,proto3" json:"membership,omitempty"`
}

func (x *NewConfig) Reset() {
	*x = NewConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isspb_isspb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewConfig) ProtoMessage() {}

func (x *NewConfig) ProtoReflect() protoreflect.Message {
	mi := &file_isspb_isspb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewConfig.ProtoReflect.Descriptor instead.
func (*NewConfig) Descriptor() ([]byte, []int) {
	return file_isspb_isspb_proto_rawDescGZIP(), []int{5}
}

func (x *NewConfig) GetEpochNr() uint64 {
	if x != nil {
		return x.EpochNr
	}
	return 0
}

func (x *NewConfig) GetMembership() *trantorpb.Membership {
	if x != nil {
		return x.Membership
	}
	return nil
}

var File_isspb_isspb_proto protoreflect.FileDescriptor

var file_isspb_isspb_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x73, 0x73, 0x70, 0x62, 0x2f, 0x69, 0x73, 0x73, 0x70, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x69, 0x73, 0x73, 0x70, 0x62, 0x1a, 0x19, 0x74, 0x72, 0x61, 0x6e,
	0x74, 0x6f, 0x72, 0x70, 0x62, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x70, 0x62, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x70, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x70, 0x62, 0x2f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x69, 0x72,
	0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6e, 0x65, 0x74, 0x2f, 0x63,
	0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x0a, 0x49, 0x53, 0x53, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4d, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x10, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x3a, 0x04, 0xc8, 0xe4, 0x1d, 0x01, 0x42, 0x0c, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x04, 0xc8, 0xe4, 0x1d, 0x01, 0x22, 0xfc, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x40, 0x0a, 0x0f, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x69, 0x73,
	0x73, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x70, 0x75, 0x73, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x0a, 0x73, 0x62, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x69, 0x73, 0x73, 0x70, 0x62,
	0x2e, 0x53, 0x42, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x48, 0x00, 0x52, 0x09, 0x73, 0x62,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x69, 0x73, 0x73, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x43, 0x65, 0x72,
	0x74, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74,
	0x12, 0x31, 0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x69, 0x73, 0x73, 0x70, 0x62, 0x2e, 0x4e, 0x65, 0x77,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x65, 0x77, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x3a, 0x04, 0x90, 0xa6, 0x1d, 0x01, 0x42, 0x0c, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x04, 0x80, 0xa6, 0x1d, 0x01, 0x22, 0x16, 0x0a, 0x0e, 0x50, 0x75, 0x73, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3a, 0x04, 0x98, 0xa6, 0x1d, 0x01, 0x22,
	0xa8, 0x02, 0x0a, 0x09, 0x53, 0x42, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x12, 0x48, 0x0a,
	0x02, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x38, 0x82, 0xa6, 0x1d, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x74, 0x65, 0x6a, 0x70,
	0x61, 0x76, 0x6c, 0x6f, 0x76, 0x69, 0x63, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65,
	0x71, 0x4e, 0x72, 0x52, 0x02, 0x73, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x62, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x61, 0x62,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x48, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0x82, 0xa6, 0x1d, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x74, 0x65, 0x6a, 0x70, 0x61, 0x76, 0x6c, 0x6f,
	0x76, 0x69, 0x63, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x73, 0x74, 0x64, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x53, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x32, 0x82, 0xa6, 0x1d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x74, 0x65, 0x6a, 0x70, 0x61, 0x76, 0x6c, 0x6f, 0x76,
	0x69, 0x63, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x73, 0x74, 0x64, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x3a, 0x04, 0x98, 0xa6, 0x1d, 0x01, 0x22, 0x9d, 0x01, 0x0a, 0x0b, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x12, 0x48, 0x0a, 0x02, 0x73, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x38, 0x82, 0xa6, 0x1d, 0x34, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x74, 0x65, 0x6a, 0x70, 0x61, 0x76, 0x6c,
	0x6f, 0x76, 0x69, 0x63, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x74, 0x6f, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x71, 0x4e, 0x72,
	0x52, 0x02, 0x73, 0x6e, 0x12, 0x28, 0x0a, 0x04, 0x63, 0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x70, 0x62, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x3a, 0x04, 0x98, 0xa6, 0x1d, 0x01, 0x22, 0x9f, 0x01, 0x0a, 0x09, 0x4e,
	0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x55, 0x0a, 0x08, 0x65, 0x70, 0x6f, 0x63,
	0x68, 0x5f, 0x6e, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x3a, 0x82, 0xa6, 0x1d, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x74, 0x65, 0x6a,
	0x70, 0x61, 0x76, 0x6c, 0x6f, 0x76, 0x69, 0x63, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45,
	0x70, 0x6f, 0x63, 0x68, 0x4e, 0x72, 0x52, 0x07, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x4e, 0x72, 0x12,
	0x35, 0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x3a, 0x04, 0x98, 0xa6, 0x1d, 0x01, 0x42, 0x2b, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x74, 0x65, 0x6a,
	0x70, 0x61, 0x76, 0x6c, 0x6f, 0x76, 0x69, 0x63, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x62, 0x2f, 0x69, 0x73, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_isspb_isspb_proto_rawDescOnce sync.Once
	file_isspb_isspb_proto_rawDescData = file_isspb_isspb_proto_rawDesc
)

func file_isspb_isspb_proto_rawDescGZIP() []byte {
	file_isspb_isspb_proto_rawDescOnce.Do(func() {
		file_isspb_isspb_proto_rawDescData = protoimpl.X.CompressGZIP(file_isspb_isspb_proto_rawDescData)
	})
	return file_isspb_isspb_proto_rawDescData
}

var file_isspb_isspb_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_isspb_isspb_proto_goTypes = []interface{}{
	(*ISSMessage)(nil),                    // 0: isspb.ISSMessage
	(*Event)(nil),                         // 1: isspb.Event
	(*PushCheckpoint)(nil),                // 2: isspb.PushCheckpoint
	(*SBDeliver)(nil),                     // 3: isspb.SBDeliver
	(*DeliverCert)(nil),                   // 4: isspb.DeliverCert
	(*NewConfig)(nil),                     // 5: isspb.NewConfig
	(*checkpointpb.StableCheckpoint)(nil), // 6: checkpointpb.StableCheckpoint
	(*availabilitypb.Cert)(nil),           // 7: availabilitypb.Cert
	(*trantorpb.Membership)(nil),          // 8: trantorpb.Membership
}
var file_isspb_isspb_proto_depIdxs = []int32{
	6, // 0: isspb.ISSMessage.stable_checkpoint:type_name -> checkpointpb.StableCheckpoint
	2, // 1: isspb.Event.push_checkpoint:type_name -> isspb.PushCheckpoint
	3, // 2: isspb.Event.sb_deliver:type_name -> isspb.SBDeliver
	4, // 3: isspb.Event.deliver_cert:type_name -> isspb.DeliverCert
	5, // 4: isspb.Event.new_config:type_name -> isspb.NewConfig
	7, // 5: isspb.DeliverCert.cert:type_name -> availabilitypb.Cert
	8, // 6: isspb.NewConfig.membership:type_name -> trantorpb.Membership
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_isspb_isspb_proto_init() }
func file_isspb_isspb_proto_init() {
	if File_isspb_isspb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_isspb_isspb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ISSMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_isspb_isspb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_isspb_isspb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushCheckpoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_isspb_isspb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SBDeliver); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_isspb_isspb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliverCert); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_isspb_isspb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_isspb_isspb_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ISSMessage_StableCheckpoint)(nil),
	}
	file_isspb_isspb_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Event_PushCheckpoint)(nil),
		(*Event_SbDeliver)(nil),
		(*Event_DeliverCert)(nil),
		(*Event_NewConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_isspb_isspb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_isspb_isspb_proto_goTypes,
		DependencyIndexes: file_isspb_isspb_proto_depIdxs,
		MessageInfos:      file_isspb_isspb_proto_msgTypes,
	}.Build()
	File_isspb_isspb_proto = out.File
	file_isspb_isspb_proto_rawDesc = nil
	file_isspb_isspb_proto_goTypes = nil
	file_isspb_isspb_proto_depIdxs = nil
}

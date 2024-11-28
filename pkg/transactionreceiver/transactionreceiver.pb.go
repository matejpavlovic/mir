//
//Copyright IBM Corp. All Rights Reserved.
//
//SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: transactionreceiver/transactionreceiver.proto

package transactionreceiver

import (
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

type ByeBye struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ByeBye) Reset() {
	*x = ByeBye{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transactionreceiver_transactionreceiver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByeBye) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByeBye) ProtoMessage() {}

func (x *ByeBye) ProtoReflect() protoreflect.Message {
	mi := &file_transactionreceiver_transactionreceiver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByeBye.ProtoReflect.Descriptor instead.
func (*ByeBye) Descriptor() ([]byte, []int) {
	return file_transactionreceiver_transactionreceiver_proto_rawDescGZIP(), []int{0}
}

var File_transactionreceiver_transactionreceiver_proto protoreflect.FileDescriptor

var file_transactionreceiver_transactionreceiver_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x1a, 0x19, 0x74, 0x72, 0x61, 0x6e, 0x74,
	0x6f, 0x72, 0x70, 0x62, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x08, 0x0a, 0x06, 0x42, 0x79, 0x65, 0x42, 0x79, 0x65, 0x32, 0x4b,
	0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x12,
	0x16, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x10, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x2e, 0x42, 0x79, 0x65, 0x42, 0x79, 0x65, 0x28, 0x01, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x74, 0x65, 0x6a, 0x70,
	0x61, 0x76, 0x6c, 0x6f, 0x76, 0x69, 0x63, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transactionreceiver_transactionreceiver_proto_rawDescOnce sync.Once
	file_transactionreceiver_transactionreceiver_proto_rawDescData = file_transactionreceiver_transactionreceiver_proto_rawDesc
)

func file_transactionreceiver_transactionreceiver_proto_rawDescGZIP() []byte {
	file_transactionreceiver_transactionreceiver_proto_rawDescOnce.Do(func() {
		file_transactionreceiver_transactionreceiver_proto_rawDescData = protoimpl.X.CompressGZIP(file_transactionreceiver_transactionreceiver_proto_rawDescData)
	})
	return file_transactionreceiver_transactionreceiver_proto_rawDescData
}

var file_transactionreceiver_transactionreceiver_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transactionreceiver_transactionreceiver_proto_goTypes = []interface{}{
	(*ByeBye)(nil),                // 0: receiver.ByeBye
	(*trantorpb.Transaction)(nil), // 1: trantorpb.Transaction
}
var file_transactionreceiver_transactionreceiver_proto_depIdxs = []int32{
	1, // 0: receiver.TransactionReceiver.Listen:input_type -> trantorpb.Transaction
	0, // 1: receiver.TransactionReceiver.Listen:output_type -> receiver.ByeBye
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transactionreceiver_transactionreceiver_proto_init() }
func file_transactionreceiver_transactionreceiver_proto_init() {
	if File_transactionreceiver_transactionreceiver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transactionreceiver_transactionreceiver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByeBye); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transactionreceiver_transactionreceiver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transactionreceiver_transactionreceiver_proto_goTypes,
		DependencyIndexes: file_transactionreceiver_transactionreceiver_proto_depIdxs,
		MessageInfos:      file_transactionreceiver_transactionreceiver_proto_msgTypes,
	}.Build()
	File_transactionreceiver_transactionreceiver_proto = out.File
	file_transactionreceiver_transactionreceiver_proto_rawDesc = nil
	file_transactionreceiver_transactionreceiver_proto_goTypes = nil
	file_transactionreceiver_transactionreceiver_proto_depIdxs = nil
}

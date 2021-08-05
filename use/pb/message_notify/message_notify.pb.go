// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: message_notify.proto

package message_notify

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PN_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PoolAddress             string `protobuf:"bytes,1,opt,name=pool_address,json=poolAddress,proto3" json:"pool_address,omitempty"`
	LastCheckBlockHeight    string `protobuf:"bytes,2,opt,name=last_check_block_height,json=lastCheckBlockHeight,proto3" json:"last_check_block_height,omitempty"`
	LastCheckBlockHeightTvl string `protobuf:"bytes,3,opt,name=last_check_block_height_tvl,json=lastCheckBlockHeightTvl,proto3" json:"last_check_block_height_tvl,omitempty"`
	CurCheckBlockHeight     string `protobuf:"bytes,4,opt,name=cur_check_block_height,json=curCheckBlockHeight,proto3" json:"cur_check_block_height,omitempty"`
	CurCheckBlockHeightTvl  string `protobuf:"bytes,5,opt,name=cur_check_block_height_tvl,json=curCheckBlockHeightTvl,proto3" json:"cur_check_block_height_tvl,omitempty"`
	MasterChelfAddr         string `protobuf:"bytes,6,opt,name=master_chelf_addr,json=masterChelfAddr,proto3" json:"master_chelf_addr,omitempty"`
}

func (x *PN_Request) Reset() {
	*x = PN_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_notify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PN_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PN_Request) ProtoMessage() {}

func (x *PN_Request) ProtoReflect() protoreflect.Message {
	mi := &file_message_notify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PN_Request.ProtoReflect.Descriptor instead.
func (*PN_Request) Descriptor() ([]byte, []int) {
	return file_message_notify_proto_rawDescGZIP(), []int{0}
}

func (x *PN_Request) GetPoolAddress() string {
	if x != nil {
		return x.PoolAddress
	}
	return ""
}

func (x *PN_Request) GetLastCheckBlockHeight() string {
	if x != nil {
		return x.LastCheckBlockHeight
	}
	return ""
}

func (x *PN_Request) GetLastCheckBlockHeightTvl() string {
	if x != nil {
		return x.LastCheckBlockHeightTvl
	}
	return ""
}

func (x *PN_Request) GetCurCheckBlockHeight() string {
	if x != nil {
		return x.CurCheckBlockHeight
	}
	return ""
}

func (x *PN_Request) GetCurCheckBlockHeightTvl() string {
	if x != nil {
		return x.CurCheckBlockHeightTvl
	}
	return ""
}

func (x *PN_Request) GetMasterChelfAddr() string {
	if x != nil {
		return x.MasterChelfAddr
	}
	return ""
}

type SuccessReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` //true成功 false失败
}

func (x *SuccessReply) Reset() {
	*x = SuccessReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_notify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessReply) ProtoMessage() {}

func (x *SuccessReply) ProtoReflect() protoreflect.Message {
	mi := &file_message_notify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessReply.ProtoReflect.Descriptor instead.
func (*SuccessReply) Descriptor() ([]byte, []int) {
	return file_message_notify_proto_rawDescGZIP(), []int{1}
}

func (x *SuccessReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type FEN_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PoolAddress string `protobuf:"bytes,1,opt,name=pool_address,json=poolAddress,proto3" json:"pool_address,omitempty"` //发生闪电贷的池子合约地址
	Hash        string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`                                  //闪电贷交易哈希
}

func (x *FEN_Request) Reset() {
	*x = FEN_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_notify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FEN_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FEN_Request) ProtoMessage() {}

func (x *FEN_Request) ProtoReflect() protoreflect.Message {
	mi := &file_message_notify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FEN_Request.ProtoReflect.Descriptor instead.
func (*FEN_Request) Descriptor() ([]byte, []int) {
	return file_message_notify_proto_rawDescGZIP(), []int{2}
}

func (x *FEN_Request) GetPoolAddress() string {
	if x != nil {
		return x.PoolAddress
	}
	return ""
}

func (x *FEN_Request) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type FEN_Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` //true成功 false失败
}

func (x *FEN_Reply) Reset() {
	*x = FEN_Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_notify_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FEN_Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FEN_Reply) ProtoMessage() {}

func (x *FEN_Reply) ProtoReflect() protoreflect.Message {
	mi := &file_message_notify_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FEN_Reply.ProtoReflect.Descriptor instead.
func (*FEN_Reply) Descriptor() ([]byte, []int) {
	return file_message_notify_proto_rawDescGZIP(), []int{3}
}

func (x *FEN_Reply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_message_notify_proto protoreflect.FileDescriptor

var file_message_notify_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x22, 0xc1, 0x02, 0x0a, 0x0a, 0x50, 0x4e, 0x5f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6f, 0x6f,
	0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x35, 0x0a, 0x17, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x3c, 0x0a, 0x1b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x74, 0x76, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x54, 0x76, 0x6c, 0x12, 0x33, 0x0a,
	0x16, 0x63, 0x75, 0x72, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63,
	0x75, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x3a, 0x0a, 0x1a, 0x63, 0x75, 0x72, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x74, 0x76, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x63, 0x75, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x54, 0x76, 0x6c, 0x12, 0x2a,
	0x0a, 0x11, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x68, 0x65, 0x6c, 0x66, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x43, 0x68, 0x65, 0x6c, 0x66, 0x41, 0x64, 0x64, 0x72, 0x22, 0x28, 0x0a, 0x0c, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x44, 0x0a, 0x0b, 0x46, 0x45, 0x4e, 0x5f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6f, 0x6f, 0x6c, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x25, 0x0a, 0x09, 0x46, 0x45,
	0x4e, 0x5f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x32, 0xa7, 0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x4c, 0x0a, 0x10, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1b, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x46, 0x45, 0x4e, 0x5f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x46, 0x45, 0x4e, 0x5f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x48, 0x0a, 0x0a, 0x50, 0x6f, 0x6f, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x1a, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2e, 0x50, 0x4e, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e,
	0x2f, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_notify_proto_rawDescOnce sync.Once
	file_message_notify_proto_rawDescData = file_message_notify_proto_rawDesc
)

func file_message_notify_proto_rawDescGZIP() []byte {
	file_message_notify_proto_rawDescOnce.Do(func() {
		file_message_notify_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_notify_proto_rawDescData)
	})
	return file_message_notify_proto_rawDescData
}

var file_message_notify_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_message_notify_proto_goTypes = []interface{}{
	(*PN_Request)(nil),   // 0: message_notify.PN_Request
	(*SuccessReply)(nil), // 1: message_notify.SuccessReply
	(*FEN_Request)(nil),  // 2: message_notify.FEN_Request
	(*FEN_Reply)(nil),    // 3: message_notify.FEN_Reply
}
var file_message_notify_proto_depIdxs = []int32{
	2, // 0: message_notify.MessageNotify.FlashEventNotify:input_type -> message_notify.FEN_Request
	0, // 1: message_notify.MessageNotify.PoolNotify:input_type -> message_notify.PN_Request
	3, // 2: message_notify.MessageNotify.FlashEventNotify:output_type -> message_notify.FEN_Reply
	1, // 3: message_notify.MessageNotify.PoolNotify:output_type -> message_notify.SuccessReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_notify_proto_init() }
func file_message_notify_proto_init() {
	if File_message_notify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_notify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PN_Request); i {
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
		file_message_notify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessReply); i {
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
		file_message_notify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FEN_Request); i {
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
		file_message_notify_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FEN_Reply); i {
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
			RawDescriptor: file_message_notify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_notify_proto_goTypes,
		DependencyIndexes: file_message_notify_proto_depIdxs,
		MessageInfos:      file_message_notify_proto_msgTypes,
	}.Build()
	File_message_notify_proto = out.File
	file_message_notify_proto_rawDesc = nil
	file_message_notify_proto_goTypes = nil
	file_message_notify_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MessageNotifyClient is the client API for MessageNotify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageNotifyClient interface {
	FlashEventNotify(ctx context.Context, in *FEN_Request, opts ...grpc.CallOption) (*FEN_Reply, error)
	PoolNotify(ctx context.Context, in *PN_Request, opts ...grpc.CallOption) (*SuccessReply, error)
}

type messageNotifyClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageNotifyClient(cc grpc.ClientConnInterface) MessageNotifyClient {
	return &messageNotifyClient{cc}
}

func (c *messageNotifyClient) FlashEventNotify(ctx context.Context, in *FEN_Request, opts ...grpc.CallOption) (*FEN_Reply, error) {
	out := new(FEN_Reply)
	err := c.cc.Invoke(ctx, "/message_notify.MessageNotify/FlashEventNotify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageNotifyClient) PoolNotify(ctx context.Context, in *PN_Request, opts ...grpc.CallOption) (*SuccessReply, error) {
	out := new(SuccessReply)
	err := c.cc.Invoke(ctx, "/message_notify.MessageNotify/PoolNotify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageNotifyServer is the server API for MessageNotify service.
type MessageNotifyServer interface {
	FlashEventNotify(context.Context, *FEN_Request) (*FEN_Reply, error)
	PoolNotify(context.Context, *PN_Request) (*SuccessReply, error)
}

// UnimplementedMessageNotifyServer can be embedded to have forward compatible implementations.
type UnimplementedMessageNotifyServer struct {
}

func (*UnimplementedMessageNotifyServer) FlashEventNotify(context.Context, *FEN_Request) (*FEN_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlashEventNotify not implemented")
}
func (*UnimplementedMessageNotifyServer) PoolNotify(context.Context, *PN_Request) (*SuccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PoolNotify not implemented")
}

func RegisterMessageNotifyServer(s *grpc.Server, srv MessageNotifyServer) {
	s.RegisterService(&_MessageNotify_serviceDesc, srv)
}

func _MessageNotify_FlashEventNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FEN_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageNotifyServer).FlashEventNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message_notify.MessageNotify/FlashEventNotify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageNotifyServer).FlashEventNotify(ctx, req.(*FEN_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageNotify_PoolNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PN_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageNotifyServer).PoolNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message_notify.MessageNotify/PoolNotify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageNotifyServer).PoolNotify(ctx, req.(*PN_Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageNotify_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message_notify.MessageNotify",
	HandlerType: (*MessageNotifyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FlashEventNotify",
			Handler:    _MessageNotify_FlashEventNotify_Handler,
		},
		{
			MethodName: "PoolNotify",
			Handler:    _MessageNotify_PoolNotify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message_notify.proto",
}

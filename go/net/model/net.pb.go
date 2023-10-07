// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.3
// source: net.proto

package model

import (
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

type Action int32

const (
	Action_Action_Invalid Action = 0
	Action_Action_Post    Action = 1
	Action_Action_Put     Action = 2
	Action_Action_Patch   Action = 3
	Action_Action_Delete  Action = 4
	Action_Action_Get     Action = 5
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "Action_Invalid",
		1: "Action_Post",
		2: "Action_Put",
		3: "Action_Patch",
		4: "Action_Delete",
		5: "Action_Get",
	}
	Action_value = map[string]int32{
		"Action_Invalid": 0,
		"Action_Post":    1,
		"Action_Put":     2,
		"Action_Patch":   3,
		"Action_Delete":  4,
		"Action_Get":     5,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_net_proto_enumTypes[0].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_net_proto_enumTypes[0]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_net_proto_rawDescGZIP(), []int{0}
}

//
//Secure Message is to transmit a piece of data, securly, from one process to one or more processes.
type SecureMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The source uuid
	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// The destination id, can be a process destination id or a topic.
	Destination string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	// To uniquely identify the source packet, the sender process maintain a sequence number.
	Sequence int32 `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// Priority of this packet
	Priority int32 `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
	// The protobuf marshaled data
	ProtoData []byte `protobuf:"bytes,5,opt,name=proto_data,json=protoData,proto3" json:"proto_data,omitempty"`
	// The protobuf type name of the serialized data
	ProtoTypeName string `protobuf:"bytes,6,opt,name=proto_type_name,json=protoTypeName,proto3" json:"proto_type_name,omitempty"`
	// Action to do with this protobuf
	Action Action `protobuf:"varint,7,opt,name=action,proto3,enum=net_model.Action" json:"action,omitempty"`
}

func (x *SecureMessage) Reset() {
	*x = SecureMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecureMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecureMessage) ProtoMessage() {}

func (x *SecureMessage) ProtoReflect() protoreflect.Message {
	mi := &file_net_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecureMessage.ProtoReflect.Descriptor instead.
func (*SecureMessage) Descriptor() ([]byte, []int) {
	return file_net_proto_rawDescGZIP(), []int{0}
}

func (x *SecureMessage) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *SecureMessage) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *SecureMessage) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *SecureMessage) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *SecureMessage) GetProtoData() []byte {
	if x != nil {
		return x.ProtoData
	}
	return nil
}

func (x *SecureMessage) GetProtoTypeName() string {
	if x != nil {
		return x.ProtoTypeName
	}
	return ""
}

func (x *SecureMessage) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_Action_Invalid
}

type NetConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxDataSize        uint64 `protobuf:"varint,1,opt,name=max_data_size,json=maxDataSize,proto3" json:"max_data_size,omitempty"`
	DefaultTxQueueSize uint64 `protobuf:"varint,2,opt,name=default_tx_queue_size,json=defaultTxQueueSize,proto3" json:"default_tx_queue_size,omitempty"`
	DefaultRxQueueSize uint64 `protobuf:"varint,3,opt,name=default_rx_queue_size,json=defaultRxQueueSize,proto3" json:"default_rx_queue_size,omitempty"`
}

func (x *NetConfig) Reset() {
	*x = NetConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetConfig) ProtoMessage() {}

func (x *NetConfig) ProtoReflect() protoreflect.Message {
	mi := &file_net_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetConfig.ProtoReflect.Descriptor instead.
func (*NetConfig) Descriptor() ([]byte, []int) {
	return file_net_proto_rawDescGZIP(), []int{1}
}

func (x *NetConfig) GetMaxDataSize() uint64 {
	if x != nil {
		return x.MaxDataSize
	}
	return 0
}

func (x *NetConfig) GetDefaultTxQueueSize() uint64 {
	if x != nil {
		return x.DefaultTxQueueSize
	}
	return 0
}

func (x *NetConfig) GetDefaultRxQueueSize() uint64 {
	if x != nil {
		return x.DefaultRxQueueSize
	}
	return 0
}

var File_net_proto protoreflect.FileDescriptor

var file_net_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6e, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6e, 0x65, 0x74,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0xf3, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x75, 0x72,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x29, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x11, 0x2e, 0x6e, 0x65, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x95, 0x01, 0x0a,
	0x09, 0x4e, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x61,
	0x78, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x44, 0x61, 0x74, 0x61, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x31,
	0x0a, 0x15, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x74, 0x78, 0x5f, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x78, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x31, 0x0a, 0x15, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x72, 0x78, 0x5f,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x12, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x78, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x2a, 0x72, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x50, 0x6f, 0x73,
	0x74, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x50, 0x75,
	0x74, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x47, 0x65, 0x74, 0x10, 0x05, 0x42, 0x51, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e,
	0x6d, 0x79, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x42, 0x08, 0x4e, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x01, 0x5a,
	0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x69, 0x63,
	0x68, 0x6c, 0x65, 0x72, 0x2f, 0x6d, 0x79, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x67,
	0x6f, 0x2f, 0x6e, 0x65, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_net_proto_rawDescOnce sync.Once
	file_net_proto_rawDescData = file_net_proto_rawDesc
)

func file_net_proto_rawDescGZIP() []byte {
	file_net_proto_rawDescOnce.Do(func() {
		file_net_proto_rawDescData = protoimpl.X.CompressGZIP(file_net_proto_rawDescData)
	})
	return file_net_proto_rawDescData
}

var file_net_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_net_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_net_proto_goTypes = []interface{}{
	(Action)(0),           // 0: net_model.Action
	(*SecureMessage)(nil), // 1: net_model.SecureMessage
	(*NetConfig)(nil),     // 2: net_model.NetConfig
}
var file_net_proto_depIdxs = []int32{
	0, // 0: net_model.SecureMessage.action:type_name -> net_model.Action
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_net_proto_init() }
func file_net_proto_init() {
	if File_net_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_net_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecureMessage); i {
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
		file_net_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetConfig); i {
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
			RawDescriptor: file_net_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_net_proto_goTypes,
		DependencyIndexes: file_net_proto_depIdxs,
		EnumInfos:         file_net_proto_enumTypes,
		MessageInfos:      file_net_proto_msgTypes,
	}.Build()
	File_net_proto = out.File
	file_net_proto_rawDesc = nil
	file_net_proto_goTypes = nil
	file_net_proto_depIdxs = nil
}

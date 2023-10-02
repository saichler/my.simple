// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.3
// source: directory.proto

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

type State int32

const (
	State_Dir_Invalid State = 0
	State_Dir_Live    State = 1
	State_Dir_Evicted State = 2
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "Dir_Invalid",
		1: "Dir_Live",
		2: "Dir_Evicted",
	}
	State_value = map[string]int32{
		"Dir_Invalid": 0,
		"Dir_Live":    1,
		"Dir_Evicted": 2,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_directory_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_directory_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_directory_proto_rawDescGZIP(), []int{0}
}

type Directory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	States   map[string]*Entry `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Services map[string]*Entry `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Directory) Reset() {
	*x = Directory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Directory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Directory) ProtoMessage() {}

func (x *Directory) ProtoReflect() protoreflect.Message {
	mi := &file_directory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Directory.ProtoReflect.Descriptor instead.
func (*Directory) Descriptor() ([]byte, []int) {
	return file_directory_proto_rawDescGZIP(), []int{0}
}

func (x *Directory) GetStates() map[string]*Entry {
	if x != nil {
		return x.States
	}
	return nil
}

func (x *Directory) GetServices() map[string]*Entry {
	if x != nil {
		return x.Services
	}
	return nil
}

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port      string `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	ServiceId string `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_directory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_directory_proto_rawDescGZIP(), []int{1}
}

func (x *Entry) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *Entry) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *Entry) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceId string            `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Providers map[string]*Entry `protobuf:"bytes,2,rep,name=providers,proto3" json:"providers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_directory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_directory_proto_rawDescGZIP(), []int{2}
}

func (x *Service) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *Service) GetProviders() map[string]*Entry {
	if x != nil {
		return x.Providers
	}
	return nil
}

var File_directory_proto protoreflect.FileDescriptor

var file_directory_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x22, 0xb9, 0x02, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x3e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x12, 0x44, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x51, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x53, 0x0a, 0x0d, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x58,
	0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xc5, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x54, 0x0a, 0x0e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x2a, 0x37, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x69, 0x72,
	0x5f, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x69,
	0x72, 0x5f, 0x4c, 0x69, 0x76, 0x65, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x69, 0x72, 0x5f,
	0x45, 0x76, 0x69, 0x63, 0x74, 0x65, 0x64, 0x10, 0x02, 0x42, 0x63, 0x0a, 0x1d, 0x63, 0x6f, 0x6d,
	0x2e, 0x6d, 0x79, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0e, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x69, 0x63, 0x68, 0x6c, 0x65,
	0x72, 0x2f, 0x6d, 0x79, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_directory_proto_rawDescOnce sync.Once
	file_directory_proto_rawDescData = file_directory_proto_rawDesc
)

func file_directory_proto_rawDescGZIP() []byte {
	file_directory_proto_rawDescOnce.Do(func() {
		file_directory_proto_rawDescData = protoimpl.X.CompressGZIP(file_directory_proto_rawDescData)
	})
	return file_directory_proto_rawDescData
}

var file_directory_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_directory_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_directory_proto_goTypes = []interface{}{
	(State)(0),        // 0: directory_model.State
	(*Directory)(nil), // 1: directory_model.Directory
	(*Entry)(nil),     // 2: directory_model.Entry
	(*Service)(nil),   // 3: directory_model.Service
	nil,               // 4: directory_model.Directory.StatesEntry
	nil,               // 5: directory_model.Directory.ServicesEntry
	nil,               // 6: directory_model.Service.ProvidersEntry
}
var file_directory_proto_depIdxs = []int32{
	4, // 0: directory_model.Directory.states:type_name -> directory_model.Directory.StatesEntry
	5, // 1: directory_model.Directory.services:type_name -> directory_model.Directory.ServicesEntry
	6, // 2: directory_model.Service.providers:type_name -> directory_model.Service.ProvidersEntry
	2, // 3: directory_model.Directory.StatesEntry.value:type_name -> directory_model.Entry
	2, // 4: directory_model.Directory.ServicesEntry.value:type_name -> directory_model.Entry
	2, // 5: directory_model.Service.ProvidersEntry.value:type_name -> directory_model.Entry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_directory_proto_init() }
func file_directory_proto_init() {
	if File_directory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_directory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Directory); i {
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
		file_directory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_directory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
			RawDescriptor: file_directory_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_directory_proto_goTypes,
		DependencyIndexes: file_directory_proto_depIdxs,
		EnumInfos:         file_directory_proto_enumTypes,
		MessageInfos:      file_directory_proto_msgTypes,
	}.Build()
	File_directory_proto = out.File
	file_directory_proto_rawDesc = nil
	file_directory_proto_goTypes = nil
	file_directory_proto_depIdxs = nil
}

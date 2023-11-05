// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.4
// source: introspect.proto

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

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type name, in case of a map, this is the value type.
	TypeName string `protobuf:"bytes,1,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	// In case this attribute is a cell in a map or a slice, this is the key type
	KeyTypeName string `protobuf:"bytes,2,opt,name=key_type_name,json=keyTypeName,proto3" json:"key_type_name,omitempty"`
	// The parent node, nil if root.
	Parent *Node `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
	// The attribute name in the parent
	FieldName string `protobuf:"bytes,4,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	// In case this attribute is a struct, a map between the attribute name and the child registry node.
	Attributes map[string]*Node `protobuf:"bytes,5,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// If this attribute is a map
	IsMap bool `protobuf:"varint,6,opt,name=is_map,json=isMap,proto3" json:"is_map,omitempty"`
	// If this attribute is a slice
	IsSlice bool `protobuf:"varint,7,opt,name=is_slice,json=isSlice,proto3" json:"is_slice,omitempty"`
	//The cached key so we won't need to calculate it all the time.
	CachedKey string `protobuf:"bytes,8,opt,name=cached_key,json=cachedKey,proto3" json:"cached_key,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_introspect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_introspect_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_introspect_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *Node) GetKeyTypeName() string {
	if x != nil {
		return x.KeyTypeName
	}
	return ""
}

func (x *Node) GetParent() *Node {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *Node) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *Node) GetAttributes() map[string]*Node {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Node) GetIsMap() bool {
	if x != nil {
		return x.IsMap
	}
	return false
}

func (x *Node) GetIsSlice() bool {
	if x != nil {
		return x.IsSlice
	}
	return false
}

func (x *Node) GetCachedKey() string {
	if x != nil {
		return x.CachedKey
	}
	return ""
}

var File_introspect_proto protoreflect.FileDescriptor

var file_introspect_proto_rawDesc = []byte{
	0x0a, 0x10, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x22, 0x86, 0x03, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x6b, 0x65,
	0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x70, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x4d, 0x61, 0x70, 0x12, 0x19, 0x0a, 0x08,
	0x69, 0x73, 0x5f, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x69, 0x73, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x64, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x1a, 0x55, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x6e, 0x74,
	0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x66, 0x0a,
	0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x79, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x69,
	0x6e, 0x74, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42,
	0x0f, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x61, 0x69, 0x63, 0x68, 0x6c, 0x65, 0x72, 0x2f, 0x6d, 0x79, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_introspect_proto_rawDescOnce sync.Once
	file_introspect_proto_rawDescData = file_introspect_proto_rawDesc
)

func file_introspect_proto_rawDescGZIP() []byte {
	file_introspect_proto_rawDescOnce.Do(func() {
		file_introspect_proto_rawDescData = protoimpl.X.CompressGZIP(file_introspect_proto_rawDescData)
	})
	return file_introspect_proto_rawDescData
}

var file_introspect_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_introspect_proto_goTypes = []interface{}{
	(*Node)(nil), // 0: introspect_model.Node
	nil,          // 1: introspect_model.Node.AttributesEntry
}
var file_introspect_proto_depIdxs = []int32{
	0, // 0: introspect_model.Node.parent:type_name -> introspect_model.Node
	1, // 1: introspect_model.Node.attributes:type_name -> introspect_model.Node.AttributesEntry
	0, // 2: introspect_model.Node.AttributesEntry.value:type_name -> introspect_model.Node
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_introspect_proto_init() }
func file_introspect_proto_init() {
	if File_introspect_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_introspect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
			RawDescriptor: file_introspect_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_introspect_proto_goTypes,
		DependencyIndexes: file_introspect_proto_depIdxs,
		MessageInfos:      file_introspect_proto_msgTypes,
	}.Build()
	File_introspect_proto = out.File
	file_introspect_proto_rawDesc = nil
	file_introspect_proto_goTypes = nil
	file_introspect_proto_depIdxs = nil
}

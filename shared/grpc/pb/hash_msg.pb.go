// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: hash_msg.proto

package pb

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

type Hash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Hash) Reset() {
	*x = Hash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hash_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hash) ProtoMessage() {}

func (x *Hash) ProtoReflect() protoreflect.Message {
	mi := &file_hash_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hash.ProtoReflect.Descriptor instead.
func (*Hash) Descriptor() ([]byte, []int) {
	return file_hash_msg_proto_rawDescGZIP(), []int{0}
}

func (x *Hash) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Exists struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Exists) Reset() {
	*x = Exists{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hash_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Exists) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exists) ProtoMessage() {}

func (x *Exists) ProtoReflect() protoreflect.Message {
	mi := &file_hash_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Exists.ProtoReflect.Descriptor instead.
func (*Exists) Descriptor() ([]byte, []int) {
	return file_hash_msg_proto_rawDescGZIP(), []int{1}
}

func (x *Exists) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type HashingPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Src string `protobuf:"bytes,1,opt,name=src,proto3" json:"src,omitempty"`
}

func (x *HashingPayload) Reset() {
	*x = HashingPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hash_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashingPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashingPayload) ProtoMessage() {}

func (x *HashingPayload) ProtoReflect() protoreflect.Message {
	mi := &file_hash_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashingPayload.ProtoReflect.Descriptor instead.
func (*HashingPayload) Descriptor() ([]byte, []int) {
	return file_hash_msg_proto_rawDescGZIP(), []int{2}
}

func (x *HashingPayload) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

type HashedPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *HashedPayload) Reset() {
	*x = HashedPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hash_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashedPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashedPayload) ProtoMessage() {}

func (x *HashedPayload) ProtoReflect() protoreflect.Message {
	mi := &file_hash_msg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashedPayload.ProtoReflect.Descriptor instead.
func (*HashedPayload) Descriptor() ([]byte, []int) {
	return file_hash_msg_proto_rawDescGZIP(), []int{3}
}

func (x *HashedPayload) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

var File_hash_msg_proto protoreflect.FileDescriptor

var file_hash_msg_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x1c, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x1e, 0x0a, 0x06, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x22, 0x0a, 0x0e, 0x48, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x72, 0x63, 0x22, 0x23, 0x0a, 0x0d, 0x48, 0x61, 0x73, 0x68, 0x65, 0x64,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hash_msg_proto_rawDescOnce sync.Once
	file_hash_msg_proto_rawDescData = file_hash_msg_proto_rawDesc
)

func file_hash_msg_proto_rawDescGZIP() []byte {
	file_hash_msg_proto_rawDescOnce.Do(func() {
		file_hash_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_hash_msg_proto_rawDescData)
	})
	return file_hash_msg_proto_rawDescData
}

var file_hash_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_hash_msg_proto_goTypes = []interface{}{
	(*Hash)(nil),           // 0: pb.Hash
	(*Exists)(nil),         // 1: pb.Exists
	(*HashingPayload)(nil), // 2: pb.HashingPayload
	(*HashedPayload)(nil),  // 3: pb.HashedPayload
}
var file_hash_msg_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hash_msg_proto_init() }
func file_hash_msg_proto_init() {
	if File_hash_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hash_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hash); i {
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
		file_hash_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Exists); i {
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
		file_hash_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashingPayload); i {
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
		file_hash_msg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashedPayload); i {
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
			RawDescriptor: file_hash_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hash_msg_proto_goTypes,
		DependencyIndexes: file_hash_msg_proto_depIdxs,
		MessageInfos:      file_hash_msg_proto_msgTypes,
	}.Build()
	File_hash_msg_proto = out.File
	file_hash_msg_proto_rawDesc = nil
	file_hash_msg_proto_goTypes = nil
	file_hash_msg_proto_depIdxs = nil
}
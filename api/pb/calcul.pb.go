// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: calcul.proto

package calcul

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int64 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int64 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calcul_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_calcul_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReq.ProtoReflect.Descriptor instead.
func (*AddReq) Descriptor() ([]byte, []int) {
	return file_calcul_proto_rawDescGZIP(), []int{0}
}

func (x *AddReq) GetA() int64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *AddReq) GetB() int64 {
	if x != nil {
		return x.B
	}
	return 0
}

type AddRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AddRes) Reset() {
	*x = AddRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calcul_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRes) ProtoMessage() {}

func (x *AddRes) ProtoReflect() protoreflect.Message {
	mi := &file_calcul_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRes.ProtoReflect.Descriptor instead.
func (*AddRes) Descriptor() ([]byte, []int) {
	return file_calcul_proto_rawDescGZIP(), []int{1}
}

func (x *AddRes) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calcul_proto protoreflect.FileDescriptor

var file_calcul_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x22, 0x24, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71,
	0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x61, 0x12, 0x0c,
	0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x62, 0x22, 0x20, 0x0a, 0x06,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x36,
	0x0a, 0x0d, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x25, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x2e,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x2e,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calcul_proto_rawDescOnce sync.Once
	file_calcul_proto_rawDescData = file_calcul_proto_rawDesc
)

func file_calcul_proto_rawDescGZIP() []byte {
	file_calcul_proto_rawDescOnce.Do(func() {
		file_calcul_proto_rawDescData = protoimpl.X.CompressGZIP(file_calcul_proto_rawDescData)
	})
	return file_calcul_proto_rawDescData
}

var file_calcul_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_calcul_proto_goTypes = []interface{}{
	(*AddReq)(nil), // 0: calcul.AddReq
	(*AddRes)(nil), // 1: calcul.AddRes
}
var file_calcul_proto_depIdxs = []int32{
	0, // 0: calcul.CalculService.Add:input_type -> calcul.AddReq
	1, // 1: calcul.CalculService.Add:output_type -> calcul.AddRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calcul_proto_init() }
func file_calcul_proto_init() {
	if File_calcul_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calcul_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReq); i {
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
		file_calcul_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRes); i {
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
			RawDescriptor: file_calcul_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calcul_proto_goTypes,
		DependencyIndexes: file_calcul_proto_depIdxs,
		MessageInfos:      file_calcul_proto_msgTypes,
	}.Build()
	File_calcul_proto = out.File
	file_calcul_proto_rawDesc = nil
	file_calcul_proto_goTypes = nil
	file_calcul_proto_depIdxs = nil
}
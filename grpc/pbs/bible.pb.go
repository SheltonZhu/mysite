// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: grpc/protos/bible.proto

package pbs

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

type Bible struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Bible) Reset() {
	*x = Bible{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_protos_bible_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bible) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bible) ProtoMessage() {}

func (x *Bible) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_protos_bible_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bible.ProtoReflect.Descriptor instead.
func (*Bible) Descriptor() ([]byte, []int) {
	return file_grpc_protos_bible_proto_rawDescGZIP(), []int{0}
}

func (x *Bible) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Bible) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Bibles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bibles []*Bible `protobuf:"bytes,1,rep,name=bibles,proto3" json:"bibles,omitempty"`
}

func (x *Bibles) Reset() {
	*x = Bibles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_protos_bible_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bibles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bibles) ProtoMessage() {}

func (x *Bibles) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_protos_bible_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bibles.ProtoReflect.Descriptor instead.
func (*Bibles) Descriptor() ([]byte, []int) {
	return file_grpc_protos_bible_proto_rawDescGZIP(), []int{1}
}

func (x *Bibles) GetBibles() []*Bible {
	if x != nil {
		return x.Bibles
	}
	return nil
}

var File_grpc_protos_bible_proto protoreflect.FileDescriptor

var file_grpc_protos_bible_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x69,
	0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x1a, 0x16, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75,
	0x74, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x05, 0x42, 0x69, 0x62,
	0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x2f, 0x0a, 0x06, 0x42, 0x69, 0x62, 0x6c, 0x65, 0x73,
	0x12, 0x25, 0x0a, 0x06, 0x62, 0x69, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x69, 0x62, 0x6c, 0x65, 0x52,
	0x06, 0x62, 0x69, 0x62, 0x6c, 0x65, 0x73, 0x32, 0x98, 0x01, 0x0a, 0x0c, 0x42, 0x69, 0x62, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x69, 0x62,
	0x6c, 0x65, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x69, 0x62, 0x6c, 0x65, 0x73, 0x22, 0x00, 0x12, 0x33, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_protos_bible_proto_rawDescOnce sync.Once
	file_grpc_protos_bible_proto_rawDescData = file_grpc_protos_bible_proto_rawDesc
)

func file_grpc_protos_bible_proto_rawDescGZIP() []byte {
	file_grpc_protos_bible_proto_rawDescOnce.Do(func() {
		file_grpc_protos_bible_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_protos_bible_proto_rawDescData)
	})
	return file_grpc_protos_bible_proto_rawDescData
}

var file_grpc_protos_bible_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_protos_bible_proto_goTypes = []interface{}{
	(*Bible)(nil),       // 0: protos.Bible
	(*Bibles)(nil),      // 1: protos.Bibles
	(*Int32Value)(nil),  // 2: protos.Int32Value
	(*Empty)(nil),       // 3: protos.Empty
	(*StringValue)(nil), // 4: protos.StringValue
}
var file_grpc_protos_bible_proto_depIdxs = []int32{
	0, // 0: protos.Bibles.bibles:type_name -> protos.Bible
	2, // 1: protos.BibleService.Get:input_type -> protos.Int32Value
	3, // 2: protos.BibleService.List:input_type -> protos.Empty
	4, // 3: protos.BibleService.Create:input_type -> protos.StringValue
	0, // 4: protos.BibleService.Get:output_type -> protos.Bible
	1, // 5: protos.BibleService.List:output_type -> protos.Bibles
	2, // 6: protos.BibleService.Create:output_type -> protos.Int32Value
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_protos_bible_proto_init() }
func file_grpc_protos_bible_proto_init() {
	if File_grpc_protos_bible_proto != nil {
		return
	}
	file_grpc_protos_util_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_grpc_protos_bible_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bible); i {
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
		file_grpc_protos_bible_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bibles); i {
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
			RawDescriptor: file_grpc_protos_bible_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_protos_bible_proto_goTypes,
		DependencyIndexes: file_grpc_protos_bible_proto_depIdxs,
		MessageInfos:      file_grpc_protos_bible_proto_msgTypes,
	}.Build()
	File_grpc_protos_bible_proto = out.File
	file_grpc_protos_bible_proto_rawDesc = nil
	file_grpc_protos_bible_proto_goTypes = nil
	file_grpc_protos_bible_proto_depIdxs = nil
}

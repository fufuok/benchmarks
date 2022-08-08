// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.8.0
// source: proto/benchmark.proto

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

type Message_Corpus int32

const (
	Message_UNIVERSAL Message_Corpus = 0
	Message_WEB       Message_Corpus = 1
	Message_IMAGES    Message_Corpus = 2
	Message_LOCAL     Message_Corpus = 3
	Message_NEWS      Message_Corpus = 4
	Message_PRODUCTS  Message_Corpus = 5
	Message_VIDEO     Message_Corpus = 6
)

// Enum value maps for Message_Corpus.
var (
	Message_Corpus_name = map[int32]string{
		0: "UNIVERSAL",
		1: "WEB",
		2: "IMAGES",
		3: "LOCAL",
		4: "NEWS",
		5: "PRODUCTS",
		6: "VIDEO",
	}
	Message_Corpus_value = map[string]int32{
		"UNIVERSAL": 0,
		"WEB":       1,
		"IMAGES":    2,
		"LOCAL":     3,
		"NEWS":      4,
		"PRODUCTS":  5,
		"VIDEO":     6,
	}
)

func (x Message_Corpus) Enum() *Message_Corpus {
	p := new(Message_Corpus)
	*p = x
	return p
}

func (x Message_Corpus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Message_Corpus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_benchmark_proto_enumTypes[0].Descriptor()
}

func (Message_Corpus) Type() protoreflect.EnumType {
	return &file_proto_benchmark_proto_enumTypes[0]
}

func (x Message_Corpus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Message_Corpus.Descriptor instead.
func (Message_Corpus) EnumDescriptor() ([]byte, []int) {
	return file_proto_benchmark_proto_rawDescGZIP(), []int{0, 0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query         string         `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber    int32          `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultPerPage int32          `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"`
	Comment       []byte         `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
	Corpus        Message_Corpus `protobuf:"varint,5,opt,name=corpus,proto3,enum=Message_Corpus" json:"corpus,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_benchmark_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_benchmark_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_benchmark_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *Message) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *Message) GetResultPerPage() int32 {
	if x != nil {
		return x.ResultPerPage
	}
	return 0
}

func (x *Message) GetComment() []byte {
	if x != nil {
		return x.Comment
	}
	return nil
}

func (x *Message) GetCorpus() Message_Corpus {
	if x != nil {
		return x.Corpus
	}
	return Message_UNIVERSAL
}

var File_proto_benchmark_proto protoreflect.FileDescriptor

var file_proto_benchmark_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x65, 0x72, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x06,
	0x63, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x52, 0x06, 0x63,
	0x6f, 0x72, 0x70, 0x75, 0x73, 0x22, 0x5a, 0x0a, 0x06, 0x43, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x12,
	0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x49, 0x56, 0x45, 0x52, 0x53, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x57, 0x45, 0x42, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x4d, 0x41, 0x47, 0x45,
	0x53, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x45, 0x57, 0x53, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x4f, 0x44,
	0x55, 0x43, 0x54, 0x53, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10,
	0x06, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_benchmark_proto_rawDescOnce sync.Once
	file_proto_benchmark_proto_rawDescData = file_proto_benchmark_proto_rawDesc
)

func file_proto_benchmark_proto_rawDescGZIP() []byte {
	file_proto_benchmark_proto_rawDescOnce.Do(func() {
		file_proto_benchmark_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_benchmark_proto_rawDescData)
	})
	return file_proto_benchmark_proto_rawDescData
}

var file_proto_benchmark_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_benchmark_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_benchmark_proto_goTypes = []interface{}{
	(Message_Corpus)(0), // 0: Message.Corpus
	(*Message)(nil),     // 1: Message
}
var file_proto_benchmark_proto_depIdxs = []int32{
	0, // 0: Message.corpus:type_name -> Message.Corpus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_benchmark_proto_init() }
func file_proto_benchmark_proto_init() {
	if File_proto_benchmark_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_benchmark_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_proto_benchmark_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_benchmark_proto_goTypes,
		DependencyIndexes: file_proto_benchmark_proto_depIdxs,
		EnumInfos:         file_proto_benchmark_proto_enumTypes,
		MessageInfos:      file_proto_benchmark_proto_msgTypes,
	}.Build()
	File_proto_benchmark_proto = out.File
	file_proto_benchmark_proto_rawDesc = nil
	file_proto_benchmark_proto_goTypes = nil
	file_proto_benchmark_proto_depIdxs = nil
}

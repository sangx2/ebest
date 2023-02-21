// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v4.22.0
// source: proto/news.proto

package proto

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

type NewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerName string `protobuf:"bytes,6,opt,name=ServerName,proto3" json:"ServerName,omitempty"`
}

func (x *NewsRequest) Reset() {
	*x = NewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_news_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsRequest) ProtoMessage() {}

func (x *NewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_news_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsRequest.ProtoReflect.Descriptor instead.
func (*NewsRequest) Descriptor() ([]byte, []int) {
	return file_proto_news_proto_rawDescGZIP(), []int{0}
}

func (x *NewsRequest) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

type NewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language  string `protobuf:"bytes,1,opt,name=Language,proto3" json:"Language,omitempty"`
	Date      string `protobuf:"bytes,2,opt,name=Date,proto3" json:"Date,omitempty"`
	Time      string `protobuf:"bytes,3,opt,name=Time,proto3" json:"Time,omitempty"`
	Publisher string `protobuf:"bytes,4,opt,name=Publisher,proto3" json:"Publisher,omitempty"`
	Title     string `protobuf:"bytes,5,opt,name=Title,proto3" json:"Title,omitempty"`
	Body      string `protobuf:"bytes,6,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *NewsResponse) Reset() {
	*x = NewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_news_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsResponse) ProtoMessage() {}

func (x *NewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_news_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsResponse.ProtoReflect.Descriptor instead.
func (*NewsResponse) Descriptor() ([]byte, []int) {
	return file_proto_news_proto_rawDescGZIP(), []int{1}
}

func (x *NewsResponse) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *NewsResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *NewsResponse) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *NewsResponse) GetPublisher() string {
	if x != nil {
		return x.Publisher
	}
	return ""
}

func (x *NewsResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NewsResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_proto_news_proto protoreflect.FileDescriptor

var file_proto_news_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x0b, 0x4e, 0x65, 0x77,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x0c, 0x4e, 0x65, 0x77,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x42, 0x6f, 0x64, 0x79, 0x32, 0x3b, 0x0a, 0x04, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x33, 0x0a,
	0x04, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65,
	0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_news_proto_rawDescOnce sync.Once
	file_proto_news_proto_rawDescData = file_proto_news_proto_rawDesc
)

func file_proto_news_proto_rawDescGZIP() []byte {
	file_proto_news_proto_rawDescOnce.Do(func() {
		file_proto_news_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_news_proto_rawDescData)
	})
	return file_proto_news_proto_rawDescData
}

var file_proto_news_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_news_proto_goTypes = []interface{}{
	(*NewsRequest)(nil),  // 0: proto.NewsRequest
	(*NewsResponse)(nil), // 1: proto.NewsResponse
}
var file_proto_news_proto_depIdxs = []int32{
	0, // 0: proto.News.News:input_type -> proto.NewsRequest
	1, // 1: proto.News.News:output_type -> proto.NewsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_news_proto_init() }
func file_proto_news_proto_init() {
	if File_proto_news_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_news_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsRequest); i {
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
		file_proto_news_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsResponse); i {
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
			RawDescriptor: file_proto_news_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_news_proto_goTypes,
		DependencyIndexes: file_proto_news_proto_depIdxs,
		MessageInfos:      file_proto_news_proto_msgTypes,
	}.Build()
	File_proto_news_proto = out.File
	file_proto_news_proto_rawDesc = nil
	file_proto_news_proto_goTypes = nil
	file_proto_news_proto_depIdxs = nil
}
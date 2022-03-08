// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: svcpodserv.proto

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svcpodserv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_svcpodserv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_svcpodserv_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svcpodserv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_svcpodserv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_svcpodserv_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_svcpodserv_proto protoreflect.FileDescriptor

var file_svcpodserv_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x76, 0x63, 0x70, 0x6f, 0x64, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x73, 0x76, 0x63, 0x70, 0x6f, 0x64, 0x73, 0x65, 0x72, 0x76, 0x22, 0x23,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x47, 0x0a, 0x0a, 0x53, 0x76, 0x63,
	0x50, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x12, 0x39, 0x0a, 0x0a, 0x41, 0x73, 0x6b, 0x46, 0x6f,
	0x72, 0x48, 0x65, 0x6c, 0x70, 0x12, 0x13, 0x2e, 0x73, 0x76, 0x63, 0x70, 0x6f, 0x64, 0x73, 0x65,
	0x72, 0x76, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x76, 0x63,
	0x70, 0x6f, 0x64, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svcpodserv_proto_rawDescOnce sync.Once
	file_svcpodserv_proto_rawDescData = file_svcpodserv_proto_rawDesc
)

func file_svcpodserv_proto_rawDescGZIP() []byte {
	file_svcpodserv_proto_rawDescOnce.Do(func() {
		file_svcpodserv_proto_rawDescData = protoimpl.X.CompressGZIP(file_svcpodserv_proto_rawDescData)
	})
	return file_svcpodserv_proto_rawDescData
}

var file_svcpodserv_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_svcpodserv_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: svcpodserv.Request
	(*Response)(nil), // 1: svcpodserv.Response
}
var file_svcpodserv_proto_depIdxs = []int32{
	0, // 0: svcpodserv.SvcPodServ.AskForHelp:input_type -> svcpodserv.Request
	1, // 1: svcpodserv.SvcPodServ.AskForHelp:output_type -> svcpodserv.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_svcpodserv_proto_init() }
func file_svcpodserv_proto_init() {
	if File_svcpodserv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svcpodserv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_svcpodserv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_svcpodserv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svcpodserv_proto_goTypes,
		DependencyIndexes: file_svcpodserv_proto_depIdxs,
		MessageInfos:      file_svcpodserv_proto_msgTypes,
	}.Build()
	File_svcpodserv_proto = out.File
	file_svcpodserv_proto_rawDesc = nil
	file_svcpodserv_proto_goTypes = nil
	file_svcpodserv_proto_depIdxs = nil
}

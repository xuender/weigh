// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: pb/request.proto

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

// Request.
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id.
	// Deprecated: discard.
	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// Uri.
	// Deprecated: use URL.
	Uri  string  `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	Head []*Head `protobuf:"bytes,4,rep,name=head,proto3" json:"head,omitempty"`
	// Entity.
	// Deprecated: use Body.
	Entity []byte `protobuf:"bytes,5,opt,name=entity,proto3" json:"entity,omitempty"`
	URL    string `protobuf:"bytes,6,opt,name=URL,proto3" json:"URL,omitempty"`
	Body   []byte `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_pb_request_proto_msgTypes[0]
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
	return file_pb_request_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Request) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Request) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Request) GetHead() []*Head {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *Request) GetEntity() []byte {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *Request) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *Request) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_pb_request_proto protoreflect.FileDescriptor

var file_pb_request_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0d, 0x70, 0x62, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x1c, 0x0a, 0x04, 0x68,
	0x65, 0x61, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x55, 0x52, 0x4c, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_request_proto_rawDescOnce sync.Once
	file_pb_request_proto_rawDescData = file_pb_request_proto_rawDesc
)

func file_pb_request_proto_rawDescGZIP() []byte {
	file_pb_request_proto_rawDescOnce.Do(func() {
		file_pb_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_request_proto_rawDescData)
	})
	return file_pb_request_proto_rawDescData
}

var file_pb_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_request_proto_goTypes = []interface{}{
	(*Request)(nil), // 0: pb.Request
	(*Head)(nil),    // 1: pb.Head
}
var file_pb_request_proto_depIdxs = []int32{
	1, // 0: pb.Request.head:type_name -> pb.Head
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_request_proto_init() }
func file_pb_request_proto_init() {
	if File_pb_request_proto != nil {
		return
	}
	file_pb_head_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_request_proto_goTypes,
		DependencyIndexes: file_pb_request_proto_depIdxs,
		MessageInfos:      file_pb_request_proto_msgTypes,
	}.Build()
	File_pb_request_proto = out.File
	file_pb_request_proto_rawDesc = nil
	file_pb_request_proto_goTypes = nil
	file_pb_request_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/v1/error.proto

package apiv1

import (
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// Any error information beyond code and status should be included here and
// added to the error in the status details field. The frontend knows how to
// render all of the fields in a user-friendly way. If there is extremely
// verbose error information, consider adding it using a different type, e.g.
// from the errdetails package. Any details not using this type will still be
// accessible to the user in a raw format.
type ErrorDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If there are any underlying errors that were being wrapped, they are
	// presented here.
	Wrapped []*status.Status `protobuf:"bytes,1,rep,name=wrapped,proto3" json:"wrapped,omitempty"`
}

func (x *ErrorDetails) Reset() {
	*x = ErrorDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDetails) ProtoMessage() {}

func (x *ErrorDetails) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDetails.ProtoReflect.Descriptor instead.
func (*ErrorDetails) Descriptor() ([]byte, []int) {
	return file_api_v1_error_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorDetails) GetWrapped() []*status.Status {
	if x != nil {
		return x.Wrapped
	}
	return nil
}

var File_api_v1_error_proto protoreflect.FileDescriptor

var file_api_v1_error_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a,
	0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2c, 0x0a,
	0x07, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x07, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x42, 0xaa, 0x01, 0x0a, 0x12,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x42, 0x0a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x2e, 0x67, 0x6f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x2e, 0x69, 0x6f,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x41, 0x41, 0x58, 0xaa, 0x02, 0x0e, 0x41, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c,
	0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x41, 0x64, 0x6d, 0x69, 0x72, 0x61,
	0x6c, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x41, 0x64, 0x6d, 0x69, 0x72,
	0x61, 0x6c, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x41, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x3a,
	0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_error_proto_rawDescOnce sync.Once
	file_api_v1_error_proto_rawDescData = file_api_v1_error_proto_rawDesc
)

func file_api_v1_error_proto_rawDescGZIP() []byte {
	file_api_v1_error_proto_rawDescOnce.Do(func() {
		file_api_v1_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_error_proto_rawDescData)
	})
	return file_api_v1_error_proto_rawDescData
}

var file_api_v1_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_v1_error_proto_goTypes = []interface{}{
	(*ErrorDetails)(nil),  // 0: admiral.api.v1.ErrorDetails
	(*status.Status)(nil), // 1: google.rpc.Status
}
var file_api_v1_error_proto_depIdxs = []int32{
	1, // 0: admiral.api.v1.ErrorDetails.wrapped:type_name -> google.rpc.Status
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_v1_error_proto_init() }
func file_api_v1_error_proto_init() {
	if File_api_v1_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorDetails); i {
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
			RawDescriptor: file_api_v1_error_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_error_proto_goTypes,
		DependencyIndexes: file_api_v1_error_proto_depIdxs,
		MessageInfos:      file_api_v1_error_proto_msgTypes,
	}.Build()
	File_api_v1_error_proto = out.File
	file_api_v1_error_proto_rawDesc = nil
	file_api_v1_error_proto_goTypes = nil
	file_api_v1_error_proto_depIdxs = nil
}

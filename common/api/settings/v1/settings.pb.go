// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: settings/v1/settings.proto

package settingsv1

import (
	_ "go.datalift.io/admiral/common/api/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type OIDCConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Issuer      string   `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	ClientId    string   `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	CliClientId string   `protobuf:"bytes,4,opt,name=cli_client_id,json=cliClientId,proto3" json:"cli_client_id,omitempty"`
	Scopes      []string `protobuf:"bytes,5,rep,name=scopes,proto3" json:"scopes,omitempty"`
}

func (x *OIDCConfig) Reset() {
	*x = OIDCConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_v1_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OIDCConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OIDCConfig) ProtoMessage() {}

func (x *OIDCConfig) ProtoReflect() protoreflect.Message {
	mi := &file_settings_v1_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OIDCConfig.ProtoReflect.Descriptor instead.
func (*OIDCConfig) Descriptor() ([]byte, []int) {
	return file_settings_v1_settings_proto_rawDescGZIP(), []int{0}
}

func (x *OIDCConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OIDCConfig) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *OIDCConfig) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *OIDCConfig) GetCliClientId() string {
	if x != nil {
		return x.CliClientId
	}
	return ""
}

func (x *OIDCConfig) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

type SettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SettingsRequest) Reset() {
	*x = SettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_v1_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingsRequest) ProtoMessage() {}

func (x *SettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_settings_v1_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingsRequest.ProtoReflect.Descriptor instead.
func (*SettingsRequest) Descriptor() ([]byte, []int) {
	return file_settings_v1_settings_proto_rawDescGZIP(), []int{1}
}

type SettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url        string      `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	OidcConfig *OIDCConfig `protobuf:"bytes,2,opt,name=oidc_config,json=oidcConfig,proto3" json:"oidc_config,omitempty"`
}

func (x *SettingsResponse) Reset() {
	*x = SettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_v1_settings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingsResponse) ProtoMessage() {}

func (x *SettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_settings_v1_settings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingsResponse.ProtoReflect.Descriptor instead.
func (*SettingsResponse) Descriptor() ([]byte, []int) {
	return file_settings_v1_settings_proto_rawDescGZIP(), []int{2}
}

func (x *SettingsResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SettingsResponse) GetOidcConfig() *OIDCConfig {
	if x != nil {
		return x.OidcConfig
	}
	return nil
}

var File_settings_v1_settings_proto protoreflect.FileDescriptor

var file_settings_v1_settings_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x64,
	0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a,
	0x0a, 0x4f, 0x49, 0x44, 0x43, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x69,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73,
	0x22, 0x11, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x66, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x40, 0x0a, 0x0b, 0x6f, 0x69, 0x64,
	0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x49, 0x44, 0x43, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0a, 0x6f, 0x69, 0x64, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0x86, 0x01, 0x0a, 0x0b,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x41, 0x50, 0x49, 0x12, 0x77, 0x0a, 0x08, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x24, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61,
	0x6c, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0xaa, 0xe1, 0x1c, 0x02, 0x08, 0x02, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x42, 0xd0, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31,
	0x42, 0x0d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x38, 0x67, 0x6f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x2e, 0x69,
	0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x53,
	0x58, 0xaa, 0x02, 0x13, 0x41, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x13, 0x41, 0x64, 0x6d, 0x69, 0x72, 0x61,
	0x6c, 0x5c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1f,
	0x41, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x5c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x15, 0x41, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x3a, 0x3a, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_settings_v1_settings_proto_rawDescOnce sync.Once
	file_settings_v1_settings_proto_rawDescData = file_settings_v1_settings_proto_rawDesc
)

func file_settings_v1_settings_proto_rawDescGZIP() []byte {
	file_settings_v1_settings_proto_rawDescOnce.Do(func() {
		file_settings_v1_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_settings_v1_settings_proto_rawDescData)
	})
	return file_settings_v1_settings_proto_rawDescData
}

var file_settings_v1_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_settings_v1_settings_proto_goTypes = []interface{}{
	(*OIDCConfig)(nil),       // 0: admiral.settings.v1.OIDCConfig
	(*SettingsRequest)(nil),  // 1: admiral.settings.v1.SettingsRequest
	(*SettingsResponse)(nil), // 2: admiral.settings.v1.SettingsResponse
}
var file_settings_v1_settings_proto_depIdxs = []int32{
	0, // 0: admiral.settings.v1.SettingsResponse.oidc_config:type_name -> admiral.settings.v1.OIDCConfig
	1, // 1: admiral.settings.v1.SettingsAPI.Settings:input_type -> admiral.settings.v1.SettingsRequest
	2, // 2: admiral.settings.v1.SettingsAPI.Settings:output_type -> admiral.settings.v1.SettingsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_settings_v1_settings_proto_init() }
func file_settings_v1_settings_proto_init() {
	if File_settings_v1_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_settings_v1_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OIDCConfig); i {
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
		file_settings_v1_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettingsRequest); i {
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
		file_settings_v1_settings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettingsResponse); i {
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
			RawDescriptor: file_settings_v1_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_settings_v1_settings_proto_goTypes,
		DependencyIndexes: file_settings_v1_settings_proto_depIdxs,
		MessageInfos:      file_settings_v1_settings_proto_msgTypes,
	}.Build()
	File_settings_v1_settings_proto = out.File
	file_settings_v1_settings_proto_rawDesc = nil
	file_settings_v1_settings_proto_goTypes = nil
	file_settings_v1_settings_proto_depIdxs = nil
}

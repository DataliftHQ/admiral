// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: authn/v1/authn.proto

package authnv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "go.datalift.io/admiral/common/api/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateTokenRequest_TokenType int32

const (
	CreateTokenRequest_UNSPECIFIED CreateTokenRequest_TokenType = 0
	// Token is issued on behalf of another service for programmatic access.
	CreateTokenRequest_SERVICE CreateTokenRequest_TokenType = 1
)

// Enum value maps for CreateTokenRequest_TokenType.
var (
	CreateTokenRequest_TokenType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "SERVICE",
	}
	CreateTokenRequest_TokenType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"SERVICE":     1,
	}
)

func (x CreateTokenRequest_TokenType) Enum() *CreateTokenRequest_TokenType {
	p := new(CreateTokenRequest_TokenType)
	*p = x
	return p
}

func (x CreateTokenRequest_TokenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateTokenRequest_TokenType) Descriptor() protoreflect.EnumDescriptor {
	return file_authn_v1_authn_proto_enumTypes[0].Descriptor()
}

func (CreateTokenRequest_TokenType) Type() protoreflect.EnumType {
	return &file_authn_v1_authn_proto_enumTypes[0]
}

func (x CreateTokenRequest_TokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateTokenRequest_TokenType.Descriptor instead.
func (CreateTokenRequest_TokenType) EnumDescriptor() ([]byte, []int) {
	return file_authn_v1_authn_proto_rawDescGZIP(), []int{4, 0}
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedirectUrl string `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_v1_authn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authn_v1_authn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_authn_v1_authn_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If the user presents a refresh token and refresh succeeds, they will receive the token in response. Otherwise,
	// they will be redirected through the authentication flow.
	//
	// Types that are assignable to Return:
	//
	//	*LoginResponse_AuthUrl
	//	*LoginResponse_Token_
	Return isLoginResponse_Return `protobuf_oneof:"return"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_v1_authn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authn_v1_authn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_authn_v1_authn_proto_rawDescGZIP(), []int{1}
}

func (m *LoginResponse) GetReturn() isLoginResponse_Return {
	if m != nil {
		return m.Return
	}
	return nil
}

func (x *LoginResponse) GetAuthUrl() string {
	if x, ok := x.GetReturn().(*LoginResponse_AuthUrl); ok {
		return x.AuthUrl
	}
	return ""
}

func (x *LoginResponse) GetToken() *LoginResponse_Token {
	if x, ok := x.GetReturn().(*LoginResponse_Token_); ok {
		return x.Token
	}
	return nil
}

type isLoginResponse_Return interface {
	isLoginResponse_Return()
}

type LoginResponse_AuthUrl struct {
	AuthUrl string `protobuf:"bytes,1,opt,name=auth_url,json=authUrl,proto3,oneof"`
}

type LoginResponse_Token_ struct {
	Token *LoginResponse_Token `protobuf:"bytes,2,opt,name=token,proto3,oneof"`
}

func (*LoginResponse_AuthUrl) isLoginResponse_Return() {}

func (*LoginResponse_Token_) isLoginResponse_Return() {}

// See https://www.oauth.com/oauth2-servers/authorization/the-authorization-response/ for description of the parameters.
type CallbackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code             string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	State            string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Error            string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	ErrorDescription string `protobuf:"bytes,4,opt,name=error_description,json=errorDescription,proto3" json:"error_description,omitempty"`
}

func (x *CallbackRequest) Reset() {
	*x = CallbackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_v1_authn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackRequest) ProtoMessage() {}

func (x *CallbackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authn_v1_authn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackRequest.ProtoReflect.Descriptor instead.
func (*CallbackRequest) Descriptor() ([]byte, []int) {
	return file_authn_v1_authn_proto_rawDescGZIP(), []int{2}
}

func (x *CallbackRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CallbackRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *CallbackRequest) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CallbackRequest) GetErrorDescription() string {
	if x != nil {
		return x.ErrorDescription
	}
	return ""
}

type CallbackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is the token that the user should present. Note: this response is only valid in a gRPC context. In an HTTP
	// context the user will be redirected.
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// The refresh token will be empty if no refresh token was issued.
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *CallbackResponse) Reset() {
	*x = CallbackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_v1_authn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackResponse) ProtoMessage() {}

func (x *CallbackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authn_v1_authn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackResponse.ProtoReflect.Descriptor instead.
func (*CallbackResponse) Descriptor() ([]byte, []int) {
	return file_authn_v1_authn_proto_rawDescGZIP(), []int{3}
}

func (x *CallbackResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *CallbackResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type CreateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The subject to issue this token for.
	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	// The duration until this token should expire. If unset, the token will never expire.
	Expiry *durationpb.Duration `protobuf:"bytes,2,opt,name=expiry,proto3" json:"expiry,omitempty"`
	// The kind of token to issue. This provides namespacing to avoid naming collisions.
	TokenType CreateTokenRequest_TokenType `protobuf:"varint,3,opt,name=token_type,json=tokenType,proto3,enum=admiral.authn.v1.CreateTokenRequest_TokenType" json:"token_type,omitempty"`
}

func (x *CreateTokenRequest) Reset() {
	*x = CreateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_v1_authn_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenRequest) ProtoMessage() {}

func (x *CreateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authn_v1_authn_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenRequest.ProtoReflect.Descriptor instead.
func (*CreateTokenRequest) Descriptor() ([]byte, []int) {
	return file_authn_v1_authn_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTokenRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *CreateTokenRequest) GetExpiry() *durationpb.Duration {
	if x != nil {
		return x.Expiry
	}
	return nil
}

func (x *CreateTokenRequest) GetTokenType() CreateTokenRequest_TokenType {
	if x != nil {
		return x.TokenType
	}
	return CreateTokenRequest_UNSPECIFIED
}

type CreateTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The access token associated with the newly created token.
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *CreateTokenResponse) Reset() {
	*x = CreateTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_v1_authn_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenResponse) ProtoMessage() {}

func (x *CreateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authn_v1_authn_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenResponse.ProtoReflect.Descriptor instead.
func (*CreateTokenResponse) Descriptor() ([]byte, []int) {
	return file_authn_v1_authn_proto_rawDescGZIP(), []int{5}
}

func (x *CreateTokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type LoginResponse_Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *LoginResponse_Token) Reset() {
	*x = LoginResponse_Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_v1_authn_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse_Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse_Token) ProtoMessage() {}

func (x *LoginResponse_Token) ProtoReflect() protoreflect.Message {
	mi := &file_authn_v1_authn_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse_Token.ProtoReflect.Descriptor instead.
func (*LoginResponse_Token) Descriptor() ([]byte, []int) {
	return file_authn_v1_authn_proto_rawDescGZIP(), []int{1, 0}
}

func (x *LoginResponse_Token) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginResponse_Token) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

var File_authn_v1_authn_proto protoreflect.FileDescriptor

var file_authn_v1_authn_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x0c,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x3a,
	0x04, 0xb8, 0xe1, 0x1c, 0x01, 0x22, 0xcc, 0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x61, 0x75, 0x74,
	0x68, 0x55, 0x72, 0x6c, 0x12, 0x3d, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x48, 0x00, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x1a, 0x4f, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x3a, 0x04, 0xb8, 0xe1, 0x1c, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x22, 0x84, 0x01, 0x0a, 0x0f, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x04, 0xb8, 0xe1, 0x1c, 0x01, 0x22, 0x60, 0x0a, 0x10, 0x43,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x3a, 0x04, 0xb8, 0xe1, 0x1c, 0x01, 0x22, 0xf6, 0x01,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x59, 0x0a, 0x0a, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0a,
	0xfa, 0x42, 0x07, 0x82, 0x01, 0x04, 0x10, 0x01, 0x20, 0x00, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x29, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x01,
	0x3a, 0x04, 0xb8, 0xe1, 0x1c, 0x01, 0x22, 0x3e, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x3a, 0x04, 0xb8, 0xe1, 0x1c, 0x01, 0x32, 0xe9, 0x02, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6e,
	0x41, 0x50, 0x49, 0x12, 0x63, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1e, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0xaa,
	0xe1, 0x1c, 0x02, 0x08, 0x02, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x6f, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x12, 0x21, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0xaa, 0xe1, 0x1c,
	0x02, 0x08, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x86, 0x01, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x24, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x72, 0x61, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0xaa, 0xe1, 0x1c, 0x02, 0x08, 0x01, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x42, 0xb8, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x72,
	0x61, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x41, 0x75, 0x74,
	0x68, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x6f, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61,
	0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x41, 0x41, 0x58, 0xaa, 0x02, 0x10, 0x41, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x41, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c,
	0x5c, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x41, 0x64, 0x6d, 0x69,
	0x72, 0x61, 0x6c, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x41, 0x64, 0x6d, 0x69, 0x72,
	0x61, 0x6c, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authn_v1_authn_proto_rawDescOnce sync.Once
	file_authn_v1_authn_proto_rawDescData = file_authn_v1_authn_proto_rawDesc
)

func file_authn_v1_authn_proto_rawDescGZIP() []byte {
	file_authn_v1_authn_proto_rawDescOnce.Do(func() {
		file_authn_v1_authn_proto_rawDescData = protoimpl.X.CompressGZIP(file_authn_v1_authn_proto_rawDescData)
	})
	return file_authn_v1_authn_proto_rawDescData
}

var file_authn_v1_authn_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_authn_v1_authn_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_authn_v1_authn_proto_goTypes = []interface{}{
	(CreateTokenRequest_TokenType)(0), // 0: admiral.authn.v1.CreateTokenRequest.TokenType
	(*LoginRequest)(nil),              // 1: admiral.authn.v1.LoginRequest
	(*LoginResponse)(nil),             // 2: admiral.authn.v1.LoginResponse
	(*CallbackRequest)(nil),           // 3: admiral.authn.v1.CallbackRequest
	(*CallbackResponse)(nil),          // 4: admiral.authn.v1.CallbackResponse
	(*CreateTokenRequest)(nil),        // 5: admiral.authn.v1.CreateTokenRequest
	(*CreateTokenResponse)(nil),       // 6: admiral.authn.v1.CreateTokenResponse
	(*LoginResponse_Token)(nil),       // 7: admiral.authn.v1.LoginResponse.Token
	(*durationpb.Duration)(nil),       // 8: google.protobuf.Duration
}
var file_authn_v1_authn_proto_depIdxs = []int32{
	7, // 0: admiral.authn.v1.LoginResponse.token:type_name -> admiral.authn.v1.LoginResponse.Token
	8, // 1: admiral.authn.v1.CreateTokenRequest.expiry:type_name -> google.protobuf.Duration
	0, // 2: admiral.authn.v1.CreateTokenRequest.token_type:type_name -> admiral.authn.v1.CreateTokenRequest.TokenType
	1, // 3: admiral.authn.v1.AuthnAPI.Login:input_type -> admiral.authn.v1.LoginRequest
	3, // 4: admiral.authn.v1.AuthnAPI.Callback:input_type -> admiral.authn.v1.CallbackRequest
	5, // 5: admiral.authn.v1.AuthnAPI.CreateToken:input_type -> admiral.authn.v1.CreateTokenRequest
	2, // 6: admiral.authn.v1.AuthnAPI.Login:output_type -> admiral.authn.v1.LoginResponse
	4, // 7: admiral.authn.v1.AuthnAPI.Callback:output_type -> admiral.authn.v1.CallbackResponse
	6, // 8: admiral.authn.v1.AuthnAPI.CreateToken:output_type -> admiral.authn.v1.CreateTokenResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_authn_v1_authn_proto_init() }
func file_authn_v1_authn_proto_init() {
	if File_authn_v1_authn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authn_v1_authn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_authn_v1_authn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_authn_v1_authn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallbackRequest); i {
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
		file_authn_v1_authn_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallbackResponse); i {
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
		file_authn_v1_authn_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTokenRequest); i {
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
		file_authn_v1_authn_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTokenResponse); i {
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
		file_authn_v1_authn_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse_Token); i {
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
	file_authn_v1_authn_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*LoginResponse_AuthUrl)(nil),
		(*LoginResponse_Token_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_authn_v1_authn_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authn_v1_authn_proto_goTypes,
		DependencyIndexes: file_authn_v1_authn_proto_depIdxs,
		EnumInfos:         file_authn_v1_authn_proto_enumTypes,
		MessageInfos:      file_authn_v1_authn_proto_msgTypes,
	}.Build()
	File_authn_v1_authn_proto = out.File
	file_authn_v1_authn_proto_rawDesc = nil
	file_authn_v1_authn_proto_goTypes = nil
	file_authn_v1_authn_proto_depIdxs = nil
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: settings/v1/settings.proto

package settingsv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SettingsAPI_Settings_FullMethodName = "/admiral.settings.v1.SettingsAPI/Settings"
)

// SettingsAPIClient is the client API for SettingsAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SettingsAPIClient interface {
	Settings(ctx context.Context, in *SettingsRequest, opts ...grpc.CallOption) (*SettingsResponse, error)
}

type settingsAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewSettingsAPIClient(cc grpc.ClientConnInterface) SettingsAPIClient {
	return &settingsAPIClient{cc}
}

func (c *settingsAPIClient) Settings(ctx context.Context, in *SettingsRequest, opts ...grpc.CallOption) (*SettingsResponse, error) {
	out := new(SettingsResponse)
	err := c.cc.Invoke(ctx, SettingsAPI_Settings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingsAPIServer is the server API for SettingsAPI service.
// All implementations should embed UnimplementedSettingsAPIServer
// for forward compatibility
type SettingsAPIServer interface {
	Settings(context.Context, *SettingsRequest) (*SettingsResponse, error)
}

// UnimplementedSettingsAPIServer should be embedded to have forward compatible implementations.
type UnimplementedSettingsAPIServer struct {
}

func (UnimplementedSettingsAPIServer) Settings(context.Context, *SettingsRequest) (*SettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Settings not implemented")
}

// UnsafeSettingsAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SettingsAPIServer will
// result in compilation errors.
type UnsafeSettingsAPIServer interface {
	mustEmbedUnimplementedSettingsAPIServer()
}

func RegisterSettingsAPIServer(s grpc.ServiceRegistrar, srv SettingsAPIServer) {
	s.RegisterService(&SettingsAPI_ServiceDesc, srv)
}

func _SettingsAPI_Settings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsAPIServer).Settings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsAPI_Settings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsAPIServer).Settings(ctx, req.(*SettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SettingsAPI_ServiceDesc is the grpc.ServiceDesc for SettingsAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SettingsAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admiral.settings.v1.SettingsAPI",
	HandlerType: (*SettingsAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Settings",
			Handler:    _SettingsAPI_Settings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "settings/v1/settings.proto",
}

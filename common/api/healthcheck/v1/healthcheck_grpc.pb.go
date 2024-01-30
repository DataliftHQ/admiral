// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: healthcheck/v1/healthcheck.proto

package healthcheckv1

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
	HealthcheckAPI_Healthcheck_FullMethodName = "/admiral.healthcheck.v1.HealthcheckAPI/Healthcheck"
)

// HealthcheckAPIClient is the client API for HealthcheckAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthcheckAPIClient interface {
	Healthcheck(ctx context.Context, in *HealthcheckRequest, opts ...grpc.CallOption) (*HealthcheckResponse, error)
}

type healthcheckAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthcheckAPIClient(cc grpc.ClientConnInterface) HealthcheckAPIClient {
	return &healthcheckAPIClient{cc}
}

func (c *healthcheckAPIClient) Healthcheck(ctx context.Context, in *HealthcheckRequest, opts ...grpc.CallOption) (*HealthcheckResponse, error) {
	out := new(HealthcheckResponse)
	err := c.cc.Invoke(ctx, HealthcheckAPI_Healthcheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthcheckAPIServer is the server API for HealthcheckAPI service.
// All implementations should embed UnimplementedHealthcheckAPIServer
// for forward compatibility
type HealthcheckAPIServer interface {
	Healthcheck(context.Context, *HealthcheckRequest) (*HealthcheckResponse, error)
}

// UnimplementedHealthcheckAPIServer should be embedded to have forward compatible implementations.
type UnimplementedHealthcheckAPIServer struct {
}

func (UnimplementedHealthcheckAPIServer) Healthcheck(context.Context, *HealthcheckRequest) (*HealthcheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Healthcheck not implemented")
}

// UnsafeHealthcheckAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthcheckAPIServer will
// result in compilation errors.
type UnsafeHealthcheckAPIServer interface {
	mustEmbedUnimplementedHealthcheckAPIServer()
}

func RegisterHealthcheckAPIServer(s grpc.ServiceRegistrar, srv HealthcheckAPIServer) {
	s.RegisterService(&HealthcheckAPI_ServiceDesc, srv)
}

func _HealthcheckAPI_Healthcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthcheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthcheckAPIServer).Healthcheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthcheckAPI_Healthcheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthcheckAPIServer).Healthcheck(ctx, req.(*HealthcheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthcheckAPI_ServiceDesc is the grpc.ServiceDesc for HealthcheckAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthcheckAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admiral.healthcheck.v1.HealthcheckAPI",
	HandlerType: (*HealthcheckAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Healthcheck",
			Handler:    _HealthcheckAPI_Healthcheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "healthcheck/v1/healthcheck.proto",
}

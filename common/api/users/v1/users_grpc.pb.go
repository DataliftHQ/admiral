// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: users/v1/users.proto

package usersv1

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
	UsersAPI_GetMe_FullMethodName   = "/admiral.users.v1.UsersAPI/GetMe"
	UsersAPI_GetUser_FullMethodName = "/admiral.users.v1.UsersAPI/GetUser"
)

// UsersAPIClient is the client API for UsersAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersAPIClient interface {
	GetMe(ctx context.Context, in *GetMeRequest, opts ...grpc.CallOption) (*GetMeResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
}

type usersAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersAPIClient(cc grpc.ClientConnInterface) UsersAPIClient {
	return &usersAPIClient{cc}
}

func (c *usersAPIClient) GetMe(ctx context.Context, in *GetMeRequest, opts ...grpc.CallOption) (*GetMeResponse, error) {
	out := new(GetMeResponse)
	err := c.cc.Invoke(ctx, UsersAPI_GetMe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersAPIClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, UsersAPI_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersAPIServer is the server API for UsersAPI service.
// All implementations should embed UnimplementedUsersAPIServer
// for forward compatibility
type UsersAPIServer interface {
	GetMe(context.Context, *GetMeRequest) (*GetMeResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
}

// UnimplementedUsersAPIServer should be embedded to have forward compatible implementations.
type UnimplementedUsersAPIServer struct {
}

func (UnimplementedUsersAPIServer) GetMe(context.Context, *GetMeRequest) (*GetMeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMe not implemented")
}
func (UnimplementedUsersAPIServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

// UnsafeUsersAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersAPIServer will
// result in compilation errors.
type UnsafeUsersAPIServer interface {
	mustEmbedUnimplementedUsersAPIServer()
}

func RegisterUsersAPIServer(s grpc.ServiceRegistrar, srv UsersAPIServer) {
	s.RegisterService(&UsersAPI_ServiceDesc, srv)
}

func _UsersAPI_GetMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersAPIServer).GetMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersAPI_GetMe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersAPIServer).GetMe(ctx, req.(*GetMeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersAPI_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersAPIServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersAPI_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersAPIServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UsersAPI_ServiceDesc is the grpc.ServiceDesc for UsersAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsersAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admiral.users.v1.UsersAPI",
	HandlerType: (*UsersAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMe",
			Handler:    _UsersAPI_GetMe_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UsersAPI_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users/v1/users.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: userserver.proto

package userserver

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

// UserServerClient is the client API for UserServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServerClient interface {
	// Add new user
	AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	// Delete user
	DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	// Get list of all users
	GetUsers(ctx context.Context, in *GetUsersParams, opts ...grpc.CallOption) (*Users, error)
}

type userServerClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServerClient(cc grpc.ClientConnInterface) UserServerClient {
	return &userServerClient{cc}
}

func (c *userServerClient) AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/userserver.UserServer/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServerClient) DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/userserver.UserServer/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServerClient) GetUsers(ctx context.Context, in *GetUsersParams, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/userserver.UserServer/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServerServer is the server API for UserServer service.
// All implementations must embed UnimplementedUserServerServer
// for forward compatibility
type UserServerServer interface {
	// Add new user
	AddUser(context.Context, *User) (*Response, error)
	// Delete user
	DeleteUser(context.Context, *User) (*Response, error)
	// Get list of all users
	GetUsers(context.Context, *GetUsersParams) (*Users, error)
	mustEmbedUnimplementedUserServerServer()
}

// UnimplementedUserServerServer must be embedded to have forward compatible implementations.
type UnimplementedUserServerServer struct {
}

func (UnimplementedUserServerServer) AddUser(context.Context, *User) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedUserServerServer) DeleteUser(context.Context, *User) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServerServer) GetUsers(context.Context, *GetUsersParams) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUserServerServer) mustEmbedUnimplementedUserServerServer() {}

// UnsafeUserServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServerServer will
// result in compilation errors.
type UnsafeUserServerServer interface {
	mustEmbedUnimplementedUserServerServer()
}

func RegisterUserServerServer(s grpc.ServiceRegistrar, srv UserServerServer) {
	s.RegisterService(&UserServer_ServiceDesc, srv)
}

func _UserServer_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userserver.UserServer/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).AddUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServer_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userserver.UserServer/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).DeleteUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServer_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userserver.UserServer/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).GetUsers(ctx, req.(*GetUsersParams))
	}
	return interceptor(ctx, in, info, handler)
}

// UserServer_ServiceDesc is the grpc.ServiceDesc for UserServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userserver.UserServer",
	HandlerType: (*UserServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _UserServer_AddUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserServer_DeleteUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _UserServer_GetUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userserver.proto",
}

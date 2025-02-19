// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: schema/v1/user_schema.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	UserSchema_ListUser_FullMethodName             = "/v1.UserSchema/ListUser"
	UserSchema_CreateUserPreference_FullMethodName = "/v1.UserSchema/CreateUserPreference"
	UserSchema_Me_FullMethodName                   = "/v1.UserSchema/Me"
)

// UserSchemaClient is the client API for UserSchema service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserSchemaClient interface {
	ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error)
	CreateUserPreference(ctx context.Context, in *CreateUserPreferenceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Me(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserResponse, error)
}

type userSchemaClient struct {
	cc grpc.ClientConnInterface
}

func NewUserSchemaClient(cc grpc.ClientConnInterface) UserSchemaClient {
	return &userSchemaClient{cc}
}

func (c *userSchemaClient) ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUserResponse)
	err := c.cc.Invoke(ctx, UserSchema_ListUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSchemaClient) CreateUserPreference(ctx context.Context, in *CreateUserPreferenceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserSchema_CreateUserPreference_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSchemaClient) Me(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserSchema_Me_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSchemaServer is the server API for UserSchema service.
// All implementations must embed UnimplementedUserSchemaServer
// for forward compatibility
type UserSchemaServer interface {
	ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error)
	CreateUserPreference(context.Context, *CreateUserPreferenceRequest) (*emptypb.Empty, error)
	Me(context.Context, *emptypb.Empty) (*UserResponse, error)
	mustEmbedUnimplementedUserSchemaServer()
}

// UnimplementedUserSchemaServer must be embedded to have forward compatible implementations.
type UnimplementedUserSchemaServer struct {
}

func (UnimplementedUserSchemaServer) ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedUserSchemaServer) CreateUserPreference(context.Context, *CreateUserPreferenceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserPreference not implemented")
}
func (UnimplementedUserSchemaServer) Me(context.Context, *emptypb.Empty) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Me not implemented")
}
func (UnimplementedUserSchemaServer) mustEmbedUnimplementedUserSchemaServer() {}

// UnsafeUserSchemaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserSchemaServer will
// result in compilation errors.
type UnsafeUserSchemaServer interface {
	mustEmbedUnimplementedUserSchemaServer()
}

func RegisterUserSchemaServer(s grpc.ServiceRegistrar, srv UserSchemaServer) {
	s.RegisterService(&UserSchema_ServiceDesc, srv)
}

func _UserSchema_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSchemaServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserSchema_ListUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSchemaServer).ListUser(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSchema_CreateUserPreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserPreferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSchemaServer).CreateUserPreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserSchema_CreateUserPreference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSchemaServer).CreateUserPreference(ctx, req.(*CreateUserPreferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSchema_Me_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSchemaServer).Me(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserSchema_Me_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSchemaServer).Me(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// UserSchema_ServiceDesc is the grpc.ServiceDesc for UserSchema service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserSchema_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.UserSchema",
	HandlerType: (*UserSchemaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUser",
			Handler:    _UserSchema_ListUser_Handler,
		},
		{
			MethodName: "CreateUserPreference",
			Handler:    _UserSchema_CreateUserPreference_Handler,
		},
		{
			MethodName: "Me",
			Handler:    _UserSchema_Me_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schema/v1/user_schema.proto",
}

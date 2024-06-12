// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: schema/v1/package_schema.proto

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
	PackageSchema_ListPackage_FullMethodName  = "/v1.PackageSchema/ListPackage"
	PackageSchema_OrderPackage_FullMethodName = "/v1.PackageSchema/OrderPackage"
)

// PackageSchemaClient is the client API for PackageSchema service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PackageSchemaClient interface {
	ListPackage(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListPackageResponse, error)
	OrderPackage(ctx context.Context, in *OrderPackageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type packageSchemaClient struct {
	cc grpc.ClientConnInterface
}

func NewPackageSchemaClient(cc grpc.ClientConnInterface) PackageSchemaClient {
	return &packageSchemaClient{cc}
}

func (c *packageSchemaClient) ListPackage(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListPackageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPackageResponse)
	err := c.cc.Invoke(ctx, PackageSchema_ListPackage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageSchemaClient) OrderPackage(ctx context.Context, in *OrderPackageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PackageSchema_OrderPackage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PackageSchemaServer is the server API for PackageSchema service.
// All implementations must embed UnimplementedPackageSchemaServer
// for forward compatibility
type PackageSchemaServer interface {
	ListPackage(context.Context, *emptypb.Empty) (*ListPackageResponse, error)
	OrderPackage(context.Context, *OrderPackageRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPackageSchemaServer()
}

// UnimplementedPackageSchemaServer must be embedded to have forward compatible implementations.
type UnimplementedPackageSchemaServer struct {
}

func (UnimplementedPackageSchemaServer) ListPackage(context.Context, *emptypb.Empty) (*ListPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPackage not implemented")
}
func (UnimplementedPackageSchemaServer) OrderPackage(context.Context, *OrderPackageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderPackage not implemented")
}
func (UnimplementedPackageSchemaServer) mustEmbedUnimplementedPackageSchemaServer() {}

// UnsafePackageSchemaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PackageSchemaServer will
// result in compilation errors.
type UnsafePackageSchemaServer interface {
	mustEmbedUnimplementedPackageSchemaServer()
}

func RegisterPackageSchemaServer(s grpc.ServiceRegistrar, srv PackageSchemaServer) {
	s.RegisterService(&PackageSchema_ServiceDesc, srv)
}

func _PackageSchema_ListPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageSchemaServer).ListPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PackageSchema_ListPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageSchemaServer).ListPackage(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageSchema_OrderPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageSchemaServer).OrderPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PackageSchema_OrderPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageSchemaServer).OrderPackage(ctx, req.(*OrderPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PackageSchema_ServiceDesc is the grpc.ServiceDesc for PackageSchema service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PackageSchema_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PackageSchema",
	HandlerType: (*PackageSchemaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPackage",
			Handler:    _PackageSchema_ListPackage_Handler,
		},
		{
			MethodName: "OrderPackage",
			Handler:    _PackageSchema_OrderPackage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schema/v1/package_schema.proto",
}

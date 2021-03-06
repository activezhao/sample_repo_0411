// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.5.0
// source: Zhaoh0411.proto

package sample_package_0411

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

// Zhaoh0411Client is the client API for Zhaoh0411 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Zhaoh0411Client interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type zhaoh0411Client struct {
	cc grpc.ClientConnInterface
}

func NewZhaoh0411Client(cc grpc.ClientConnInterface) Zhaoh0411Client {
	return &zhaoh0411Client{cc}
}

func (c *zhaoh0411Client) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/Zhaoh0411/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Zhaoh0411Server is the server API for Zhaoh0411 service.
// All implementations must embed UnimplementedZhaoh0411Server
// for forward compatibility
type Zhaoh0411Server interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedZhaoh0411Server()
}

// UnimplementedZhaoh0411Server must be embedded to have forward compatible implementations.
type UnimplementedZhaoh0411Server struct {
}

func (UnimplementedZhaoh0411Server) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedZhaoh0411Server) mustEmbedUnimplementedZhaoh0411Server() {}

// UnsafeZhaoh0411Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Zhaoh0411Server will
// result in compilation errors.
type UnsafeZhaoh0411Server interface {
	mustEmbedUnimplementedZhaoh0411Server()
}

func RegisterZhaoh0411Server(s grpc.ServiceRegistrar, srv Zhaoh0411Server) {
	s.RegisterService(&Zhaoh0411_ServiceDesc, srv)
}

func _Zhaoh0411_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Zhaoh0411Server).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Zhaoh0411/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Zhaoh0411Server).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Zhaoh0411_ServiceDesc is the grpc.ServiceDesc for Zhaoh0411 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Zhaoh0411_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Zhaoh0411",
	HandlerType: (*Zhaoh0411Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Zhaoh0411_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Zhaoh0411.proto",
}

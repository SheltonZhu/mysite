// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pbs

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

// BibleServiceClient is the client API for BibleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BibleServiceClient interface {
	Get(ctx context.Context, in *Int32Value, opts ...grpc.CallOption) (*Bible, error)
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Bibles, error)
	Create(ctx context.Context, in *StringValue, opts ...grpc.CallOption) (*Int32Value, error)
	Delete(ctx context.Context, in *Int32Value, opts ...grpc.CallOption) (*Int32Value, error)
}

type bibleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBibleServiceClient(cc grpc.ClientConnInterface) BibleServiceClient {
	return &bibleServiceClient{cc}
}

func (c *bibleServiceClient) Get(ctx context.Context, in *Int32Value, opts ...grpc.CallOption) (*Bible, error) {
	out := new(Bible)
	err := c.cc.Invoke(ctx, "/protos.BibleService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bibleServiceClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Bibles, error) {
	out := new(Bibles)
	err := c.cc.Invoke(ctx, "/protos.BibleService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bibleServiceClient) Create(ctx context.Context, in *StringValue, opts ...grpc.CallOption) (*Int32Value, error) {
	out := new(Int32Value)
	err := c.cc.Invoke(ctx, "/protos.BibleService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bibleServiceClient) Delete(ctx context.Context, in *Int32Value, opts ...grpc.CallOption) (*Int32Value, error) {
	out := new(Int32Value)
	err := c.cc.Invoke(ctx, "/protos.BibleService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BibleServiceServer is the server API for BibleService service.
// All implementations must embed UnimplementedBibleServiceServer
// for forward compatibility
type BibleServiceServer interface {
	Get(context.Context, *Int32Value) (*Bible, error)
	List(context.Context, *Empty) (*Bibles, error)
	Create(context.Context, *StringValue) (*Int32Value, error)
	Delete(context.Context, *Int32Value) (*Int32Value, error)
	mustEmbedUnimplementedBibleServiceServer()
}

// UnimplementedBibleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBibleServiceServer struct {
}

func (UnimplementedBibleServiceServer) Get(context.Context, *Int32Value) (*Bible, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBibleServiceServer) List(context.Context, *Empty) (*Bibles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedBibleServiceServer) Create(context.Context, *StringValue) (*Int32Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBibleServiceServer) Delete(context.Context, *Int32Value) (*Int32Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBibleServiceServer) mustEmbedUnimplementedBibleServiceServer() {}

// UnsafeBibleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BibleServiceServer will
// result in compilation errors.
type UnsafeBibleServiceServer interface {
	mustEmbedUnimplementedBibleServiceServer()
}

func RegisterBibleServiceServer(s grpc.ServiceRegistrar, srv BibleServiceServer) {
	s.RegisterService(&BibleService_ServiceDesc, srv)
}

func _BibleService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int32Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BibleServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.BibleService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BibleServiceServer).Get(ctx, req.(*Int32Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _BibleService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BibleServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.BibleService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BibleServiceServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BibleService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BibleServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.BibleService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BibleServiceServer).Create(ctx, req.(*StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _BibleService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int32Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BibleServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.BibleService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BibleServiceServer).Delete(ctx, req.(*Int32Value))
	}
	return interceptor(ctx, in, info, handler)
}

// BibleService_ServiceDesc is the grpc.ServiceDesc for BibleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BibleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.BibleService",
	HandlerType: (*BibleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _BibleService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _BibleService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _BibleService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BibleService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/protos/bible.proto",
}

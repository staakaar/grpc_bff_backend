// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: api/item.proto

package __proto

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

// ItemServiceClient is the client API for ItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ItemServiceClient interface {
	// サービスが持つメソッドの定義
	GetItem(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*ItemResponse, error)
	ItemServerStream(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (ItemService_ItemServerStreamClient, error)
	ItemClientStream(ctx context.Context, opts ...grpc.CallOption) (ItemService_ItemClientStreamClient, error)
}

type itemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewItemServiceClient(cc grpc.ClientConnInterface) ItemServiceClient {
	return &itemServiceClient{cc}
}

func (c *itemServiceClient) GetItem(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*ItemResponse, error) {
	out := new(ItemResponse)
	err := c.cc.Invoke(ctx, "/item.ItemService/GetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) ItemServerStream(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (ItemService_ItemServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ItemService_ServiceDesc.Streams[0], "/item.ItemService/ItemServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &itemServiceItemServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ItemService_ItemServerStreamClient interface {
	Recv() (*ItemResponse, error)
	grpc.ClientStream
}

type itemServiceItemServerStreamClient struct {
	grpc.ClientStream
}

func (x *itemServiceItemServerStreamClient) Recv() (*ItemResponse, error) {
	m := new(ItemResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *itemServiceClient) ItemClientStream(ctx context.Context, opts ...grpc.CallOption) (ItemService_ItemClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ItemService_ServiceDesc.Streams[1], "/item.ItemService/ItemClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &itemServiceItemClientStreamClient{stream}
	return x, nil
}

type ItemService_ItemClientStreamClient interface {
	Send(*ItemRequest) error
	CloseAndRecv() (*ItemResponse, error)
	grpc.ClientStream
}

type itemServiceItemClientStreamClient struct {
	grpc.ClientStream
}

func (x *itemServiceItemClientStreamClient) Send(m *ItemRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *itemServiceItemClientStreamClient) CloseAndRecv() (*ItemResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ItemResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ItemServiceServer is the server API for ItemService service.
// All implementations must embed UnimplementedItemServiceServer
// for forward compatibility
type ItemServiceServer interface {
	// サービスが持つメソッドの定義
	GetItem(context.Context, *ItemRequest) (*ItemResponse, error)
	ItemServerStream(*ItemRequest, ItemService_ItemServerStreamServer) error
	ItemClientStream(ItemService_ItemClientStreamServer) error
	mustEmbedUnimplementedItemServiceServer()
}

// UnimplementedItemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedItemServiceServer struct {
}

func (UnimplementedItemServiceServer) GetItem(context.Context, *ItemRequest) (*ItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}
func (UnimplementedItemServiceServer) ItemServerStream(*ItemRequest, ItemService_ItemServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ItemServerStream not implemented")
}
func (UnimplementedItemServiceServer) ItemClientStream(ItemService_ItemClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ItemClientStream not implemented")
}
func (UnimplementedItemServiceServer) mustEmbedUnimplementedItemServiceServer() {}

// UnsafeItemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ItemServiceServer will
// result in compilation errors.
type UnsafeItemServiceServer interface {
	mustEmbedUnimplementedItemServiceServer()
}

func RegisterItemServiceServer(s grpc.ServiceRegistrar, srv ItemServiceServer) {
	s.RegisterService(&ItemService_ServiceDesc, srv)
}

func _ItemService_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/GetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetItem(ctx, req.(*ItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_ItemServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ItemRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ItemServiceServer).ItemServerStream(m, &itemServiceItemServerStreamServer{stream})
}

type ItemService_ItemServerStreamServer interface {
	Send(*ItemResponse) error
	grpc.ServerStream
}

type itemServiceItemServerStreamServer struct {
	grpc.ServerStream
}

func (x *itemServiceItemServerStreamServer) Send(m *ItemResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ItemService_ItemClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ItemServiceServer).ItemClientStream(&itemServiceItemClientStreamServer{stream})
}

type ItemService_ItemClientStreamServer interface {
	SendAndClose(*ItemResponse) error
	Recv() (*ItemRequest, error)
	grpc.ServerStream
}

type itemServiceItemClientStreamServer struct {
	grpc.ServerStream
}

func (x *itemServiceItemClientStreamServer) SendAndClose(m *ItemResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *itemServiceItemClientStreamServer) Recv() (*ItemRequest, error) {
	m := new(ItemRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ItemService_ServiceDesc is the grpc.ServiceDesc for ItemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ItemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "item.ItemService",
	HandlerType: (*ItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetItem",
			Handler:    _ItemService_GetItem_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ItemServerStream",
			Handler:       _ItemService_ItemServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ItemClientStream",
			Handler:       _ItemService_ItemClientStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api/item.proto",
}

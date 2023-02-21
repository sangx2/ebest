// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// NewsClient is the client API for News service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewsClient interface {
	News(ctx context.Context, in *NewsRequest, opts ...grpc.CallOption) (News_NewsClient, error)
}

type newsClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsClient(cc grpc.ClientConnInterface) NewsClient {
	return &newsClient{cc}
}

func (c *newsClient) News(ctx context.Context, in *NewsRequest, opts ...grpc.CallOption) (News_NewsClient, error) {
	stream, err := c.cc.NewStream(ctx, &News_ServiceDesc.Streams[0], "/proto.News/News", opts...)
	if err != nil {
		return nil, err
	}
	x := &newsNewsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type News_NewsClient interface {
	Recv() (*NewsResponse, error)
	grpc.ClientStream
}

type newsNewsClient struct {
	grpc.ClientStream
}

func (x *newsNewsClient) Recv() (*NewsResponse, error) {
	m := new(NewsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NewsServer is the server API for News service.
// All implementations must embed UnimplementedNewsServer
// for forward compatibility
type NewsServer interface {
	News(*NewsRequest, News_NewsServer) error
	mustEmbedUnimplementedNewsServer()
}

// UnimplementedNewsServer must be embedded to have forward compatible implementations.
type UnimplementedNewsServer struct {
}

func (UnimplementedNewsServer) News(*NewsRequest, News_NewsServer) error {
	return status.Errorf(codes.Unimplemented, "method News not implemented")
}
func (UnimplementedNewsServer) mustEmbedUnimplementedNewsServer() {}

// UnsafeNewsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewsServer will
// result in compilation errors.
type UnsafeNewsServer interface {
	mustEmbedUnimplementedNewsServer()
}

func RegisterNewsServer(s grpc.ServiceRegistrar, srv NewsServer) {
	s.RegisterService(&News_ServiceDesc, srv)
}

func _News_News_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NewsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NewsServer).News(m, &newsNewsServer{stream})
}

type News_NewsServer interface {
	Send(*NewsResponse) error
	grpc.ServerStream
}

type newsNewsServer struct {
	grpc.ServerStream
}

func (x *newsNewsServer) Send(m *NewsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// News_ServiceDesc is the grpc.ServiceDesc for News service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var News_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.News",
	HandlerType: (*NewsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "News",
			Handler:       _News_News_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/news.proto",
}

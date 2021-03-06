// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/go-queue-server.proto

package api

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

// DataFetcherClient is the client API for DataFetcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataFetcherClient interface {
	FetchQueueStream(ctx context.Context, in *StreamSize, opts ...grpc.CallOption) (DataFetcher_FetchQueueStreamClient, error)
}

type dataFetcherClient struct {
	cc grpc.ClientConnInterface
}

func NewDataFetcherClient(cc grpc.ClientConnInterface) DataFetcherClient {
	return &dataFetcherClient{cc}
}

func (c *dataFetcherClient) FetchQueueStream(ctx context.Context, in *StreamSize, opts ...grpc.CallOption) (DataFetcher_FetchQueueStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataFetcher_ServiceDesc.Streams[0], "/api.DataFetcher/FetchQueueStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataFetcherFetchQueueStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataFetcher_FetchQueueStreamClient interface {
	Recv() (*Records, error)
	grpc.ClientStream
}

type dataFetcherFetchQueueStreamClient struct {
	grpc.ClientStream
}

func (x *dataFetcherFetchQueueStreamClient) Recv() (*Records, error) {
	m := new(Records)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataFetcherServer is the server API for DataFetcher service.
// All implementations must embed UnimplementedDataFetcherServer
// for forward compatibility
type DataFetcherServer interface {
	FetchQueueStream(*StreamSize, DataFetcher_FetchQueueStreamServer) error
	mustEmbedUnimplementedDataFetcherServer()
}

// UnimplementedDataFetcherServer must be embedded to have forward compatible implementations.
type UnimplementedDataFetcherServer struct {
}

func (UnimplementedDataFetcherServer) FetchQueueStream(*StreamSize, DataFetcher_FetchQueueStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchQueueStream not implemented")
}
func (UnimplementedDataFetcherServer) mustEmbedUnimplementedDataFetcherServer() {}

// UnsafeDataFetcherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataFetcherServer will
// result in compilation errors.
type UnsafeDataFetcherServer interface {
	mustEmbedUnimplementedDataFetcherServer()
}

func RegisterDataFetcherServer(s grpc.ServiceRegistrar, srv DataFetcherServer) {
	s.RegisterService(&DataFetcher_ServiceDesc, srv)
}

func _DataFetcher_FetchQueueStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamSize)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataFetcherServer).FetchQueueStream(m, &dataFetcherFetchQueueStreamServer{stream})
}

type DataFetcher_FetchQueueStreamServer interface {
	Send(*Records) error
	grpc.ServerStream
}

type dataFetcherFetchQueueStreamServer struct {
	grpc.ServerStream
}

func (x *dataFetcherFetchQueueStreamServer) Send(m *Records) error {
	return x.ServerStream.SendMsg(m)
}

// DataFetcher_ServiceDesc is the grpc.ServiceDesc for DataFetcher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataFetcher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.DataFetcher",
	HandlerType: (*DataFetcherServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchQueueStream",
			Handler:       _DataFetcher_FetchQueueStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/go-queue-server.proto",
}

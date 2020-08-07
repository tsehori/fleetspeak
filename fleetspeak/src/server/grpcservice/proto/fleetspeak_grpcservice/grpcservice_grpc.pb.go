// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package fleetspeak_grpcservice

import (
	context "context"

	fleetspeak "github.com/google/fleetspeak/fleetspeak/src/common/proto/fleetspeak"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProcessorClient is the client API for Processor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProcessorClient interface {
	// Process accepts message and processes it.
	Process(ctx context.Context, in *fleetspeak.Message, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error)
}

type processorClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessorClient(cc grpc.ClientConnInterface) ProcessorClient {
	return &processorClient{cc}
}

func (c *processorClient) Process(ctx context.Context, in *fleetspeak.Message, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error) {
	out := new(fleetspeak.EmptyMessage)
	err := c.cc.Invoke(ctx, "/fleetspeak.grpcservice.Processor/Process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessorServer is the server API for Processor service.
// All implementations should embed UnimplementedProcessorServer
// for forward compatibility
type ProcessorServer interface {
	// Process accepts message and processes it.
	Process(context.Context, *fleetspeak.Message) (*fleetspeak.EmptyMessage, error)
}

// UnimplementedProcessorServer should be embedded to have forward compatible implementations.
type UnimplementedProcessorServer struct {
}

func (*UnimplementedProcessorServer) Process(context.Context, *fleetspeak.Message) (*fleetspeak.EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Process not implemented")
}

func RegisterProcessorServer(s *grpc.Server, srv ProcessorServer) {
	s.RegisterService(&_Processor_serviceDesc, srv)
}

func _Processor_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fleetspeak.Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessorServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.grpcservice.Processor/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessorServer).Process(ctx, req.(*fleetspeak.Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Processor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fleetspeak.grpcservice.Processor",
	HandlerType: (*ProcessorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _Processor_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fleetspeak/src/server/grpcservice/proto/fleetspeak_grpcservice/grpcservice.proto",
}

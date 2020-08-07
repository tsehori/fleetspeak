// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package fleetspeak_server

import (
	context "context"
	fleetspeak "fleetspeak/src/common/proto/fleetspeak"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	// CreateBroadcast creates a FS broadcast, potentially sending a message to
	// many machines in the fleet.
	CreateBroadcast(ctx context.Context, in *CreateBroadcastRequest, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error)
	// ListActiveBroadcasts lists the currently active FS broadcasts.
	ListActiveBroadcasts(ctx context.Context, in *ListActiveBroadcastsRequest, opts ...grpc.CallOption) (*ListActiveBroadcastsResponse, error)
	// ListClients lists the currently active FS clients.
	ListClients(ctx context.Context, in *ListClientsRequest, opts ...grpc.CallOption) (*ListClientsResponse, error)
	// ListClientContacts lists the contact history for a client.
	ListClientContacts(ctx context.Context, in *ListClientContactsRequest, opts ...grpc.CallOption) (*ListClientContactsResponse, error)
	// GetMessageStatus retrieves the current status of a message.
	GetMessageStatus(ctx context.Context, in *GetMessageStatusRequest, opts ...grpc.CallOption) (*GetMessageStatusResponse, error)
	// InsertMessage inserts a message into the Fleetspeak system to be processed
	// by the server or delivered to a client.
	// TODO: Have this method return the message that is written to the
	// datastore (or at least the id).
	InsertMessage(ctx context.Context, in *fleetspeak.Message, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error)
	// DeletePendingMessages clears message queues for given clients.
	DeletePendingMessages(ctx context.Context, in *DeletePendingMessagesRequest, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error)
	// StoreFile inserts a file into the Fleetspeak system.
	StoreFile(ctx context.Context, in *StoreFileRequest, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error)
	// KeepAlive does as little as possible.
	KeepAlive(ctx context.Context, in *fleetspeak.EmptyMessage, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error)
	// BlacklistClient marks a client_id as invalid, forcing all Fleetspeak
	// clients using it to rekey.
	BlacklistClient(ctx context.Context, in *BlacklistClientRequest, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error)
	GetMetricValues(ctx context.Context, in *GetMetricValuesRequest, opts ...grpc.CallOption) (*GetMetricValuesResponse, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) CreateBroadcast(ctx context.Context, in *CreateBroadcastRequest, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error) {
	out := new(fleetspeak.EmptyMessage)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/CreateBroadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) ListActiveBroadcasts(ctx context.Context, in *ListActiveBroadcastsRequest, opts ...grpc.CallOption) (*ListActiveBroadcastsResponse, error) {
	out := new(ListActiveBroadcastsResponse)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/ListActiveBroadcasts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) ListClients(ctx context.Context, in *ListClientsRequest, opts ...grpc.CallOption) (*ListClientsResponse, error) {
	out := new(ListClientsResponse)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/ListClients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) ListClientContacts(ctx context.Context, in *ListClientContactsRequest, opts ...grpc.CallOption) (*ListClientContactsResponse, error) {
	out := new(ListClientContactsResponse)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/ListClientContacts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetMessageStatus(ctx context.Context, in *GetMessageStatusRequest, opts ...grpc.CallOption) (*GetMessageStatusResponse, error) {
	out := new(GetMessageStatusResponse)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/GetMessageStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) InsertMessage(ctx context.Context, in *fleetspeak.Message, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error) {
	out := new(fleetspeak.EmptyMessage)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/InsertMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) DeletePendingMessages(ctx context.Context, in *DeletePendingMessagesRequest, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error) {
	out := new(fleetspeak.EmptyMessage)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/DeletePendingMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) StoreFile(ctx context.Context, in *StoreFileRequest, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error) {
	out := new(fleetspeak.EmptyMessage)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/StoreFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) KeepAlive(ctx context.Context, in *fleetspeak.EmptyMessage, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error) {
	out := new(fleetspeak.EmptyMessage)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/KeepAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) BlacklistClient(ctx context.Context, in *BlacklistClientRequest, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error) {
	out := new(fleetspeak.EmptyMessage)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/BlacklistClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetMetricValues(ctx context.Context, in *GetMetricValuesRequest, opts ...grpc.CallOption) (*GetMetricValuesResponse, error) {
	out := new(GetMetricValuesResponse)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/GetMetricValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	// CreateBroadcast creates a FS broadcast, potentially sending a message to
	// many machines in the fleet.
	CreateBroadcast(context.Context, *CreateBroadcastRequest) (*fleetspeak.EmptyMessage, error)
	// ListActiveBroadcasts lists the currently active FS broadcasts.
	ListActiveBroadcasts(context.Context, *ListActiveBroadcastsRequest) (*ListActiveBroadcastsResponse, error)
	// ListClients lists the currently active FS clients.
	ListClients(context.Context, *ListClientsRequest) (*ListClientsResponse, error)
	// ListClientContacts lists the contact history for a client.
	ListClientContacts(context.Context, *ListClientContactsRequest) (*ListClientContactsResponse, error)
	// GetMessageStatus retrieves the current status of a message.
	GetMessageStatus(context.Context, *GetMessageStatusRequest) (*GetMessageStatusResponse, error)
	// InsertMessage inserts a message into the Fleetspeak system to be processed
	// by the server or delivered to a client.
	// TODO: Have this method return the message that is written to the
	// datastore (or at least the id).
	InsertMessage(context.Context, *fleetspeak.Message) (*fleetspeak.EmptyMessage, error)
	// DeletePendingMessages clears message queues for given clients.
	DeletePendingMessages(context.Context, *DeletePendingMessagesRequest) (*fleetspeak.EmptyMessage, error)
	// StoreFile inserts a file into the Fleetspeak system.
	StoreFile(context.Context, *StoreFileRequest) (*fleetspeak.EmptyMessage, error)
	// KeepAlive does as little as possible.
	KeepAlive(context.Context, *fleetspeak.EmptyMessage) (*fleetspeak.EmptyMessage, error)
	// BlacklistClient marks a client_id as invalid, forcing all Fleetspeak
	// clients using it to rekey.
	BlacklistClient(context.Context, *BlacklistClientRequest) (*fleetspeak.EmptyMessage, error)
	GetMetricValues(context.Context, *GetMetricValuesRequest) (*GetMetricValuesResponse, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (*UnimplementedAdminServer) CreateBroadcast(context.Context, *CreateBroadcastRequest) (*fleetspeak.EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBroadcast not implemented")
}
func (*UnimplementedAdminServer) ListActiveBroadcasts(context.Context, *ListActiveBroadcastsRequest) (*ListActiveBroadcastsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListActiveBroadcasts not implemented")
}
func (*UnimplementedAdminServer) ListClients(context.Context, *ListClientsRequest) (*ListClientsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClients not implemented")
}
func (*UnimplementedAdminServer) ListClientContacts(context.Context, *ListClientContactsRequest) (*ListClientContactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClientContacts not implemented")
}
func (*UnimplementedAdminServer) GetMessageStatus(context.Context, *GetMessageStatusRequest) (*GetMessageStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageStatus not implemented")
}
func (*UnimplementedAdminServer) InsertMessage(context.Context, *fleetspeak.Message) (*fleetspeak.EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertMessage not implemented")
}
func (*UnimplementedAdminServer) DeletePendingMessages(context.Context, *DeletePendingMessagesRequest) (*fleetspeak.EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePendingMessages not implemented")
}
func (*UnimplementedAdminServer) StoreFile(context.Context, *StoreFileRequest) (*fleetspeak.EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreFile not implemented")
}
func (*UnimplementedAdminServer) KeepAlive(context.Context, *fleetspeak.EmptyMessage) (*fleetspeak.EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeepAlive not implemented")
}
func (*UnimplementedAdminServer) BlacklistClient(context.Context, *BlacklistClientRequest) (*fleetspeak.EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlacklistClient not implemented")
}
func (*UnimplementedAdminServer) GetMetricValues(context.Context, *GetMetricValuesRequest) (*GetMetricValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricValues not implemented")
}
func (*UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_CreateBroadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBroadcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).CreateBroadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/CreateBroadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).CreateBroadcast(ctx, req.(*CreateBroadcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_ListActiveBroadcasts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListActiveBroadcastsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ListActiveBroadcasts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/ListActiveBroadcasts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ListActiveBroadcasts(ctx, req.(*ListActiveBroadcastsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_ListClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ListClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/ListClients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ListClients(ctx, req.(*ListClientsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_ListClientContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientContactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ListClientContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/ListClientContacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ListClientContacts(ctx, req.(*ListClientContactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetMessageStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetMessageStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/GetMessageStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetMessageStatus(ctx, req.(*GetMessageStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_InsertMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fleetspeak.Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).InsertMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/InsertMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).InsertMessage(ctx, req.(*fleetspeak.Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_DeletePendingMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePendingMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).DeletePendingMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/DeletePendingMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).DeletePendingMessages(ctx, req.(*DeletePendingMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_StoreFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).StoreFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/StoreFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).StoreFile(ctx, req.(*StoreFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_KeepAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fleetspeak.EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).KeepAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/KeepAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).KeepAlive(ctx, req.(*fleetspeak.EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_BlacklistClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlacklistClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).BlacklistClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/BlacklistClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).BlacklistClient(ctx, req.(*BlacklistClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetMetricValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetMetricValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/GetMetricValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetMetricValues(ctx, req.(*GetMetricValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fleetspeak.server.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBroadcast",
			Handler:    _Admin_CreateBroadcast_Handler,
		},
		{
			MethodName: "ListActiveBroadcasts",
			Handler:    _Admin_ListActiveBroadcasts_Handler,
		},
		{
			MethodName: "ListClients",
			Handler:    _Admin_ListClients_Handler,
		},
		{
			MethodName: "ListClientContacts",
			Handler:    _Admin_ListClientContacts_Handler,
		},
		{
			MethodName: "GetMessageStatus",
			Handler:    _Admin_GetMessageStatus_Handler,
		},
		{
			MethodName: "InsertMessage",
			Handler:    _Admin_InsertMessage_Handler,
		},
		{
			MethodName: "DeletePendingMessages",
			Handler:    _Admin_DeletePendingMessages_Handler,
		},
		{
			MethodName: "StoreFile",
			Handler:    _Admin_StoreFile_Handler,
		},
		{
			MethodName: "KeepAlive",
			Handler:    _Admin_KeepAlive_Handler,
		},
		{
			MethodName: "BlacklistClient",
			Handler:    _Admin_BlacklistClient_Handler,
		},
		{
			MethodName: "GetMetricValues",
			Handler:    _Admin_GetMetricValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fleetspeak/src/server/proto/fleetspeak_server/admin.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: fleetspeak/src/inttesting/frr/proto/fleetspeak_frr/frr.proto

package fleetspeak_frr

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	fleetspeak "github.com/google/fleetspeak/fleetspeak/src/common/proto/fleetspeak"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Contains the information needed to configure a frr server component.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address to reach the master frr server over grpc.
	MasterServer string `protobuf:"bytes,1,opt,name=master_server,json=masterServer,proto3" json:"master_server,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetMasterServer() string {
	if x != nil {
		return x.MasterServer
	}
	return ""
}

// A TrafficRequest message is sent from the server to the client which tells the
// client to send random data back.
type TrafficRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An identifier used to identify the frr master instance responsible for this.
	MasterId int64 `protobuf:"varint,1,opt,name=master_id,json=masterId,proto3" json:"master_id,omitempty"`
	// An identifier used to link responses to requests.
	RequestId int64 `protobuf:"varint,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// How many messages to send back for this request. Defaults to 1.
	NumMessages int64 `protobuf:"varint,3,opt,name=num_messages,json=numMessages,proto3" json:"num_messages,omitempty"`
	// How large should each message be, in bytes. Defaults to 1024.
	MessageSize int64 `protobuf:"varint,4,opt,name=message_size,json=messageSize,proto3" json:"message_size,omitempty"`
	// How long to wait between messages. Defaults to 0.
	MessageDelayMs int64 `protobuf:"varint,5,opt,name=message_delay_ms,json=messageDelayMs,proto3" json:"message_delay_ms,omitempty"`
	// How much to jitter the previous parameters - all parameters will be
	// multiplied by a number between 1.0 and 1.0 + jitter.
	Jitter float32 `protobuf:"fixed32,6,opt,name=jitter,proto3" json:"jitter,omitempty"`
}

func (x *TrafficRequestData) Reset() {
	*x = TrafficRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrafficRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrafficRequestData) ProtoMessage() {}

func (x *TrafficRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrafficRequestData.ProtoReflect.Descriptor instead.
func (*TrafficRequestData) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDescGZIP(), []int{1}
}

func (x *TrafficRequestData) GetMasterId() int64 {
	if x != nil {
		return x.MasterId
	}
	return 0
}

func (x *TrafficRequestData) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *TrafficRequestData) GetNumMessages() int64 {
	if x != nil {
		return x.NumMessages
	}
	return 0
}

func (x *TrafficRequestData) GetMessageSize() int64 {
	if x != nil {
		return x.MessageSize
	}
	return 0
}

func (x *TrafficRequestData) GetMessageDelayMs() int64 {
	if x != nil {
		return x.MessageDelayMs
	}
	return 0
}

func (x *TrafficRequestData) GetJitter() float32 {
	if x != nil {
		return x.Jitter
	}
	return 0
}

type TrafficResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MasterId      int64  `protobuf:"varint,1,opt,name=master_id,json=masterId,proto3" json:"master_id,omitempty"`
	RequestId     int64  `protobuf:"varint,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	ResponseIndex int64  `protobuf:"varint,3,opt,name=response_index,json=responseIndex,proto3" json:"response_index,omitempty"`
	Data          []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	// Set when this is the last message responsive to the given request.
	Fin bool `protobuf:"varint,5,opt,name=fin,proto3" json:"fin,omitempty"`
}

func (x *TrafficResponseData) Reset() {
	*x = TrafficResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrafficResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrafficResponseData) ProtoMessage() {}

func (x *TrafficResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrafficResponseData.ProtoReflect.Descriptor instead.
func (*TrafficResponseData) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDescGZIP(), []int{2}
}

func (x *TrafficResponseData) GetMasterId() int64 {
	if x != nil {
		return x.MasterId
	}
	return 0
}

func (x *TrafficResponseData) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *TrafficResponseData) GetResponseIndex() int64 {
	if x != nil {
		return x.ResponseIndex
	}
	return 0
}

func (x *TrafficResponseData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TrafficResponseData) GetFin() bool {
	if x != nil {
		return x.Fin
	}
	return false
}

// A FileRequest is sent from the server to the client and tells
// the client to attempt to download a file from the server.
type FileRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An identifier used to identify the frr master instance
	// responsible for this.
	MasterId int64 `protobuf:"varint,1,opt,name=master_id,json=masterId,proto3" json:"master_id,omitempty"`
	// The name of the file to download.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FileRequestData) Reset() {
	*x = FileRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequestData) ProtoMessage() {}

func (x *FileRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequestData.ProtoReflect.Descriptor instead.
func (*FileRequestData) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDescGZIP(), []int{3}
}

func (x *FileRequestData) GetMasterId() int64 {
	if x != nil {
		return x.MasterId
	}
	return 0
}

func (x *FileRequestData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// A FileResponse is sent from the client to the server and
// reports that the client successfully downloaded a file from the
// server.
type FileResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An identifier used to identify the frr master instance
	// responsible for the underlying request.
	MasterId int64 `protobuf:"varint,1,opt,name=master_id,json=masterId,proto3" json:"master_id,omitempty"`
	// The name of the file that was downloaded.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The size of the file that was downloaded.
	Size uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *FileResponseData) Reset() {
	*x = FileResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResponseData) ProtoMessage() {}

func (x *FileResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResponseData.ProtoReflect.Descriptor instead.
func (*FileResponseData) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDescGZIP(), []int{4}
}

func (x *FileResponseData) GetMasterId() int64 {
	if x != nil {
		return x.MasterId
	}
	return 0
}

func (x *FileResponseData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileResponseData) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type MessageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId []byte               `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Data     *TrafficResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MessageInfo) Reset() {
	*x = MessageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageInfo) ProtoMessage() {}

func (x *MessageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageInfo.ProtoReflect.Descriptor instead.
func (*MessageInfo) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDescGZIP(), []int{5}
}

func (x *MessageInfo) GetClientId() []byte {
	if x != nil {
		return x.ClientId
	}
	return nil
}

func (x *MessageInfo) GetData() *TrafficResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type FileResponseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId []byte            `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Data     *FileResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FileResponseInfo) Reset() {
	*x = FileResponseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileResponseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResponseInfo) ProtoMessage() {}

func (x *FileResponseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResponseInfo.ProtoReflect.Descriptor instead.
func (*FileResponseInfo) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDescGZIP(), []int{6}
}

func (x *FileResponseInfo) GetClientId() []byte {
	if x != nil {
		return x.ClientId
	}
	return nil
}

func (x *FileResponseInfo) GetData() *FileResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type CompletedRequestsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *CompletedRequestsRequest) Reset() {
	*x = CompletedRequestsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletedRequestsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletedRequestsRequest) ProtoMessage() {}

func (x *CompletedRequestsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletedRequestsRequest.ProtoReflect.Descriptor instead.
func (*CompletedRequestsRequest) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDescGZIP(), []int{7}
}

func (x *CompletedRequestsRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type CompletedRequestsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestIds []int64 `protobuf:"varint,1,rep,packed,name=request_ids,json=requestIds,proto3" json:"request_ids,omitempty"`
}

func (x *CompletedRequestsResponse) Reset() {
	*x = CompletedRequestsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletedRequestsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletedRequestsResponse) ProtoMessage() {}

func (x *CompletedRequestsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletedRequestsResponse.ProtoReflect.Descriptor instead.
func (*CompletedRequestsResponse) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDescGZIP(), []int{8}
}

func (x *CompletedRequestsResponse) GetRequestIds() []int64 {
	if x != nil {
		return x.RequestIds
	}
	return nil
}

type CreateHuntRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *TrafficRequestData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Limit uint64              `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *CreateHuntRequest) Reset() {
	*x = CreateHuntRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHuntRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHuntRequest) ProtoMessage() {}

func (x *CreateHuntRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHuntRequest.ProtoReflect.Descriptor instead.
func (*CreateHuntRequest) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDescGZIP(), []int{9}
}

func (x *CreateHuntRequest) GetData() *TrafficRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CreateHuntRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type CreateHuntResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateHuntResponse) Reset() {
	*x = CreateHuntResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHuntResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHuntResponse) ProtoMessage() {}

func (x *CreateHuntResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHuntResponse.ProtoReflect.Descriptor instead.
func (*CreateHuntResponse) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDescGZIP(), []int{10}
}

var File_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto protoreflect.FileDescriptor

var file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x69, 0x6e, 0x74, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x66, 0x72, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b,
	0x5f, 0x66, 0x72, 0x72, 0x2f, 0x66, 0x72, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2e, 0x66, 0x72, 0x72, 0x1a, 0x33,
	0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x65, 0x65,
	0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23, 0x0a,
	0x0d, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x22, 0xd8, 0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6e, 0x75, 0x6d,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x6d, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65,
	0x6c, 0x61, 0x79, 0x4d, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6a, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6a, 0x69, 0x74, 0x74, 0x65, 0x72, 0x22, 0x9e, 0x01,
	0x0a, 0x13, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03,
	0x66, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x69, 0x6e, 0x22, 0x42,
	0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x57, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x63, 0x0a, 0x0b, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65,
	0x61, 0x6b, 0x2e, 0x66, 0x72, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x65, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2e, 0x66, 0x72, 0x72,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x37, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x3c, 0x0a, 0x19, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x73, 0x22, 0x61,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2e, 0x66,
	0x72, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf1, 0x02, 0x0a, 0x06, 0x4d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x50, 0x0a, 0x15, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x2e, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2e, 0x66, 0x72, 0x72, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x18, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74,
	0x73, 0x70, 0x65, 0x61, 0x6b, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x2e, 0x66, 0x6c, 0x65,
	0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2e, 0x66, 0x72, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x18, 0x2e, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x28, 0x2e,
	0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2e, 0x66, 0x72, 0x72, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73,
	0x70, 0x65, 0x61, 0x6b, 0x2e, 0x66, 0x72, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x75,
	0x6e, 0x74, 0x12, 0x21, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2e,
	0x66, 0x72, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65,
	0x61, 0x6b, 0x2e, 0x66, 0x72, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDescOnce sync.Once
	file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDescData = file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDesc
)

func file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDescGZIP() []byte {
	file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDescOnce.Do(func() {
		file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDescData = protoimpl.X.CompressGZIP(file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDescData)
	})
	return file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDescData
}

var file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_goTypes = []interface{}{
	(*Config)(nil),                    // 0: fleetspeak.frr.Config
	(*TrafficRequestData)(nil),        // 1: fleetspeak.frr.TrafficRequestData
	(*TrafficResponseData)(nil),       // 2: fleetspeak.frr.TrafficResponseData
	(*FileRequestData)(nil),           // 3: fleetspeak.frr.FileRequestData
	(*FileResponseData)(nil),          // 4: fleetspeak.frr.FileResponseData
	(*MessageInfo)(nil),               // 5: fleetspeak.frr.MessageInfo
	(*FileResponseInfo)(nil),          // 6: fleetspeak.frr.FileResponseInfo
	(*CompletedRequestsRequest)(nil),  // 7: fleetspeak.frr.CompletedRequestsRequest
	(*CompletedRequestsResponse)(nil), // 8: fleetspeak.frr.CompletedRequestsResponse
	(*CreateHuntRequest)(nil),         // 9: fleetspeak.frr.CreateHuntRequest
	(*CreateHuntResponse)(nil),        // 10: fleetspeak.frr.CreateHuntResponse
	(*fleetspeak.EmptyMessage)(nil),   // 11: fleetspeak.EmptyMessage
}
var file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_depIdxs = []int32{
	2,  // 0: fleetspeak.frr.MessageInfo.data:type_name -> fleetspeak.frr.TrafficResponseData
	4,  // 1: fleetspeak.frr.FileResponseInfo.data:type_name -> fleetspeak.frr.FileResponseData
	1,  // 2: fleetspeak.frr.CreateHuntRequest.data:type_name -> fleetspeak.frr.TrafficRequestData
	5,  // 3: fleetspeak.frr.Master.RecordTrafficResponse:input_type -> fleetspeak.frr.MessageInfo
	6,  // 4: fleetspeak.frr.Master.RecordFileResponse:input_type -> fleetspeak.frr.FileResponseInfo
	7,  // 5: fleetspeak.frr.Master.CompletedRequests:input_type -> fleetspeak.frr.CompletedRequestsRequest
	9,  // 6: fleetspeak.frr.Master.CreateHunt:input_type -> fleetspeak.frr.CreateHuntRequest
	11, // 7: fleetspeak.frr.Master.RecordTrafficResponse:output_type -> fleetspeak.EmptyMessage
	11, // 8: fleetspeak.frr.Master.RecordFileResponse:output_type -> fleetspeak.EmptyMessage
	8,  // 9: fleetspeak.frr.Master.CompletedRequests:output_type -> fleetspeak.frr.CompletedRequestsResponse
	10, // 10: fleetspeak.frr.Master.CreateHunt:output_type -> fleetspeak.frr.CreateHuntResponse
	7,  // [7:11] is the sub-list for method output_type
	3,  // [3:7] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_init() }
func file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_init() {
	if File_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrafficRequestData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrafficResponseData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRequestData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileResponseData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileResponseInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompletedRequestsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompletedRequestsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHuntRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHuntResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_goTypes,
		DependencyIndexes: file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_depIdxs,
		MessageInfos:      file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_msgTypes,
	}.Build()
	File_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto = out.File
	file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_rawDesc = nil
	file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_goTypes = nil
	file_fleetspeak_src_inttesting_frr_proto_fleetspeak_frr_frr_proto_depIdxs = nil
}

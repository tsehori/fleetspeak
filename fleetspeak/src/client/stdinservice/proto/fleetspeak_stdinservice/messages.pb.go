// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: fleetspeak/src/client/stdinservice/proto/fleetspeak_stdinservice/messages.proto

package fleetspeak_stdinservice

import (
	fleetspeak_monitoring "fleetspeak/src/common/proto/fleetspeak_monitoring"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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

type InputMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The data to be forwarded to the service.
	Input []byte `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	// Command line arguments.
	Args []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *InputMessage) Reset() {
	*x = InputMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputMessage) ProtoMessage() {}

func (x *InputMessage) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputMessage.ProtoReflect.Descriptor instead.
func (*InputMessage) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_rawDescGZIP(), []int{0}
}

func (x *InputMessage) GetInput() []byte {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *InputMessage) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

type OutputMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stdout        []byte                                         `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr        []byte                                         `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	ResourceUsage *fleetspeak_monitoring.AggregatedResourceUsage `protobuf:"bytes,3,opt,name=resource_usage,json=resourceUsage,proto3" json:"resource_usage,omitempty"`
	// When the message was generated.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *OutputMessage) Reset() {
	*x = OutputMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutputMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputMessage) ProtoMessage() {}

func (x *OutputMessage) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputMessage.ProtoReflect.Descriptor instead.
func (*OutputMessage) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_rawDescGZIP(), []int{1}
}

func (x *OutputMessage) GetStdout() []byte {
	if x != nil {
		return x.Stdout
	}
	return nil
}

func (x *OutputMessage) GetStderr() []byte {
	if x != nil {
		return x.Stderr
	}
	return nil
}

func (x *OutputMessage) GetResourceUsage() *fleetspeak_monitoring.AggregatedResourceUsage {
	if x != nil {
		return x.ResourceUsage
	}
	return nil
}

func (x *OutputMessage) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto protoreflect.FileDescriptor

var file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x64, 0x69, 0x6e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74,
	0x73, 0x70, 0x65, 0x61, 0x6b, 0x5f, 0x73, 0x74, 0x64, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x17, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2e, 0x73, 0x74,
	0x64, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x66, 0x6c, 0x65,
	0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70,
	0x65, 0x61, 0x6b, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a,
	0x0c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0xd0, 0x01, 0x0a, 0x0d, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64,
	0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x12, 0x55, 0x0a, 0x0e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_rawDescOnce sync.Once
	file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_rawDescData = file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_rawDesc
)

func file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_rawDescGZIP() []byte {
	file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_rawDescOnce.Do(func() {
		file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_rawDescData)
	})
	return file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_rawDescData
}

var file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_goTypes = []interface{}{
	(*InputMessage)(nil),  // 0: fleetspeak.stdinservice.InputMessage
	(*OutputMessage)(nil), // 1: fleetspeak.stdinservice.OutputMessage
	(*fleetspeak_monitoring.AggregatedResourceUsage)(nil), // 2: fleetspeak.monitoring.AggregatedResourceUsage
	(*timestamp.Timestamp)(nil),                           // 3: google.protobuf.Timestamp
}
var file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_depIdxs = []int32{
	2, // 0: fleetspeak.stdinservice.OutputMessage.resource_usage:type_name -> fleetspeak.monitoring.AggregatedResourceUsage
	3, // 1: fleetspeak.stdinservice.OutputMessage.timestamp:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_init()
}
func file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_init() {
	if File_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputMessage); i {
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
		file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutputMessage); i {
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
			RawDescriptor: file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_goTypes,
		DependencyIndexes: file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_depIdxs,
		MessageInfos:      file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_msgTypes,
	}.Build()
	File_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto = out.File
	file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_rawDesc = nil
	file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_goTypes = nil
	file_fleetspeak_src_client_stdinservice_proto_fleetspeak_stdinservice_messages_proto_depIdxs = nil
}

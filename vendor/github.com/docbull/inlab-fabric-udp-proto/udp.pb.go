// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protos/udp.proto

package inlab_udp

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type StatusCode int32

const (
	StatusCode_Unknown StatusCode = 0
	StatusCode_Ok      StatusCode = 1
	StatusCode_Failed  StatusCode = 2
)

// Enum value maps for StatusCode.
var (
	StatusCode_name = map[int32]string{
		0: "Unknown",
		1: "Ok",
		2: "Failed",
	}
	StatusCode_value = map[string]int32{
		"Unknown": 0,
		"Ok":      1,
		"Failed":  2,
	}
)

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}

func (x StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_udp_proto_enumTypes[0].Descriptor()
}

func (StatusCode) Type() protoreflect.EnumType {
	return &file_protos_udp_proto_enumTypes[0]
}

func (x StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode.Descriptor instead.
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_protos_udp_proto_rawDescGZIP(), []int{0}
}

// Envelope contains a marshalled
// GossipMessage and a signature over it.
// It may also contain a SecretEnvelope
// which is a marshalled Secret
type Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload        []byte          `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature      []byte          `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	SecretEnvelope *SecretEnvelope `protobuf:"bytes,3,opt,name=secret_envelope,json=secretEnvelope,proto3" json:"secret_envelope,omitempty"`
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_udp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_protos_udp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_protos_udp_proto_rawDescGZIP(), []int{0}
}

func (x *Envelope) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Envelope) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Envelope) GetSecretEnvelope() *SecretEnvelope {
	if x != nil {
		return x.SecretEnvelope
	}
	return nil
}

// SecretEnvelope is a marshalled Secret
// and a signature over it.
// The signature should be validated by the peer
// that signed the Envelope the SecretEnvelope
// came with
type SecretEnvelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload   []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SecretEnvelope) Reset() {
	*x = SecretEnvelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_udp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretEnvelope) ProtoMessage() {}

func (x *SecretEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_protos_udp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretEnvelope.ProtoReflect.Descriptor instead.
func (*SecretEnvelope) Descriptor() ([]byte, []int) {
	return file_protos_udp_proto_rawDescGZIP(), []int{1}
}

func (x *SecretEnvelope) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *SecretEnvelope) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code StatusCode `protobuf:"varint,1,opt,name=Code,proto3,enum=udp.StatusCode" json:"Code,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_udp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_protos_udp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_protos_udp_proto_rawDescGZIP(), []int{2}
}

func (x *Status) GetCode() StatusCode {
	if x != nil {
		return x.Code
	}
	return StatusCode_Unknown
}

var File_protos_udp_proto protoreflect.FileDescriptor

var file_protos_udp_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x64, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x75, 0x64, 0x70, 0x22, 0x80, 0x01, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x3c, 0x0a, 0x0f,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75, 0x64, 0x70, 0x2e, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x22, 0x48, 0x0a, 0x0e, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x22, 0x2d, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23,
	0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x75,
	0x64, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x2a, 0x2d, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x06,
	0x0a, 0x02, 0x4f, 0x6b, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x10, 0x02, 0x32, 0x3d, 0x0a, 0x0a, 0x55, 0x44, 0x50, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2f, 0x0a, 0x0f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72,
	0x55, 0x44, 0x50, 0x12, 0x0d, 0x2e, 0x75, 0x64, 0x70, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x1a, 0x0b, 0x2e, 0x75, 0x64, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x6f, 0x63, 0x62, 0x75, 0x6c, 0x6c, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2d, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x75, 0x64, 0x70, 0x2f, 0x69,
	0x6e, 0x6c, 0x61, 0x62, 0x2d, 0x75, 0x64, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_udp_proto_rawDescOnce sync.Once
	file_protos_udp_proto_rawDescData = file_protos_udp_proto_rawDesc
)

func file_protos_udp_proto_rawDescGZIP() []byte {
	file_protos_udp_proto_rawDescOnce.Do(func() {
		file_protos_udp_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_udp_proto_rawDescData)
	})
	return file_protos_udp_proto_rawDescData
}

var file_protos_udp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_udp_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_udp_proto_goTypes = []interface{}{
	(StatusCode)(0),        // 0: udp.StatusCode
	(*Envelope)(nil),       // 1: udp.Envelope
	(*SecretEnvelope)(nil), // 2: udp.SecretEnvelope
	(*Status)(nil),         // 3: udp.Status
}
var file_protos_udp_proto_depIdxs = []int32{
	2, // 0: udp.Envelope.secret_envelope:type_name -> udp.SecretEnvelope
	0, // 1: udp.Status.Code:type_name -> udp.StatusCode
	1, // 2: udp.UDPService.BlockDataForUDP:input_type -> udp.Envelope
	3, // 3: udp.UDPService.BlockDataForUDP:output_type -> udp.Status
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_udp_proto_init() }
func file_protos_udp_proto_init() {
	if File_protos_udp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_udp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Envelope); i {
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
		file_protos_udp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretEnvelope); i {
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
		file_protos_udp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_protos_udp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_udp_proto_goTypes,
		DependencyIndexes: file_protos_udp_proto_depIdxs,
		EnumInfos:         file_protos_udp_proto_enumTypes,
		MessageInfos:      file_protos_udp_proto_msgTypes,
	}.Build()
	File_protos_udp_proto = out.File
	file_protos_udp_proto_rawDesc = nil
	file_protos_udp_proto_goTypes = nil
	file_protos_udp_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UDPServiceClient is the client API for UDPService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UDPServiceClient interface {
	// BlockDataForUDP receives a block data from
	// the Peer container for block sending over UDP
	BlockDataForUDP(ctx context.Context, in *Envelope, opts ...grpc.CallOption) (*Status, error)
}

type uDPServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUDPServiceClient(cc grpc.ClientConnInterface) UDPServiceClient {
	return &uDPServiceClient{cc}
}

func (c *uDPServiceClient) BlockDataForUDP(ctx context.Context, in *Envelope, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/udp.UDPService/BlockDataForUDP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UDPServiceServer is the server API for UDPService service.
type UDPServiceServer interface {
	// BlockDataForUDP receives a block data from
	// the Peer container for block sending over UDP
	BlockDataForUDP(context.Context, *Envelope) (*Status, error)
}

// UnimplementedUDPServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUDPServiceServer struct {
}

func (*UnimplementedUDPServiceServer) BlockDataForUDP(context.Context, *Envelope) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockDataForUDP not implemented")
}

func RegisterUDPServiceServer(s *grpc.Server, srv UDPServiceServer) {
	s.RegisterService(&_UDPService_serviceDesc, srv)
}

func _UDPService_BlockDataForUDP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Envelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UDPServiceServer).BlockDataForUDP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/udp.UDPService/BlockDataForUDP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UDPServiceServer).BlockDataForUDP(ctx, req.(*Envelope))
	}
	return interceptor(ctx, in, info, handler)
}

var _UDPService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "udp.UDPService",
	HandlerType: (*UDPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BlockDataForUDP",
			Handler:    _UDPService_BlockDataForUDP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/udp.proto",
}

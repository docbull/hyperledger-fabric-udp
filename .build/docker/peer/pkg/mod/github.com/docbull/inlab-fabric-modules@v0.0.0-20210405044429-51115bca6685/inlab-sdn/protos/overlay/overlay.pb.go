// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: overlay/overlay.proto

package overlay

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

type PeerEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint string `protobuf:"bytes,1,opt,name=Endpoint,proto3" json:"Endpoint,omitempty"`
}

func (x *PeerEndpoint) Reset() {
	*x = PeerEndpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_overlay_overlay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerEndpoint) ProtoMessage() {}

func (x *PeerEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_overlay_overlay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerEndpoint.ProtoReflect.Descriptor instead.
func (*PeerEndpoint) Descriptor() ([]byte, []int) {
	return file_overlay_overlay_proto_rawDescGZIP(), []int{0}
}

func (x *PeerEndpoint) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type OverlayStructure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SuperPeer string   `protobuf:"bytes,1,opt,name=SuperPeer,proto3" json:"SuperPeer,omitempty"`
	Endpoint  string   `protobuf:"bytes,2,opt,name=Endpoint,proto3" json:"Endpoint,omitempty"`
	SubPeers  []string `protobuf:"bytes,3,rep,name=SubPeers,proto3" json:"SubPeers,omitempty"`
}

func (x *OverlayStructure) Reset() {
	*x = OverlayStructure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_overlay_overlay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverlayStructure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverlayStructure) ProtoMessage() {}

func (x *OverlayStructure) ProtoReflect() protoreflect.Message {
	mi := &file_overlay_overlay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverlayStructure.ProtoReflect.Descriptor instead.
func (*OverlayStructure) Descriptor() ([]byte, []int) {
	return file_overlay_overlay_proto_rawDescGZIP(), []int{1}
}

func (x *OverlayStructure) GetSuperPeer() string {
	if x != nil {
		return x.SuperPeer
	}
	return ""
}

func (x *OverlayStructure) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *OverlayStructure) GetSubPeers() []string {
	if x != nil {
		return x.SubPeers
	}
	return nil
}

var File_overlay_overlay_proto protoreflect.FileDescriptor

var file_overlay_overlay_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x2f, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79,
	0x22, 0x2a, 0x0a, 0x0c, 0x50, 0x65, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x68, 0x0a, 0x10,
	0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x53, 0x75, 0x70, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x75, 0x70, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x75,
	0x62, 0x50, 0x65, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x53, 0x75,
	0x62, 0x50, 0x65, 0x65, 0x72, 0x73, 0x32, 0x64, 0x0a, 0x17, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61,
	0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x49, 0x0a, 0x13, 0x53, 0x65, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x15, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x6c,
	0x61, 0x79, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x1a,
	0x19, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61,
	0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18,
	0x69, 0x6e, 0x6c, 0x61, 0x62, 0x2d, 0x73, 0x64, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_overlay_overlay_proto_rawDescOnce sync.Once
	file_overlay_overlay_proto_rawDescData = file_overlay_overlay_proto_rawDesc
)

func file_overlay_overlay_proto_rawDescGZIP() []byte {
	file_overlay_overlay_proto_rawDescOnce.Do(func() {
		file_overlay_overlay_proto_rawDescData = protoimpl.X.CompressGZIP(file_overlay_overlay_proto_rawDescData)
	})
	return file_overlay_overlay_proto_rawDescData
}

var file_overlay_overlay_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_overlay_overlay_proto_goTypes = []interface{}{
	(*PeerEndpoint)(nil),     // 0: overlay.PeerEndpoint
	(*OverlayStructure)(nil), // 1: overlay.OverlayStructure
}
var file_overlay_overlay_proto_depIdxs = []int32{
	0, // 0: overlay.OverlayStructureService.SeeOverlayStructure:input_type -> overlay.PeerEndpoint
	1, // 1: overlay.OverlayStructureService.SeeOverlayStructure:output_type -> overlay.OverlayStructure
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_overlay_overlay_proto_init() }
func file_overlay_overlay_proto_init() {
	if File_overlay_overlay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_overlay_overlay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerEndpoint); i {
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
		file_overlay_overlay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OverlayStructure); i {
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
			RawDescriptor: file_overlay_overlay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_overlay_overlay_proto_goTypes,
		DependencyIndexes: file_overlay_overlay_proto_depIdxs,
		MessageInfos:      file_overlay_overlay_proto_msgTypes,
	}.Build()
	File_overlay_overlay_proto = out.File
	file_overlay_overlay_proto_rawDesc = nil
	file_overlay_overlay_proto_goTypes = nil
	file_overlay_overlay_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OverlayStructureServiceClient is the client API for OverlayStructureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OverlayStructureServiceClient interface {
	SeeOverlayStructure(ctx context.Context, in *PeerEndpoint, opts ...grpc.CallOption) (*OverlayStructure, error)
}

type overlayStructureServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOverlayStructureServiceClient(cc grpc.ClientConnInterface) OverlayStructureServiceClient {
	return &overlayStructureServiceClient{cc}
}

func (c *overlayStructureServiceClient) SeeOverlayStructure(ctx context.Context, in *PeerEndpoint, opts ...grpc.CallOption) (*OverlayStructure, error) {
	out := new(OverlayStructure)
	err := c.cc.Invoke(ctx, "/overlay.OverlayStructureService/SeeOverlayStructure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OverlayStructureServiceServer is the server API for OverlayStructureService service.
type OverlayStructureServiceServer interface {
	SeeOverlayStructure(context.Context, *PeerEndpoint) (*OverlayStructure, error)
}

// UnimplementedOverlayStructureServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOverlayStructureServiceServer struct {
}

func (*UnimplementedOverlayStructureServiceServer) SeeOverlayStructure(context.Context, *PeerEndpoint) (*OverlayStructure, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SeeOverlayStructure not implemented")
}

func RegisterOverlayStructureServiceServer(s *grpc.Server, srv OverlayStructureServiceServer) {
	s.RegisterService(&_OverlayStructureService_serviceDesc, srv)
}

func _OverlayStructureService_SeeOverlayStructure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverlayStructureServiceServer).SeeOverlayStructure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.OverlayStructureService/SeeOverlayStructure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverlayStructureServiceServer).SeeOverlayStructure(ctx, req.(*PeerEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

var _OverlayStructureService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "overlay.OverlayStructureService",
	HandlerType: (*OverlayStructureServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SeeOverlayStructure",
			Handler:    _OverlayStructureService_SeeOverlayStructure_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "overlay/overlay.proto",
}

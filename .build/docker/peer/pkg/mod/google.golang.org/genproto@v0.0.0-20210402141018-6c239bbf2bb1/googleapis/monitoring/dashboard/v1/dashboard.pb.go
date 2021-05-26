// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: google/monitoring/dashboard/v1/dashboard.proto

package dashboard

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A Google Stackdriver dashboard. Dashboards define the content and layout
// of pages in the Stackdriver web application.
type Dashboard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the dashboard.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The mutable, human-readable name.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// `etag` is used for optimistic concurrency control as a way to help
	// prevent simultaneous updates of a policy from overwriting each other.
	// An `etag` is returned in the response to `GetDashboard`, and
	// users are expected to put that etag in the request to `UpdateDashboard` to
	// ensure that their change will be applied to the same version of the
	// Dashboard configuration. The field should not be passed during
	// dashboard creation.
	Etag string `protobuf:"bytes,4,opt,name=etag,proto3" json:"etag,omitempty"`
	// A dashboard's root container element that defines the layout style.
	//
	// Types that are assignable to Layout:
	//	*Dashboard_GridLayout
	//	*Dashboard_MosaicLayout
	//	*Dashboard_RowLayout
	//	*Dashboard_ColumnLayout
	Layout isDashboard_Layout `protobuf_oneof:"layout"`
}

func (x *Dashboard) Reset() {
	*x = Dashboard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_dashboard_v1_dashboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dashboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dashboard) ProtoMessage() {}

func (x *Dashboard) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_dashboard_v1_dashboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dashboard.ProtoReflect.Descriptor instead.
func (*Dashboard) Descriptor() ([]byte, []int) {
	return file_google_monitoring_dashboard_v1_dashboard_proto_rawDescGZIP(), []int{0}
}

func (x *Dashboard) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Dashboard) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Dashboard) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (m *Dashboard) GetLayout() isDashboard_Layout {
	if m != nil {
		return m.Layout
	}
	return nil
}

func (x *Dashboard) GetGridLayout() *GridLayout {
	if x, ok := x.GetLayout().(*Dashboard_GridLayout); ok {
		return x.GridLayout
	}
	return nil
}

func (x *Dashboard) GetMosaicLayout() *MosaicLayout {
	if x, ok := x.GetLayout().(*Dashboard_MosaicLayout); ok {
		return x.MosaicLayout
	}
	return nil
}

func (x *Dashboard) GetRowLayout() *RowLayout {
	if x, ok := x.GetLayout().(*Dashboard_RowLayout); ok {
		return x.RowLayout
	}
	return nil
}

func (x *Dashboard) GetColumnLayout() *ColumnLayout {
	if x, ok := x.GetLayout().(*Dashboard_ColumnLayout); ok {
		return x.ColumnLayout
	}
	return nil
}

type isDashboard_Layout interface {
	isDashboard_Layout()
}

type Dashboard_GridLayout struct {
	// Content is arranged with a basic layout that re-flows a simple list of
	// informational elements like widgets or tiles.
	GridLayout *GridLayout `protobuf:"bytes,5,opt,name=grid_layout,json=gridLayout,proto3,oneof"`
}

type Dashboard_MosaicLayout struct {
	// The content is arranged as a grid of tiles, with each content widget
	// occupying one or more grid blocks.
	MosaicLayout *MosaicLayout `protobuf:"bytes,6,opt,name=mosaic_layout,json=mosaicLayout,proto3,oneof"`
}

type Dashboard_RowLayout struct {
	// The content is divided into equally spaced rows and the widgets are
	// arranged horizontally.
	RowLayout *RowLayout `protobuf:"bytes,8,opt,name=row_layout,json=rowLayout,proto3,oneof"`
}

type Dashboard_ColumnLayout struct {
	// The content is divided into equally spaced columns and the widgets are
	// arranged vertically.
	ColumnLayout *ColumnLayout `protobuf:"bytes,9,opt,name=column_layout,json=columnLayout,proto3,oneof"`
}

func (*Dashboard_GridLayout) isDashboard_Layout() {}

func (*Dashboard_MosaicLayout) isDashboard_Layout() {}

func (*Dashboard_RowLayout) isDashboard_Layout() {}

func (*Dashboard_ColumnLayout) isDashboard_Layout() {}

var File_google_monitoring_dashboard_v1_dashboard_proto protoreflect.FileDescriptor

var file_google_monitoring_dashboard_v1_dashboard_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x04, 0x0a, 0x09, 0x44,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x12, 0x4d, 0x0a,
	0x0b, 0x67, 0x72, 0x69, 0x64, 0x5f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x69, 0x64, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x48, 0x00,
	0x52, 0x0a, 0x67, 0x72, 0x69, 0x64, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x53, 0x0a, 0x0d,
	0x6d, 0x6f, 0x73, 0x61, 0x69, 0x63, 0x5f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x73, 0x61, 0x69, 0x63, 0x4c, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x48, 0x00, 0x52, 0x0c, 0x6d, 0x6f, 0x73, 0x61, 0x69, 0x63, 0x4c, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x12, 0x4a, 0x0a, 0x0a, 0x72, 0x6f, 0x77, 0x5f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x77, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74,
	0x48, 0x00, 0x52, 0x09, 0x72, 0x6f, 0x77, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x53, 0x0a,
	0x0d, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4c, 0x61, 0x79, 0x6f,
	0x75, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4c, 0x61, 0x79, 0x6f,
	0x75, 0x74, 0x3a, 0x53, 0xea, 0x41, 0x50, 0x0a, 0x23, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x29, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d,
	0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x7d, 0x42, 0x08, 0x0a, 0x06, 0x6c, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x42, 0xab, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0xea, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x3a, 0x3a, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_monitoring_dashboard_v1_dashboard_proto_rawDescOnce sync.Once
	file_google_monitoring_dashboard_v1_dashboard_proto_rawDescData = file_google_monitoring_dashboard_v1_dashboard_proto_rawDesc
)

func file_google_monitoring_dashboard_v1_dashboard_proto_rawDescGZIP() []byte {
	file_google_monitoring_dashboard_v1_dashboard_proto_rawDescOnce.Do(func() {
		file_google_monitoring_dashboard_v1_dashboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_monitoring_dashboard_v1_dashboard_proto_rawDescData)
	})
	return file_google_monitoring_dashboard_v1_dashboard_proto_rawDescData
}

var file_google_monitoring_dashboard_v1_dashboard_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_monitoring_dashboard_v1_dashboard_proto_goTypes = []interface{}{
	(*Dashboard)(nil),    // 0: google.monitoring.dashboard.v1.Dashboard
	(*GridLayout)(nil),   // 1: google.monitoring.dashboard.v1.GridLayout
	(*MosaicLayout)(nil), // 2: google.monitoring.dashboard.v1.MosaicLayout
	(*RowLayout)(nil),    // 3: google.monitoring.dashboard.v1.RowLayout
	(*ColumnLayout)(nil), // 4: google.monitoring.dashboard.v1.ColumnLayout
}
var file_google_monitoring_dashboard_v1_dashboard_proto_depIdxs = []int32{
	1, // 0: google.monitoring.dashboard.v1.Dashboard.grid_layout:type_name -> google.monitoring.dashboard.v1.GridLayout
	2, // 1: google.monitoring.dashboard.v1.Dashboard.mosaic_layout:type_name -> google.monitoring.dashboard.v1.MosaicLayout
	3, // 2: google.monitoring.dashboard.v1.Dashboard.row_layout:type_name -> google.monitoring.dashboard.v1.RowLayout
	4, // 3: google.monitoring.dashboard.v1.Dashboard.column_layout:type_name -> google.monitoring.dashboard.v1.ColumnLayout
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_monitoring_dashboard_v1_dashboard_proto_init() }
func file_google_monitoring_dashboard_v1_dashboard_proto_init() {
	if File_google_monitoring_dashboard_v1_dashboard_proto != nil {
		return
	}
	file_google_monitoring_dashboard_v1_layouts_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_monitoring_dashboard_v1_dashboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dashboard); i {
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
	file_google_monitoring_dashboard_v1_dashboard_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Dashboard_GridLayout)(nil),
		(*Dashboard_MosaicLayout)(nil),
		(*Dashboard_RowLayout)(nil),
		(*Dashboard_ColumnLayout)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_monitoring_dashboard_v1_dashboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_monitoring_dashboard_v1_dashboard_proto_goTypes,
		DependencyIndexes: file_google_monitoring_dashboard_v1_dashboard_proto_depIdxs,
		MessageInfos:      file_google_monitoring_dashboard_v1_dashboard_proto_msgTypes,
	}.Build()
	File_google_monitoring_dashboard_v1_dashboard_proto = out.File
	file_google_monitoring_dashboard_v1_dashboard_proto_rawDesc = nil
	file_google_monitoring_dashboard_v1_dashboard_proto_goTypes = nil
	file_google_monitoring_dashboard_v1_dashboard_proto_depIdxs = nil
}

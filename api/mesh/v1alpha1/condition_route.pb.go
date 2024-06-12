// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: api/mesh/v1alpha1/condition_route.proto

package v1alpha1

import (
	reflect "reflect"
	sync "sync"
)

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

import (
	_ "github.com/apache/dubbo-kubernetes/api/mesh"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConditionRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigVersion string   `protobuf:"bytes,1,opt,name=configVersion,proto3" json:"configVersion,omitempty"`
	Priority      int32    `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	Enabled       bool     `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Force         bool     `protobuf:"varint,4,opt,name=force,proto3" json:"force,omitempty"`
	Runtime       bool     `protobuf:"varint,5,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Key           string   `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
	Scope         string   `protobuf:"bytes,7,opt,name=scope,proto3" json:"scope,omitempty"`
	Conditions    []string `protobuf:"bytes,8,rep,name=conditions,proto3" json:"conditions,omitempty"`
}

func (x *ConditionRoute) Reset() {
	*x = ConditionRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mesh_v1alpha1_condition_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConditionRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConditionRoute) ProtoMessage() {}

func (x *ConditionRoute) ProtoReflect() protoreflect.Message {
	mi := &file_api_mesh_v1alpha1_condition_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConditionRoute.ProtoReflect.Descriptor instead.
func (*ConditionRoute) Descriptor() ([]byte, []int) {
	return file_api_mesh_v1alpha1_condition_route_proto_rawDescGZIP(), []int{0}
}

func (x *ConditionRoute) GetConfigVersion() string {
	if x != nil {
		return x.ConfigVersion
	}
	return ""
}

func (x *ConditionRoute) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *ConditionRoute) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *ConditionRoute) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

func (x *ConditionRoute) GetRuntime() bool {
	if x != nil {
		return x.Runtime
	}
	return false
}

func (x *ConditionRoute) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ConditionRoute) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *ConditionRoute) GetConditions() []string {
	if x != nil {
		return x.Conditions
	}
	return nil
}

var File_api_mesh_v1alpha1_condition_route_proto protoreflect.FileDescriptor

var file_api_mesh_v1alpha1_condition_route_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x75, 0x62, 0x62, 0x6f,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x16,
	0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x83, 0x01,
	0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x18, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0xaa, 0x8c,
	0x89, 0xa6, 0x01, 0x10, 0x12, 0x0e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x06, 0x22, 0x04, 0x6d, 0x65, 0x73, 0x68,
	0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x04, 0x52, 0x02, 0x10, 0x01, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x12,
	0x3a, 0x10, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x13, 0x3a, 0x11, 0x12, 0x0f, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0xaa, 0x8c, 0x89, 0xa6, 0x01,
	0x02, 0x68, 0x01, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2d, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_mesh_v1alpha1_condition_route_proto_rawDescOnce sync.Once
	file_api_mesh_v1alpha1_condition_route_proto_rawDescData = file_api_mesh_v1alpha1_condition_route_proto_rawDesc
)

func file_api_mesh_v1alpha1_condition_route_proto_rawDescGZIP() []byte {
	file_api_mesh_v1alpha1_condition_route_proto_rawDescOnce.Do(func() {
		file_api_mesh_v1alpha1_condition_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_mesh_v1alpha1_condition_route_proto_rawDescData)
	})
	return file_api_mesh_v1alpha1_condition_route_proto_rawDescData
}

var file_api_mesh_v1alpha1_condition_route_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_mesh_v1alpha1_condition_route_proto_goTypes = []interface{}{
	(*ConditionRoute)(nil), // 0: dubbo.mesh.v1alpha1.ConditionRoute
}
var file_api_mesh_v1alpha1_condition_route_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_mesh_v1alpha1_condition_route_proto_init() }
func file_api_mesh_v1alpha1_condition_route_proto_init() {
	if File_api_mesh_v1alpha1_condition_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_mesh_v1alpha1_condition_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConditionRoute); i {
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
			RawDescriptor: file_api_mesh_v1alpha1_condition_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_mesh_v1alpha1_condition_route_proto_goTypes,
		DependencyIndexes: file_api_mesh_v1alpha1_condition_route_proto_depIdxs,
		MessageInfos:      file_api_mesh_v1alpha1_condition_route_proto_msgTypes,
	}.Build()
	File_api_mesh_v1alpha1_condition_route_proto = out.File
	file_api_mesh_v1alpha1_condition_route_proto_rawDesc = nil
	file_api_mesh_v1alpha1_condition_route_proto_goTypes = nil
	file_api_mesh_v1alpha1_condition_route_proto_depIdxs = nil
}

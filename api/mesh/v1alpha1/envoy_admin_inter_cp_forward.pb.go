// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: api/mesh/v1alpha1/envoy_admin_inter_cp_forward.proto

package v1alpha1

import (
	reflect "reflect"
)

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_mesh_v1alpha1_envoy_admin_inter_cp_forward_proto protoreflect.FileDescriptor

var file_api_mesh_v1alpha1_envoy_admin_inter_cp_forward_proto_rawDesc = []byte{
	0x0a, 0x34, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x70, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x61, 0x70, 0x69,
	0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64,
	0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa6, 0x02, 0x0a, 0x1f, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x43, 0x50, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x46, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x09,
	0x58, 0x44, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x2e, 0x64, 0x75, 0x62, 0x62,
	0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x58, 0x44, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x58, 0x44, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x21, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x08, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x24, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x64, 0x75, 0x62,
	0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2d, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_api_mesh_v1alpha1_envoy_admin_inter_cp_forward_proto_goTypes = []interface{}{
	(*XDSConfigRequest)(nil),  // 0: dubbo.mesh.v1alpha1.XDSConfigRequest
	(*StatsRequest)(nil),      // 1: dubbo.mesh.v1alpha1.StatsRequest
	(*ClustersRequest)(nil),   // 2: dubbo.mesh.v1alpha1.ClustersRequest
	(*XDSConfigResponse)(nil), // 3: dubbo.mesh.v1alpha1.XDSConfigResponse
	(*StatsResponse)(nil),     // 4: dubbo.mesh.v1alpha1.StatsResponse
	(*ClustersResponse)(nil),  // 5: dubbo.mesh.v1alpha1.ClustersResponse
}
var file_api_mesh_v1alpha1_envoy_admin_inter_cp_forward_proto_depIdxs = []int32{
	0, // 0: dubbo.mesh.v1alpha1.InterCPEnvoyAdminForwardService.XDSConfig:input_type -> dubbo.mesh.v1alpha1.XDSConfigRequest
	1, // 1: dubbo.mesh.v1alpha1.InterCPEnvoyAdminForwardService.Stats:input_type -> dubbo.mesh.v1alpha1.StatsRequest
	2, // 2: dubbo.mesh.v1alpha1.InterCPEnvoyAdminForwardService.Clusters:input_type -> dubbo.mesh.v1alpha1.ClustersRequest
	3, // 3: dubbo.mesh.v1alpha1.InterCPEnvoyAdminForwardService.XDSConfig:output_type -> dubbo.mesh.v1alpha1.XDSConfigResponse
	4, // 4: dubbo.mesh.v1alpha1.InterCPEnvoyAdminForwardService.Stats:output_type -> dubbo.mesh.v1alpha1.StatsResponse
	5, // 5: dubbo.mesh.v1alpha1.InterCPEnvoyAdminForwardService.Clusters:output_type -> dubbo.mesh.v1alpha1.ClustersResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_mesh_v1alpha1_envoy_admin_inter_cp_forward_proto_init() }
func file_api_mesh_v1alpha1_envoy_admin_inter_cp_forward_proto_init() {
	if File_api_mesh_v1alpha1_envoy_admin_inter_cp_forward_proto != nil {
		return
	}
	file_api_mesh_v1alpha1_dds_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_mesh_v1alpha1_envoy_admin_inter_cp_forward_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_mesh_v1alpha1_envoy_admin_inter_cp_forward_proto_goTypes,
		DependencyIndexes: file_api_mesh_v1alpha1_envoy_admin_inter_cp_forward_proto_depIdxs,
	}.Build()
	File_api_mesh_v1alpha1_envoy_admin_inter_cp_forward_proto = out.File
	file_api_mesh_v1alpha1_envoy_admin_inter_cp_forward_proto_rawDesc = nil
	file_api_mesh_v1alpha1_envoy_admin_inter_cp_forward_proto_goTypes = nil
	file_api_mesh_v1alpha1_envoy_admin_inter_cp_forward_proto_depIdxs = nil
}
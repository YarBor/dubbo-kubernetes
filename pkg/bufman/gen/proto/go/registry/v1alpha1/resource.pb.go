// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: registry/v1alpha1/resource.proto

package registryv1alpha1

import (
	reflect "reflect"
	sync "sync"
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

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Resource:
	//
	//	*Resource_Repository
	//	*Resource_Plugin
	Resource isResource_Resource `protobuf_oneof:"resource"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1alpha1_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1alpha1_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_registry_v1alpha1_resource_proto_rawDescGZIP(), []int{0}
}

func (m *Resource) GetResource() isResource_Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (x *Resource) GetRepository() *Repository {
	if x, ok := x.GetResource().(*Resource_Repository); ok {
		return x.Repository
	}
	return nil
}

func (x *Resource) GetPlugin() *CuratedPlugin {
	if x, ok := x.GetResource().(*Resource_Plugin); ok {
		return x.Plugin
	}
	return nil
}

type isResource_Resource interface {
	isResource_Resource()
}

type Resource_Repository struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository,proto3,oneof"`
}

type Resource_Plugin struct {
	Plugin *CuratedPlugin `protobuf:"bytes,2,opt,name=plugin,proto3,oneof"`
}

func (*Resource_Repository) isResource_Resource() {}

func (*Resource_Plugin) isResource_Resource() {}

type GetResourceByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Owner of the requested resource.
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// Name of the requested resource.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetResourceByNameRequest) Reset() {
	*x = GetResourceByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1alpha1_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceByNameRequest) ProtoMessage() {}

func (x *GetResourceByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1alpha1_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceByNameRequest.ProtoReflect.Descriptor instead.
func (*GetResourceByNameRequest) Descriptor() ([]byte, []int) {
	return file_registry_v1alpha1_resource_proto_rawDescGZIP(), []int{1}
}

func (x *GetResourceByNameRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *GetResourceByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetResourceByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *GetResourceByNameResponse) Reset() {
	*x = GetResourceByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1alpha1_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceByNameResponse) ProtoMessage() {}

func (x *GetResourceByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1alpha1_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceByNameResponse.ProtoReflect.Descriptor instead.
func (*GetResourceByNameResponse) Descriptor() ([]byte, []int) {
	return file_registry_v1alpha1_resource_proto_rawDescGZIP(), []int{2}
}

func (x *GetResourceByNameResponse) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

var File_registry_v1alpha1_resource_proto protoreflect.FileDescriptor

var file_registry_v1alpha1_resource_proto_rawDesc = []byte{
	0x0a, 0x20, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x29, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f,
	0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x27, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x62, 0x75,
	0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x52, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x38, 0x2e, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e,
	0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x75, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x22, 0x44, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6c, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64,
	0x75, 0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x32, 0xb7, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa3, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x43,
	0x2e, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70,
	0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62,
	0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x01, 0x42, 0xe8,
	0x02, 0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75,
	0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x42, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70,
	0x61, 0x63, 0x68, 0x65, 0x2f, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xa2, 0x02, 0x05, 0x42, 0x44, 0x41, 0x4f, 0x52, 0xaa, 0x02, 0x29, 0x42, 0x75, 0x66, 0x6d, 0x61,
	0x6e, 0x2e, 0x44, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x4f,
	0x72, 0x67, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x29, 0x42, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x5c, 0x44, 0x75,
	0x62, 0x62, 0x6f, 0x5c, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xe2, 0x02, 0x35, 0x42, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x5c, 0x44, 0x75, 0x62, 0x62, 0x6f, 0x5c,
	0x41, 0x70, 0x61, 0x63, 0x68, 0x65, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x2e, 0x42, 0x75, 0x66, 0x6d, 0x61,
	0x6e, 0x3a, 0x3a, 0x44, 0x75, 0x62, 0x62, 0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65,
	0x3a, 0x3a, 0x4f, 0x72, 0x67, 0x3a, 0x3a, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x3a,
	0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_registry_v1alpha1_resource_proto_rawDescOnce sync.Once
	file_registry_v1alpha1_resource_proto_rawDescData = file_registry_v1alpha1_resource_proto_rawDesc
)

func file_registry_v1alpha1_resource_proto_rawDescGZIP() []byte {
	file_registry_v1alpha1_resource_proto_rawDescOnce.Do(func() {
		file_registry_v1alpha1_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_registry_v1alpha1_resource_proto_rawDescData)
	})
	return file_registry_v1alpha1_resource_proto_rawDescData
}

var file_registry_v1alpha1_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_registry_v1alpha1_resource_proto_goTypes = []interface{}{
	(*Resource)(nil),                  // 0: bufman.dubbo.apache.org.registry.v1alpha1.Resource
	(*GetResourceByNameRequest)(nil),  // 1: bufman.dubbo.apache.org.registry.v1alpha1.GetResourceByNameRequest
	(*GetResourceByNameResponse)(nil), // 2: bufman.dubbo.apache.org.registry.v1alpha1.GetResourceByNameResponse
	(*Repository)(nil),                // 3: bufman.dubbo.apache.org.registry.v1alpha1.Repository
	(*CuratedPlugin)(nil),             // 4: bufman.dubbo.apache.org.registry.v1alpha1.CuratedPlugin
}
var file_registry_v1alpha1_resource_proto_depIdxs = []int32{
	3, // 0: bufman.dubbo.apache.org.registry.v1alpha1.Resource.repository:type_name -> bufman.dubbo.apache.org.registry.v1alpha1.Repository
	4, // 1: bufman.dubbo.apache.org.registry.v1alpha1.Resource.plugin:type_name -> bufman.dubbo.apache.org.registry.v1alpha1.CuratedPlugin
	0, // 2: bufman.dubbo.apache.org.registry.v1alpha1.GetResourceByNameResponse.resource:type_name -> bufman.dubbo.apache.org.registry.v1alpha1.Resource
	1, // 3: bufman.dubbo.apache.org.registry.v1alpha1.ResourceService.GetResourceByName:input_type -> bufman.dubbo.apache.org.registry.v1alpha1.GetResourceByNameRequest
	2, // 4: bufman.dubbo.apache.org.registry.v1alpha1.ResourceService.GetResourceByName:output_type -> bufman.dubbo.apache.org.registry.v1alpha1.GetResourceByNameResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_registry_v1alpha1_resource_proto_init() }
func file_registry_v1alpha1_resource_proto_init() {
	if File_registry_v1alpha1_resource_proto != nil {
		return
	}
	file_registry_v1alpha1_plugin_curation_proto_init()
	file_registry_v1alpha1_repository_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_registry_v1alpha1_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_registry_v1alpha1_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceByNameRequest); i {
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
		file_registry_v1alpha1_resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceByNameResponse); i {
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
	file_registry_v1alpha1_resource_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Resource_Repository)(nil),
		(*Resource_Plugin)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_registry_v1alpha1_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_registry_v1alpha1_resource_proto_goTypes,
		DependencyIndexes: file_registry_v1alpha1_resource_proto_depIdxs,
		MessageInfos:      file_registry_v1alpha1_resource_proto_msgTypes,
	}.Build()
	File_registry_v1alpha1_resource_proto = out.File
	file_registry_v1alpha1_resource_proto_rawDesc = nil
	file_registry_v1alpha1_resource_proto_goTypes = nil
	file_registry_v1alpha1_resource_proto_depIdxs = nil
}
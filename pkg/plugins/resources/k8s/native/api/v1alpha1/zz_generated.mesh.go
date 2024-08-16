// Generated by tools/resource-gen
// Run "make generate" to update this file.

// nolint:whitespace
package v1alpha1

import (
	"fmt"
)

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

import (
	mesh_proto "github.com/apache/dubbo-kubernetes/api/mesh/v1alpha1"
	core_model "github.com/apache/dubbo-kubernetes/pkg/core/resources/model"
	"github.com/apache/dubbo-kubernetes/pkg/plugins/resources/k8s/native/pkg/model"
	"github.com/apache/dubbo-kubernetes/pkg/plugins/resources/k8s/native/pkg/registry"
	util_proto "github.com/apache/dubbo-kubernetes/pkg/util/proto"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=dubbo,scope=Cluster
type AffinityRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Mesh is the name of the dubbo mesh this resource belongs to.
	// It may be omitted for cluster-scoped resources.
	//
	// +kubebuilder:validation:Optional
	Mesh string `json:"mesh,omitempty"`
	// Spec is the specification of the Dubbo AffinityRoute resource.
	// +kubebuilder:validation:Optional
	Spec *apiextensionsv1.JSON `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type AffinityRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AffinityRoute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AffinityRoute{}, &AffinityRouteList{})
}

func (cb *AffinityRoute) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *AffinityRoute) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *AffinityRoute) GetMesh() string {
	return cb.Mesh
}

func (cb *AffinityRoute) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *AffinityRoute) GetSpec() (core_model.ResourceSpec, error) {
	spec := cb.Spec
	m := mesh_proto.AffinityRoute{}

	if spec == nil || len(spec.Raw) == 0 {
		return &m, nil
	}

	err := util_proto.FromJSON(spec.Raw, &m)
	return &m, err
}

func (cb *AffinityRoute) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
		cb.Spec = nil
		return
	}

	s, ok := spec.(*mesh_proto.AffinityRoute)
	if !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

	cb.Spec = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
}

func (cb *AffinityRoute) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *AffinityRouteList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.AffinityRoute{}, &AffinityRoute{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "AffinityRoute",
		},
	})
	registry.RegisterListType(&mesh_proto.AffinityRoute{}, &AffinityRouteList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "AffinityRouteList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=dubbo,scope=Cluster
type ConditionRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Mesh is the name of the dubbo mesh this resource belongs to.
	// It may be omitted for cluster-scoped resources.
	//
	// +kubebuilder:validation:Optional
	Mesh string `json:"mesh,omitempty"`
	// Spec is the specification of the Dubbo ConditionRoute resource.
	// +kubebuilder:validation:Optional
	Spec *apiextensionsv1.JSON `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type ConditionRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConditionRoute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConditionRoute{}, &ConditionRouteList{})
}

func (cb *ConditionRoute) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *ConditionRoute) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *ConditionRoute) GetMesh() string {
	return cb.Mesh
}

func (cb *ConditionRoute) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *ConditionRoute) GetSpec() (core_model.ResourceSpec, error) {
	spec := cb.Spec
	m := mesh_proto.ConditionRoute{}

	if spec == nil || len(spec.Raw) == 0 {
		return &m, nil
	}

	err := util_proto.FromJSON(spec.Raw, &m)
	return &m, err
}

func (cb *ConditionRoute) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
		cb.Spec = nil
		return
	}

	s, ok := spec.(*mesh_proto.ConditionRoute)
	if !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

	cb.Spec = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
}

func (cb *ConditionRoute) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *ConditionRouteList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.ConditionRoute{}, &ConditionRoute{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ConditionRoute",
		},
	})
	registry.RegisterListType(&mesh_proto.ConditionRoute{}, &ConditionRouteList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ConditionRouteList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=dubbo,scope=Namespaced
type Dataplane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Mesh is the name of the dubbo mesh this resource belongs to.
	// It may be omitted for cluster-scoped resources.
	//
	// +kubebuilder:validation:Optional
	Mesh string `json:"mesh,omitempty"`
	// Spec is the specification of the Dubbo Dataplane resource.
	// +kubebuilder:validation:Optional
	Spec *apiextensionsv1.JSON `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type DataplaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dataplane `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Dataplane{}, &DataplaneList{})
}

func (cb *Dataplane) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *Dataplane) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *Dataplane) GetMesh() string {
	return cb.Mesh
}

func (cb *Dataplane) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *Dataplane) GetSpec() (core_model.ResourceSpec, error) {
	spec := cb.Spec
	m := mesh_proto.Dataplane{}

	if spec == nil || len(spec.Raw) == 0 {
		return &m, nil
	}

	err := util_proto.FromJSON(spec.Raw, &m)
	return &m, err
}

func (cb *Dataplane) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
		cb.Spec = nil
		return
	}

	s, ok := spec.(*mesh_proto.Dataplane)
	if !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

	cb.Spec = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
}

func (cb *Dataplane) Scope() model.Scope {
	return model.ScopeNamespace
}

func (l *DataplaneList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.Dataplane{}, &Dataplane{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "Dataplane",
		},
	})
	registry.RegisterListType(&mesh_proto.Dataplane{}, &DataplaneList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "DataplaneList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=dubbo,scope=Namespaced
type DataplaneInsight struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Mesh is the name of the dubbo mesh this resource belongs to.
	// It may be omitted for cluster-scoped resources.
	//
	// +kubebuilder:validation:Optional
	Mesh string `json:"mesh,omitempty"`
	// Status is the status the dubbo resource.
	// +kubebuilder:validation:Optional
	Status *apiextensionsv1.JSON `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type DataplaneInsightList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplaneInsight `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplaneInsight{}, &DataplaneInsightList{})
}

func (cb *DataplaneInsight) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *DataplaneInsight) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *DataplaneInsight) GetMesh() string {
	return cb.Mesh
}

func (cb *DataplaneInsight) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *DataplaneInsight) GetSpec() (core_model.ResourceSpec, error) {
	spec := cb.Status
	m := mesh_proto.DataplaneInsight{}

	if spec == nil || len(spec.Raw) == 0 {
		return &m, nil
	}

	err := util_proto.FromJSON(spec.Raw, &m)
	return &m, err
}

func (cb *DataplaneInsight) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
		cb.Status = nil
		return
	}

	s, ok := spec.(*mesh_proto.DataplaneInsight)
	if !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

	cb.Status = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
}

func (cb *DataplaneInsight) Scope() model.Scope {
	return model.ScopeNamespace
}

func (l *DataplaneInsightList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.DataplaneInsight{}, &DataplaneInsight{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "DataplaneInsight",
		},
	})
	registry.RegisterListType(&mesh_proto.DataplaneInsight{}, &DataplaneInsightList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "DataplaneInsightList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=dubbo,scope=Cluster
type DynamicConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Mesh is the name of the dubbo mesh this resource belongs to.
	// It may be omitted for cluster-scoped resources.
	//
	// +kubebuilder:validation:Optional
	Mesh string `json:"mesh,omitempty"`
	// Spec is the specification of the Dubbo DynamicConfig resource.
	// +kubebuilder:validation:Optional
	Spec *apiextensionsv1.JSON `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type DynamicConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DynamicConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DynamicConfig{}, &DynamicConfigList{})
}

func (cb *DynamicConfig) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *DynamicConfig) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *DynamicConfig) GetMesh() string {
	return cb.Mesh
}

func (cb *DynamicConfig) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *DynamicConfig) GetSpec() (core_model.ResourceSpec, error) {
	spec := cb.Spec
	m := mesh_proto.DynamicConfig{}

	if spec == nil || len(spec.Raw) == 0 {
		return &m, nil
	}

	err := util_proto.FromJSON(spec.Raw, &m)
	return &m, err
}

func (cb *DynamicConfig) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
		cb.Spec = nil
		return
	}

	s, ok := spec.(*mesh_proto.DynamicConfig)
	if !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

	cb.Spec = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
}

func (cb *DynamicConfig) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *DynamicConfigList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.DynamicConfig{}, &DynamicConfig{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "DynamicConfig",
		},
	})
	registry.RegisterListType(&mesh_proto.DynamicConfig{}, &DynamicConfigList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "DynamicConfigList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=dubbo,scope=Namespaced
type Mapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Mesh is the name of the dubbo mesh this resource belongs to.
	// It may be omitted for cluster-scoped resources.
	//
	// +kubebuilder:validation:Optional
	Mesh string `json:"mesh,omitempty"`
	// Spec is the specification of the Dubbo Mapping resource.
	// +kubebuilder:validation:Optional
	Spec *apiextensionsv1.JSON `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type MappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Mapping `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Mapping{}, &MappingList{})
}

func (cb *Mapping) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *Mapping) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *Mapping) GetMesh() string {
	return cb.Mesh
}

func (cb *Mapping) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *Mapping) GetSpec() (core_model.ResourceSpec, error) {
	spec := cb.Spec
	m := mesh_proto.Mapping{}

	if spec == nil || len(spec.Raw) == 0 {
		return &m, nil
	}

	err := util_proto.FromJSON(spec.Raw, &m)
	return &m, err
}

func (cb *Mapping) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
		cb.Spec = nil
		return
	}

	s, ok := spec.(*mesh_proto.Mapping)
	if !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

	cb.Spec = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
}

func (cb *Mapping) Scope() model.Scope {
	return model.ScopeNamespace
}

func (l *MappingList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.Mapping{}, &Mapping{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "Mapping",
		},
	})
	registry.RegisterListType(&mesh_proto.Mapping{}, &MappingList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "MappingList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=dubbo,scope=Cluster
type Mesh struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Mesh is the name of the dubbo mesh this resource belongs to.
	// It may be omitted for cluster-scoped resources.
	//
	// +kubebuilder:validation:Optional
	Mesh string `json:"mesh,omitempty"`
	// Spec is the specification of the Dubbo Mesh resource.
	// +kubebuilder:validation:Optional
	Spec *apiextensionsv1.JSON `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type MeshList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Mesh `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Mesh{}, &MeshList{})
}

func (cb *Mesh) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *Mesh) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *Mesh) GetMesh() string {
	return cb.Mesh
}

func (cb *Mesh) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *Mesh) GetSpec() (core_model.ResourceSpec, error) {
	spec := cb.Spec
	m := mesh_proto.Mesh{}

	if spec == nil || len(spec.Raw) == 0 {
		return &m, nil
	}

	err := util_proto.FromJSON(spec.Raw, &m)
	return &m, err
}

func (cb *Mesh) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
		cb.Spec = nil
		return
	}

	s, ok := spec.(*mesh_proto.Mesh)
	if !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

	cb.Spec = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
}

func (cb *Mesh) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *MeshList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.Mesh{}, &Mesh{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "Mesh",
		},
	})
	registry.RegisterListType(&mesh_proto.Mesh{}, &MeshList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "MeshList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=dubbo,scope=Cluster
type MeshInsight struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Mesh is the name of the dubbo mesh this resource belongs to.
	// It may be omitted for cluster-scoped resources.
	//
	// +kubebuilder:validation:Optional
	Mesh string `json:"mesh,omitempty"`
	// Spec is the specification of the Dubbo MeshInsight resource.
	// +kubebuilder:validation:Optional
	Spec *apiextensionsv1.JSON `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type MeshInsightList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MeshInsight `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MeshInsight{}, &MeshInsightList{})
}

func (cb *MeshInsight) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *MeshInsight) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *MeshInsight) GetMesh() string {
	return cb.Mesh
}

func (cb *MeshInsight) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *MeshInsight) GetSpec() (core_model.ResourceSpec, error) {
	spec := cb.Spec
	m := mesh_proto.MeshInsight{}

	if spec == nil || len(spec.Raw) == 0 {
		return &m, nil
	}

	err := util_proto.FromJSON(spec.Raw, &m)
	return &m, err
}

func (cb *MeshInsight) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
		cb.Spec = nil
		return
	}

	s, ok := spec.(*mesh_proto.MeshInsight)
	if !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

	cb.Spec = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
}

func (cb *MeshInsight) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *MeshInsightList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.MeshInsight{}, &MeshInsight{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "MeshInsight",
		},
	})
	registry.RegisterListType(&mesh_proto.MeshInsight{}, &MeshInsightList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "MeshInsightList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=dubbo,scope=Namespaced
type MetaData struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Mesh is the name of the dubbo mesh this resource belongs to.
	// It may be omitted for cluster-scoped resources.
	//
	// +kubebuilder:validation:Optional
	Mesh string `json:"mesh,omitempty"`
	// Spec is the specification of the Dubbo MetaData resource.
	// +kubebuilder:validation:Optional
	Spec *apiextensionsv1.JSON `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type MetaDataList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetaData `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MetaData{}, &MetaDataList{})
}

func (cb *MetaData) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *MetaData) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *MetaData) GetMesh() string {
	return cb.Mesh
}

func (cb *MetaData) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *MetaData) GetSpec() (core_model.ResourceSpec, error) {
	spec := cb.Spec
	m := mesh_proto.MetaData{}

	if spec == nil || len(spec.Raw) == 0 {
		return &m, nil
	}

	err := util_proto.FromJSON(spec.Raw, &m)
	return &m, err
}

func (cb *MetaData) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
		cb.Spec = nil
		return
	}

	s, ok := spec.(*mesh_proto.MetaData)
	if !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

	cb.Spec = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
}

func (cb *MetaData) Scope() model.Scope {
	return model.ScopeNamespace
}

func (l *MetaDataList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.MetaData{}, &MetaData{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "MetaData",
		},
	})
	registry.RegisterListType(&mesh_proto.MetaData{}, &MetaDataList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "MetaDataList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=dubbo,scope=Cluster
type TagRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Mesh is the name of the dubbo mesh this resource belongs to.
	// It may be omitted for cluster-scoped resources.
	//
	// +kubebuilder:validation:Optional
	Mesh string `json:"mesh,omitempty"`
	// Spec is the specification of the Dubbo TagRoute resource.
	// +kubebuilder:validation:Optional
	Spec *apiextensionsv1.JSON `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type TagRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TagRoute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TagRoute{}, &TagRouteList{})
}

func (cb *TagRoute) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *TagRoute) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *TagRoute) GetMesh() string {
	return cb.Mesh
}

func (cb *TagRoute) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *TagRoute) GetSpec() (core_model.ResourceSpec, error) {
	spec := cb.Spec
	m := mesh_proto.TagRoute{}

	if spec == nil || len(spec.Raw) == 0 {
		return &m, nil
	}

	err := util_proto.FromJSON(spec.Raw, &m)
	return &m, err
}

func (cb *TagRoute) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
		cb.Spec = nil
		return
	}

	s, ok := spec.(*mesh_proto.TagRoute)
	if !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

	cb.Spec = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
}

func (cb *TagRoute) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *TagRouteList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.TagRoute{}, &TagRoute{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "TagRoute",
		},
	})
	registry.RegisterListType(&mesh_proto.TagRoute{}, &TagRouteList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "TagRouteList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=dubbo,scope=Cluster
type VirtualService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Mesh is the name of the dubbo mesh this resource belongs to.
	// It may be omitted for cluster-scoped resources.
	//
	// +kubebuilder:validation:Optional
	Mesh string `json:"mesh,omitempty"`
	// Spec is the specification of the Dubbo VirtualService resource.
	// +kubebuilder:validation:Optional
	Spec *apiextensionsv1.JSON `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type VirtualServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualService{}, &VirtualServiceList{})
}

func (cb *VirtualService) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *VirtualService) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *VirtualService) GetMesh() string {
	return cb.Mesh
}

func (cb *VirtualService) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *VirtualService) GetSpec() (core_model.ResourceSpec, error) {
	spec := cb.Spec
	m := mesh_proto.VirtualService{}

	if spec == nil || len(spec.Raw) == 0 {
		return &m, nil
	}

	err := util_proto.FromJSON(spec.Raw, &m)
	return &m, err
}

func (cb *VirtualService) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
		cb.Spec = nil
		return
	}

	s, ok := spec.(*mesh_proto.VirtualService)
	if !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

	cb.Spec = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
}

func (cb *VirtualService) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *VirtualServiceList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.VirtualService{}, &VirtualService{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "VirtualService",
		},
	})
	registry.RegisterListType(&mesh_proto.VirtualService{}, &VirtualServiceList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "VirtualServiceList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=dubbo,scope=Namespaced
type ZoneEgress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Mesh is the name of the dubbo mesh this resource belongs to.
	// It may be omitted for cluster-scoped resources.
	//
	// +kubebuilder:validation:Optional
	Mesh string `json:"mesh,omitempty"`
	// Spec is the specification of the Dubbo ZoneEgress resource.
	// +kubebuilder:validation:Optional
	Spec *apiextensionsv1.JSON `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type ZoneEgressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ZoneEgress `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ZoneEgress{}, &ZoneEgressList{})
}

func (cb *ZoneEgress) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *ZoneEgress) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *ZoneEgress) GetMesh() string {
	return cb.Mesh
}

func (cb *ZoneEgress) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *ZoneEgress) GetSpec() (core_model.ResourceSpec, error) {
	spec := cb.Spec
	m := mesh_proto.ZoneEgress{}

	if spec == nil || len(spec.Raw) == 0 {
		return &m, nil
	}

	err := util_proto.FromJSON(spec.Raw, &m)
	return &m, err
}

func (cb *ZoneEgress) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
		cb.Spec = nil
		return
	}

	s, ok := spec.(*mesh_proto.ZoneEgress)
	if !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

	cb.Spec = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
}

func (cb *ZoneEgress) Scope() model.Scope {
	return model.ScopeNamespace
}

func (l *ZoneEgressList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.ZoneEgress{}, &ZoneEgress{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ZoneEgress",
		},
	})
	registry.RegisterListType(&mesh_proto.ZoneEgress{}, &ZoneEgressList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ZoneEgressList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=dubbo,scope=Namespaced
type ZoneEgressInsight struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Mesh is the name of the dubbo mesh this resource belongs to.
	// It may be omitted for cluster-scoped resources.
	//
	// +kubebuilder:validation:Optional
	Mesh string `json:"mesh,omitempty"`
	// Spec is the specification of the Dubbo ZoneEgressInsight resource.
	// +kubebuilder:validation:Optional
	Spec *apiextensionsv1.JSON `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type ZoneEgressInsightList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ZoneEgressInsight `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ZoneEgressInsight{}, &ZoneEgressInsightList{})
}

func (cb *ZoneEgressInsight) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *ZoneEgressInsight) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *ZoneEgressInsight) GetMesh() string {
	return cb.Mesh
}

func (cb *ZoneEgressInsight) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *ZoneEgressInsight) GetSpec() (core_model.ResourceSpec, error) {
	spec := cb.Spec
	m := mesh_proto.ZoneEgressInsight{}

	if spec == nil || len(spec.Raw) == 0 {
		return &m, nil
	}

	err := util_proto.FromJSON(spec.Raw, &m)
	return &m, err
}

func (cb *ZoneEgressInsight) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
		cb.Spec = nil
		return
	}

	s, ok := spec.(*mesh_proto.ZoneEgressInsight)
	if !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

	cb.Spec = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
}

func (cb *ZoneEgressInsight) Scope() model.Scope {
	return model.ScopeNamespace
}

func (l *ZoneEgressInsightList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.ZoneEgressInsight{}, &ZoneEgressInsight{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ZoneEgressInsight",
		},
	})
	registry.RegisterListType(&mesh_proto.ZoneEgressInsight{}, &ZoneEgressInsightList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ZoneEgressInsightList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=dubbo,scope=Namespaced
type ZoneIngress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Mesh is the name of the dubbo mesh this resource belongs to.
	// It may be omitted for cluster-scoped resources.
	//
	// +kubebuilder:validation:Optional
	Mesh string `json:"mesh,omitempty"`
	// Spec is the specification of the Dubbo ZoneIngress resource.
	// +kubebuilder:validation:Optional
	Spec *apiextensionsv1.JSON `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type ZoneIngressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ZoneIngress `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ZoneIngress{}, &ZoneIngressList{})
}

func (cb *ZoneIngress) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *ZoneIngress) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *ZoneIngress) GetMesh() string {
	return cb.Mesh
}

func (cb *ZoneIngress) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *ZoneIngress) GetSpec() (core_model.ResourceSpec, error) {
	spec := cb.Spec
	m := mesh_proto.ZoneIngress{}

	if spec == nil || len(spec.Raw) == 0 {
		return &m, nil
	}

	err := util_proto.FromJSON(spec.Raw, &m)
	return &m, err
}

func (cb *ZoneIngress) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
		cb.Spec = nil
		return
	}

	s, ok := spec.(*mesh_proto.ZoneIngress)
	if !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

	cb.Spec = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
}

func (cb *ZoneIngress) Scope() model.Scope {
	return model.ScopeNamespace
}

func (l *ZoneIngressList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.ZoneIngress{}, &ZoneIngress{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ZoneIngress",
		},
	})
	registry.RegisterListType(&mesh_proto.ZoneIngress{}, &ZoneIngressList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ZoneIngressList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=dubbo,scope=Namespaced
type ZoneIngressInsight struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Mesh is the name of the dubbo mesh this resource belongs to.
	// It may be omitted for cluster-scoped resources.
	//
	// +kubebuilder:validation:Optional
	Mesh string `json:"mesh,omitempty"`
	// Spec is the specification of the Dubbo ZoneIngressInsight resource.
	// +kubebuilder:validation:Optional
	Spec *apiextensionsv1.JSON `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type ZoneIngressInsightList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ZoneIngressInsight `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ZoneIngressInsight{}, &ZoneIngressInsightList{})
}

func (cb *ZoneIngressInsight) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *ZoneIngressInsight) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *ZoneIngressInsight) GetMesh() string {
	return cb.Mesh
}

func (cb *ZoneIngressInsight) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *ZoneIngressInsight) GetSpec() (core_model.ResourceSpec, error) {
	spec := cb.Spec
	m := mesh_proto.ZoneIngressInsight{}

	if spec == nil || len(spec.Raw) == 0 {
		return &m, nil
	}

	err := util_proto.FromJSON(spec.Raw, &m)
	return &m, err
}

func (cb *ZoneIngressInsight) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
		cb.Spec = nil
		return
	}

	s, ok := spec.(*mesh_proto.ZoneIngressInsight)
	if !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

	cb.Spec = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
}

func (cb *ZoneIngressInsight) Scope() model.Scope {
	return model.ScopeNamespace
}

func (l *ZoneIngressInsightList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.ZoneIngressInsight{}, &ZoneIngressInsight{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ZoneIngressInsight",
		},
	})
	registry.RegisterListType(&mesh_proto.ZoneIngressInsight{}, &ZoneIngressInsightList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ZoneIngressInsightList",
		},
	})
}

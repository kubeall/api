/*
Copyright The Kubeall Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced
// +kubebuilder:resource:scope=Cluster
// WorkspaceRole
type WorkspaceRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Rules             []rbacv1.PolicyRule `json:"rules,omitempty" yaml:"rules" protobuf:"bytes,2,rep,name=rules"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkspaceRoleList contains a list of Workspace
type WorkspaceRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []WorkspaceRole `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced
// +kubebuilder:resource:scope=Cluster
// WorkspaceRoleBinding
type WorkspaceRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Subjects holds references to the objects the role applies to.
	// +optional
	Subjects []rbacv1.Subject `json:"subjects,omitempty" protobuf:"bytes,2,rep,name=subjects"`

	// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace.
	// If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef rbacv1.RoleRef `json:"roleRef" protobuf:"bytes,3,opt,name=roleRef"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkspaceRoleBindingList contains a list of Workspace
type WorkspaceRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []WorkspaceRoleBinding `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced
// +kubebuilder:resource:scope=Cluster
type KubeUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              KubeUserSpec   `json:"spec,omitempty" yaml:"spec" protobuf:"bytes,2,opt,name=spec"`
	Status            KubeUserStatus `protobuf:"bytes,3,opt,name=status"`
}
type KubeUserSpec struct {
	// username,
	UserName string `json:"userName,omitempty" protobuf:"bytes,1,opt,name=userName"`
	// user email
	Email string `json:"email,omitempty" protobuf:"bytes,2,opt,name=email"`
	// default language
	Language string `json:"language,omitempty" protobuf:"bytes,3,opt,name=language"`
	// login password
	Password string `json:"password,omitempty" protobuf:"bytes,4,opt,name=password"`
}
type KubeUserStatus struct {
	Available     bool         `json:"available,omitempty" protobuf:"varint,2,opt,name=available"`
	LastLoginTime *metav1.Time `json:"lastLoginTime,omitempty" protobuf:"bytes,3,opt,name=lastLoginTime"`
	LastRemoteIP  string       `json:"lastRemoteIp,omitempty" protobuf:"bytes,4,opt,name=lastRemoteIp"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeUserList contains a list of KubeUser
type KubeUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []KubeUser `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +genclient:nonNamespaced
// +kubebuilder:printcolumn:name="Cluster",type=string,JSONPath=`.spec.cluster`
// +kubebuilder:resource:scope=Cluster

// UserKubeConfig is the Schema for the usermanages API
type UserKubeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   UserKubeConfigSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status UserKubeConfigStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// UserKubeConfigSpec defines the desired state of UserKubeConfig
type UserKubeConfigSpec struct {
	// cluster resource name
	// +kubebuilder:validation:Required
	Cluster string `json:"cluster,omitempty" protobuf:"bytes,2,opt,name=cluster"`
	// expire time
	ExpirationSeconds *int32 `json:"expirationSeconds,omitempty" protobuf:"varint,3,opt,name=expirationSeconds"`
	// user kubeconfig content
	KubeConfig string `json:"kubeConfig,omitempty" protobuf:"bytes,4,opt,name=kubeConfig"`
}

// UserKubeConfigStatus defines the observed state of UserKubeConfig
type UserKubeConfigStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserKubeConfigList contains a list of UserKubeConfig
type UserKubeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []UserKubeConfig `json:"items" protobuf:"bytes,2,rep,name=items"`
}

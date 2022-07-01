/*
Copyright 2022 The kubeall.com Authors.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
)

const (
	WorkspaceNamespaceSelectorKey = "com.kubeall.workspace.namespace"
	WorkspaceUserSelectorKey      = "com.kubeall.workspace.user"
	UserKubeConfigSelectorKey     = "com.kubeall.user.kubeconfig"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +genclient:nonNamespaced
// +kubebuilder:printcolumn:name="Platform",type=string,JSONPath=`.status.version.platform`
// +kubebuilder:printcolumn:name="Minor",type=string,JSONPath=`.status.version.minor`
// +kubebuilder:printcolumn:name="Healthy",type=string,JSONPath=`.status.healthy`
// +kubebuilder:printcolumn:name="Provider",type=string,JSONPath=`.spec.provider`
// +kubebuilder:printcolumn:name="Category",type=string,JSONPath=`.spec.category`
// +kubebuilder:printcolumn:name="Code",type=string,JSONPath=`.spec.code`
// +kubebuilder:resource:scope=Cluster

// Cluster is the schema for the clusters API
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   ClusterSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status ClusterStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}
type ClusterSpec struct {
	// Provider of the cluster: Openshift, Kubernetes, or other cloud providers
	Provider string `json:"provider,omitempty" protobuf:"bytes,1,opt,name=provider"`
	// cluster category, such as: Strict、NonStrict、Dev、Test、Pro
	Category string `json:"category,omitempty" protobuf:"bytes,2,opt,name=category"`
	//  cluster description
	Description string `json:"description,omitempty" protobuf:"bytes,3,opt,name=description"`
	// cluster code
	Code string `json:"code,omitempty" protobuf:"bytes,4,opt,name=code"`
	// kubernetes
	// +optional
	KubeConfig string `json:"kubeconfig,omitempty" protobuf:"bytes,5,opt,name=kubeconfig"`
	// Cert Data
	// +optional
	CertData string `json:"certData,omitempty" protobuf:"bytes,6,opt,name=certData"`
	// +optional
	Token string `json:"token,omitempty" protobuf:"bytes,7,opt,name=token"`
	//cluster master url
	// +optional
	Master string `json:"master,omitempty" protobuf:"bytes,8,opt,name=master"`
}
type ClusterStatus struct {
	Healthy   bool         `json:"healthy,omitempty" yaml:"healthy" protobuf:"varint,1,opt,name=healthy"`
	LastCheck *metav1.Time `json:"lastCheck,omitempty" yaml:"lastCheck" protobuf:"bytes,2,opt,name=lastCheck"`
	// kubernetes version
	Version *version.Info `json:"version,omitempty" yaml:"version" protobuf:"bytes,3,opt,name=version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Cluster `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced
// +kubebuilder:resource:scope=Cluster

// Workspace is the Schema for the namespacegroups API
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   WorkspaceSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status WorkspaceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}
type WorkspaceSpec struct {
	// workspace code
	// +kubebuilder:validation:Required
	Code string `json:"code,omitempty" yaml:"code" protobuf:"bytes,3,opt,name=code"`
	// workspace description
	// +kubebuilder:validation:Required
	Description string `json:"description,omitempty" protobuf:"bytes,1,opt,name=description"`
	// select which namespace will join to this group
	NamespaceSelectorValue string `json:"namespaceSelector,omitempty" protobuf:"bytes,2,opt,name=namespaceSelector"`
}

// WorkspaceStatus defines the observed state of Workspace
type WorkspaceStatus struct {
	// cluster's namespaces
	ClusterNamespaces []ClusterNamespace `json:"clusterNamespaces,omitempty" protobuf:"bytes,1,rep,name=clusterNamespaces"`
}
type ClusterNamespace struct {
	Cluster    string   `json:"cluster,omitempty" protobuf:"bytes,1,opt,name=cluster"`
	Namespaces []string `json:"namespaces,omitempty" protobuf:"bytes,2,rep,name=namespaces"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkspaceList contains a list of Workspace
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Workspace `json:"items" protobuf:"bytes,2,rep,name=items"`
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

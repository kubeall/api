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
*/package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_KubeUserList = map[string]string{
	"": "KubeUserList contains a list of KubeUser",
}

func (KubeUserList) SwaggerDoc() map[string]string {
	return map_KubeUserList
}

var map_KubeUserSpec = map[string]string{
	"userName": "username,",
	"email":    "user email",
	"language": "default language",
	"password": "login password",
}

func (KubeUserSpec) SwaggerDoc() map[string]string {
	return map_KubeUserSpec
}

var map_UserKubeConfig = map[string]string{
	"": "UserKubeConfig is the Schema for the usermanages API",
}

func (UserKubeConfig) SwaggerDoc() map[string]string {
	return map_UserKubeConfig
}

var map_UserKubeConfigList = map[string]string{
	"": "UserKubeConfigList contains a list of UserKubeConfig",
}

func (UserKubeConfigList) SwaggerDoc() map[string]string {
	return map_UserKubeConfigList
}

var map_UserKubeConfigSpec = map[string]string{
	"":                  "UserKubeConfigSpec defines the desired state of UserKubeConfig",
	"cluster":           "cluster resource name",
	"expirationSeconds": "expire time",
	"kubeConfig":        "user kubeconfig content",
}

func (UserKubeConfigSpec) SwaggerDoc() map[string]string {
	return map_UserKubeConfigSpec
}

var map_UserKubeConfigStatus = map[string]string{
	"": "UserKubeConfigStatus defines the observed state of UserKubeConfig",
}

func (UserKubeConfigStatus) SwaggerDoc() map[string]string {
	return map_UserKubeConfigStatus
}

var map_WorkspaceRole = map[string]string{
	"": "WorkspaceRole",
}

func (WorkspaceRole) SwaggerDoc() map[string]string {
	return map_WorkspaceRole
}

var map_WorkspaceRoleBinding = map[string]string{
	"":         "WorkspaceRoleBinding",
	"subjects": "Subjects holds references to the objects the role applies to.",
	"roleRef":  "RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.",
}

func (WorkspaceRoleBinding) SwaggerDoc() map[string]string {
	return map_WorkspaceRoleBinding
}

var map_WorkspaceRoleBindingList = map[string]string{
	"": "WorkspaceRoleBindingList contains a list of Workspace",
}

func (WorkspaceRoleBindingList) SwaggerDoc() map[string]string {
	return map_WorkspaceRoleBindingList
}

var map_WorkspaceRoleList = map[string]string{
	"": "WorkspaceRoleList contains a list of Workspace",
}

func (WorkspaceRoleList) SwaggerDoc() map[string]string {
	return map_WorkspaceRoleList
}

// AUTO-GENERATED FUNCTIONS END HERE

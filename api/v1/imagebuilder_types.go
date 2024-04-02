/*
Copyright 2024.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type ImageBuilderSpec struct {
	PodName       string `json:"podName,omitempty"`
	Namespace     string `json:"namespace,omitempty"`
	ContainerName string `json:"containerName,omitempty"`
	Username      string `json:"username,omitempty"`
	Password      string `json:"password,omitempty"`
	To            string `json:"to,omitempty"`
}

type ImageBuilderStatus struct {
	State  string `json:"state,omitempty"`
	Reason string `json:"reason,omitempty"`
	Node   string `json:"node,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// +kubebuilder:printcolumn:name="PodName",type=string,JSONPath=`.spec.podName`
// +kubebuilder:printcolumn:name="PodNamespace",type=string,JSONPath=`.spec.namespace`
// +kubebuilder:printcolumn:name="ContainerName",type=string,JSONPath=`.spec.containerName`
// +kubebuilder:printcolumn:name="To",type=string,JSONPath=`.spec.to`
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.state`
// +kubebuilder:printcolumn:name="Node",type=string,JSONPath=`.status.node`
type ImageBuilder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImageBuilderSpec   `json:"spec,omitempty"`
	Status ImageBuilderStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ImageBuilderList contains a list of ImageBuilder
type ImageBuilderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageBuilder `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ImageBuilder{}, &ImageBuilderList{})
}

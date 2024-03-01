/*
Copyright 2018 The Kubernetes Authors.

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

package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExampleList is a list of Example objects.
type ExampleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Example `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// ExampleSpec is the specification of a Example.
type ExampleSpec struct {
	// An example description.
	Description string `json:"description,omitempty" protobuf:"bytes,1,opt,name=description"`
}

// ExampleStatus is the status of a Example.
type ExampleStatus struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Example is an example type with a spec and a status.
type Example struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   ExampleSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status ExampleStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:conversion-gen:explicit-from=net/url.Values
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ResourceCallOptions struct {
	metav1.TypeMeta `json:",inline"`

	// Path is the URL path
	// +optional
	Path string `json:"path,omitempty" protobuf:"bytes,1,opt,name=path"`
}

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type QueryRequest struct {
	metav1.TypeMeta `json:",inline"`

	Spec QueryRequestSpec `json:"spec,omitempty"`
}

type QueryRequestSpec struct {
	Foo string `json:"foo,omitempty"`
}

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type QueryResponse struct {
	metav1.TypeMeta `json:",inline"`

	Spec QueryResponseSpec `json:"spec,omitempty"`
}

type QueryResponseSpec struct {
	Bar string `json:"bar,omitempty"`
}

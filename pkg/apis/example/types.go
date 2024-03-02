/*
Copyright 2017 The Kubernetes Authors.

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

package example

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExampleList is a list of Example objects.
type ExampleList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []Example
}

// ExampleSpec is the specification of a Example.
type ExampleSpec struct {
	Description string
}

// ExampleStatus is the status of a Example.
type ExampleStatus struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Example is an example type with a spec and a status.
type Example struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   ExampleSpec
	Status ExampleStatus
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ResourceCallOptions struct {
	metav1.TypeMeta

	// Path is the URL path to use for the current proxy request
	Path string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type QueryRequest struct {
	metav1.TypeMeta `json:",inline"`

	Spec QueryRequestSpec `json:"spec,omitempty"`
}

type QueryRequestSpec struct {
	Foo string `json:"foo,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type QueryResponse struct {
	metav1.TypeMeta `json:",inline"`

	Spec QueryResponseSpec `json:"spec,omitempty"`
}

type QueryResponseSpec struct {
	Bar string `json:"bar,omitempty"`
}

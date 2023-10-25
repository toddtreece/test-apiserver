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

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package
// +k8s:conversion-gen=github.com/toddtreece/test-apiserver/pkg/apis/example
// +k8s:defaulter-gen=TypeMeta
// +groupName=example.toddtreece.com

// Package v1beta1 is the v1beta1 version of the API.
package v1beta1 // import "github.com/toddtreece/test-apiserver/pkg/apis/example/v1beta1"
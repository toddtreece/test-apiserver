//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	url "net/url"
	unsafe "unsafe"

	example "github.com/toddtreece/test-apiserver/pkg/apis/example"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Example)(nil), (*example.Example)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Example_To_example_Example(a.(*Example), b.(*example.Example), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.Example)(nil), (*Example)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_Example_To_v1beta1_Example(a.(*example.Example), b.(*Example), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ExampleList)(nil), (*example.ExampleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ExampleList_To_example_ExampleList(a.(*ExampleList), b.(*example.ExampleList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.ExampleList)(nil), (*ExampleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_ExampleList_To_v1beta1_ExampleList(a.(*example.ExampleList), b.(*ExampleList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ExampleSpec)(nil), (*example.ExampleSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ExampleSpec_To_example_ExampleSpec(a.(*ExampleSpec), b.(*example.ExampleSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.ExampleSpec)(nil), (*ExampleSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_ExampleSpec_To_v1beta1_ExampleSpec(a.(*example.ExampleSpec), b.(*ExampleSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ExampleStatus)(nil), (*example.ExampleStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ExampleStatus_To_example_ExampleStatus(a.(*ExampleStatus), b.(*example.ExampleStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.ExampleStatus)(nil), (*ExampleStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_ExampleStatus_To_v1beta1_ExampleStatus(a.(*example.ExampleStatus), b.(*ExampleStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*QueryRequest)(nil), (*example.QueryRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_QueryRequest_To_example_QueryRequest(a.(*QueryRequest), b.(*example.QueryRequest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.QueryRequest)(nil), (*QueryRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_QueryRequest_To_v1beta1_QueryRequest(a.(*example.QueryRequest), b.(*QueryRequest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*QueryRequestSpec)(nil), (*example.QueryRequestSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_QueryRequestSpec_To_example_QueryRequestSpec(a.(*QueryRequestSpec), b.(*example.QueryRequestSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.QueryRequestSpec)(nil), (*QueryRequestSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_QueryRequestSpec_To_v1beta1_QueryRequestSpec(a.(*example.QueryRequestSpec), b.(*QueryRequestSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*QueryResponse)(nil), (*example.QueryResponse)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_QueryResponse_To_example_QueryResponse(a.(*QueryResponse), b.(*example.QueryResponse), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.QueryResponse)(nil), (*QueryResponse)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_QueryResponse_To_v1beta1_QueryResponse(a.(*example.QueryResponse), b.(*QueryResponse), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*QueryResponseSpec)(nil), (*example.QueryResponseSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_QueryResponseSpec_To_example_QueryResponseSpec(a.(*QueryResponseSpec), b.(*example.QueryResponseSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.QueryResponseSpec)(nil), (*QueryResponseSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_QueryResponseSpec_To_v1beta1_QueryResponseSpec(a.(*example.QueryResponseSpec), b.(*QueryResponseSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ResourceCallOptions)(nil), (*example.ResourceCallOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ResourceCallOptions_To_example_ResourceCallOptions(a.(*ResourceCallOptions), b.(*example.ResourceCallOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.ResourceCallOptions)(nil), (*ResourceCallOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_ResourceCallOptions_To_v1beta1_ResourceCallOptions(a.(*example.ResourceCallOptions), b.(*ResourceCallOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*url.Values)(nil), (*ResourceCallOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_url_Values_To_v1beta1_ResourceCallOptions(a.(*url.Values), b.(*ResourceCallOptions), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_Example_To_example_Example(in *Example, out *example.Example, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_ExampleSpec_To_example_ExampleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_ExampleStatus_To_example_ExampleStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_Example_To_example_Example is an autogenerated conversion function.
func Convert_v1beta1_Example_To_example_Example(in *Example, out *example.Example, s conversion.Scope) error {
	return autoConvert_v1beta1_Example_To_example_Example(in, out, s)
}

func autoConvert_example_Example_To_v1beta1_Example(in *example.Example, out *Example, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_example_ExampleSpec_To_v1beta1_ExampleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_example_ExampleStatus_To_v1beta1_ExampleStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_example_Example_To_v1beta1_Example is an autogenerated conversion function.
func Convert_example_Example_To_v1beta1_Example(in *example.Example, out *Example, s conversion.Scope) error {
	return autoConvert_example_Example_To_v1beta1_Example(in, out, s)
}

func autoConvert_v1beta1_ExampleList_To_example_ExampleList(in *ExampleList, out *example.ExampleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]example.Example)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_ExampleList_To_example_ExampleList is an autogenerated conversion function.
func Convert_v1beta1_ExampleList_To_example_ExampleList(in *ExampleList, out *example.ExampleList, s conversion.Scope) error {
	return autoConvert_v1beta1_ExampleList_To_example_ExampleList(in, out, s)
}

func autoConvert_example_ExampleList_To_v1beta1_ExampleList(in *example.ExampleList, out *ExampleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Example)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_example_ExampleList_To_v1beta1_ExampleList is an autogenerated conversion function.
func Convert_example_ExampleList_To_v1beta1_ExampleList(in *example.ExampleList, out *ExampleList, s conversion.Scope) error {
	return autoConvert_example_ExampleList_To_v1beta1_ExampleList(in, out, s)
}

func autoConvert_v1beta1_ExampleSpec_To_example_ExampleSpec(in *ExampleSpec, out *example.ExampleSpec, s conversion.Scope) error {
	out.Description = in.Description
	return nil
}

// Convert_v1beta1_ExampleSpec_To_example_ExampleSpec is an autogenerated conversion function.
func Convert_v1beta1_ExampleSpec_To_example_ExampleSpec(in *ExampleSpec, out *example.ExampleSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_ExampleSpec_To_example_ExampleSpec(in, out, s)
}

func autoConvert_example_ExampleSpec_To_v1beta1_ExampleSpec(in *example.ExampleSpec, out *ExampleSpec, s conversion.Scope) error {
	out.Description = in.Description
	return nil
}

// Convert_example_ExampleSpec_To_v1beta1_ExampleSpec is an autogenerated conversion function.
func Convert_example_ExampleSpec_To_v1beta1_ExampleSpec(in *example.ExampleSpec, out *ExampleSpec, s conversion.Scope) error {
	return autoConvert_example_ExampleSpec_To_v1beta1_ExampleSpec(in, out, s)
}

func autoConvert_v1beta1_ExampleStatus_To_example_ExampleStatus(in *ExampleStatus, out *example.ExampleStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1beta1_ExampleStatus_To_example_ExampleStatus is an autogenerated conversion function.
func Convert_v1beta1_ExampleStatus_To_example_ExampleStatus(in *ExampleStatus, out *example.ExampleStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_ExampleStatus_To_example_ExampleStatus(in, out, s)
}

func autoConvert_example_ExampleStatus_To_v1beta1_ExampleStatus(in *example.ExampleStatus, out *ExampleStatus, s conversion.Scope) error {
	return nil
}

// Convert_example_ExampleStatus_To_v1beta1_ExampleStatus is an autogenerated conversion function.
func Convert_example_ExampleStatus_To_v1beta1_ExampleStatus(in *example.ExampleStatus, out *ExampleStatus, s conversion.Scope) error {
	return autoConvert_example_ExampleStatus_To_v1beta1_ExampleStatus(in, out, s)
}

func autoConvert_v1beta1_QueryRequest_To_example_QueryRequest(in *QueryRequest, out *example.QueryRequest, s conversion.Scope) error {
	if err := Convert_v1beta1_QueryRequestSpec_To_example_QueryRequestSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_QueryRequest_To_example_QueryRequest is an autogenerated conversion function.
func Convert_v1beta1_QueryRequest_To_example_QueryRequest(in *QueryRequest, out *example.QueryRequest, s conversion.Scope) error {
	return autoConvert_v1beta1_QueryRequest_To_example_QueryRequest(in, out, s)
}

func autoConvert_example_QueryRequest_To_v1beta1_QueryRequest(in *example.QueryRequest, out *QueryRequest, s conversion.Scope) error {
	if err := Convert_example_QueryRequestSpec_To_v1beta1_QueryRequestSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_example_QueryRequest_To_v1beta1_QueryRequest is an autogenerated conversion function.
func Convert_example_QueryRequest_To_v1beta1_QueryRequest(in *example.QueryRequest, out *QueryRequest, s conversion.Scope) error {
	return autoConvert_example_QueryRequest_To_v1beta1_QueryRequest(in, out, s)
}

func autoConvert_v1beta1_QueryRequestSpec_To_example_QueryRequestSpec(in *QueryRequestSpec, out *example.QueryRequestSpec, s conversion.Scope) error {
	out.Foo = in.Foo
	return nil
}

// Convert_v1beta1_QueryRequestSpec_To_example_QueryRequestSpec is an autogenerated conversion function.
func Convert_v1beta1_QueryRequestSpec_To_example_QueryRequestSpec(in *QueryRequestSpec, out *example.QueryRequestSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_QueryRequestSpec_To_example_QueryRequestSpec(in, out, s)
}

func autoConvert_example_QueryRequestSpec_To_v1beta1_QueryRequestSpec(in *example.QueryRequestSpec, out *QueryRequestSpec, s conversion.Scope) error {
	out.Foo = in.Foo
	return nil
}

// Convert_example_QueryRequestSpec_To_v1beta1_QueryRequestSpec is an autogenerated conversion function.
func Convert_example_QueryRequestSpec_To_v1beta1_QueryRequestSpec(in *example.QueryRequestSpec, out *QueryRequestSpec, s conversion.Scope) error {
	return autoConvert_example_QueryRequestSpec_To_v1beta1_QueryRequestSpec(in, out, s)
}

func autoConvert_v1beta1_QueryResponse_To_example_QueryResponse(in *QueryResponse, out *example.QueryResponse, s conversion.Scope) error {
	if err := Convert_v1beta1_QueryResponseSpec_To_example_QueryResponseSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_QueryResponse_To_example_QueryResponse is an autogenerated conversion function.
func Convert_v1beta1_QueryResponse_To_example_QueryResponse(in *QueryResponse, out *example.QueryResponse, s conversion.Scope) error {
	return autoConvert_v1beta1_QueryResponse_To_example_QueryResponse(in, out, s)
}

func autoConvert_example_QueryResponse_To_v1beta1_QueryResponse(in *example.QueryResponse, out *QueryResponse, s conversion.Scope) error {
	if err := Convert_example_QueryResponseSpec_To_v1beta1_QueryResponseSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_example_QueryResponse_To_v1beta1_QueryResponse is an autogenerated conversion function.
func Convert_example_QueryResponse_To_v1beta1_QueryResponse(in *example.QueryResponse, out *QueryResponse, s conversion.Scope) error {
	return autoConvert_example_QueryResponse_To_v1beta1_QueryResponse(in, out, s)
}

func autoConvert_v1beta1_QueryResponseSpec_To_example_QueryResponseSpec(in *QueryResponseSpec, out *example.QueryResponseSpec, s conversion.Scope) error {
	out.Bar = in.Bar
	return nil
}

// Convert_v1beta1_QueryResponseSpec_To_example_QueryResponseSpec is an autogenerated conversion function.
func Convert_v1beta1_QueryResponseSpec_To_example_QueryResponseSpec(in *QueryResponseSpec, out *example.QueryResponseSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_QueryResponseSpec_To_example_QueryResponseSpec(in, out, s)
}

func autoConvert_example_QueryResponseSpec_To_v1beta1_QueryResponseSpec(in *example.QueryResponseSpec, out *QueryResponseSpec, s conversion.Scope) error {
	out.Bar = in.Bar
	return nil
}

// Convert_example_QueryResponseSpec_To_v1beta1_QueryResponseSpec is an autogenerated conversion function.
func Convert_example_QueryResponseSpec_To_v1beta1_QueryResponseSpec(in *example.QueryResponseSpec, out *QueryResponseSpec, s conversion.Scope) error {
	return autoConvert_example_QueryResponseSpec_To_v1beta1_QueryResponseSpec(in, out, s)
}

func autoConvert_v1beta1_ResourceCallOptions_To_example_ResourceCallOptions(in *ResourceCallOptions, out *example.ResourceCallOptions, s conversion.Scope) error {
	out.Path = in.Path
	return nil
}

// Convert_v1beta1_ResourceCallOptions_To_example_ResourceCallOptions is an autogenerated conversion function.
func Convert_v1beta1_ResourceCallOptions_To_example_ResourceCallOptions(in *ResourceCallOptions, out *example.ResourceCallOptions, s conversion.Scope) error {
	return autoConvert_v1beta1_ResourceCallOptions_To_example_ResourceCallOptions(in, out, s)
}

func autoConvert_example_ResourceCallOptions_To_v1beta1_ResourceCallOptions(in *example.ResourceCallOptions, out *ResourceCallOptions, s conversion.Scope) error {
	out.Path = in.Path
	return nil
}

// Convert_example_ResourceCallOptions_To_v1beta1_ResourceCallOptions is an autogenerated conversion function.
func Convert_example_ResourceCallOptions_To_v1beta1_ResourceCallOptions(in *example.ResourceCallOptions, out *ResourceCallOptions, s conversion.Scope) error {
	return autoConvert_example_ResourceCallOptions_To_v1beta1_ResourceCallOptions(in, out, s)
}

func autoConvert_url_Values_To_v1beta1_ResourceCallOptions(in *url.Values, out *ResourceCallOptions, s conversion.Scope) error {
	// WARNING: Field TypeMeta does not have json tag, skipping.

	if values, ok := map[string][]string(*in)["path"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_string(&values, &out.Path, s); err != nil {
			return err
		}
	} else {
		out.Path = ""
	}
	return nil
}

// Convert_url_Values_To_v1beta1_ResourceCallOptions is an autogenerated conversion function.
func Convert_url_Values_To_v1beta1_ResourceCallOptions(in *url.Values, out *ResourceCallOptions, s conversion.Scope) error {
	return autoConvert_url_Values_To_v1beta1_ResourceCallOptions(in, out, s)
}

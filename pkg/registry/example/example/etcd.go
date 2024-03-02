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

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"

	"github.com/toddtreece/test-apiserver/pkg/apis/example"
)

type Storage struct {
	Example      *REST
	ResourceCall *ResourceCallREST
	Query        *QueryREST
}

type REST struct {
	*Wrapper
}

// NewREST returns a RESTStorage object that will work against API services.
func NewREST(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (*Storage, error) {
	strategy := NewStrategy(scheme)

	store := &genericregistry.Store{
		NewFunc:                   func() runtime.Object { return &example.Example{} },
		NewListFunc:               func() runtime.Object { return &example.ExampleList{} },
		PredicateFunc:             MatchExample,
		DefaultQualifiedResource:  example.Resource("examples"),
		SingularQualifiedResource: example.Resource("example"),

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,

		// TODO: define table converter that exposes more than name/creation timestamp
		TableConvertor: rest.NewDefaultTableConvertor(example.Resource("examples")),
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	exampleStorage := &REST{&Wrapper{store}}
	resourceCall := &ResourceCallREST{store}
	query := &QueryREST{store}
	return &Storage{
		Example:      exampleStorage,
		ResourceCall: resourceCall,
		Query:        query,
	}, nil

}

var _ rest.TableConvertor = (*Wrapper)(nil)
var _ rest.SingularNameProvider = (*Wrapper)(nil)
var _ rest.Storage = (*Wrapper)(nil)
var _ rest.Creater = (*Wrapper)(nil)
var _ rest.Getter = (*Wrapper)(nil)
var _ rest.Lister = (*Wrapper)(nil)

type Wrapper struct {
	r *genericregistry.Store
}

func (w *Wrapper) New() runtime.Object {
	return w.r.New()
}

func (w *Wrapper) NewList() runtime.Object {
	return w.r.NewList()
}

func (w *Wrapper) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	return w.r.ConvertToTable(ctx, object, tableOptions)
}

func (w *Wrapper) Destroy() {
	w.r.Destroy()
}

func (w *Wrapper) NamespaceScoped() bool {
	return w.r.NamespaceScoped()
}

func (w *Wrapper) GetSingularName() string {
	return w.r.GetSingularName()
}

func (w *Wrapper) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return w.r.Get(ctx, name, options)
}

func (w *Wrapper) List(ctx context.Context, options *internalversion.ListOptions) (runtime.Object, error) {
	return w.r.List(ctx, options)
}

func (w *Wrapper) Create(ctx context.Context, obj runtime.Object, validate rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	return w.r.Create(ctx, obj, validate, options)
}

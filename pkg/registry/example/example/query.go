package example

import (
	"context"

	examplev1beta1 "github.com/toddtreece/test-apiserver/pkg/apis/example/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
)

type QueryREST struct {
	Store *genericregistry.Store
}

var _ rest.Storage = (*QueryREST)(nil)
var _ rest.Creater = (*QueryREST)(nil)
var _ rest.StorageMetadata = (*QueryREST)(nil)

func (r *QueryREST) ProducesMIMETypes(verb string) []string {
	return nil
}

func (r *QueryREST) ProducesObject(verb string) interface{} {
	return &examplev1beta1.QueryResponse{}
}

func (r *QueryREST) New() runtime.Object {
	return &examplev1beta1.QueryRequest{}
}

func (r *QueryREST) Destroy() {}

func (r *QueryREST) Create(ctx context.Context, obj runtime.Object, validator rest.ValidateObjectFunc, _ *metav1.CreateOptions) (runtime.Object, error) {
	return &examplev1beta1.QueryResponse{
		Spec: examplev1beta1.QueryResponseSpec{
			Bar: "response",
		},
	}, nil
}

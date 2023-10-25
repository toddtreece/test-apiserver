package example

import (
	"context"
	"fmt"
	"net/http"

	"github.com/toddtreece/test-apiserver/pkg/apis/example"
	"k8s.io/apimachinery/pkg/runtime"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
)

type ResourceCallREST struct {
	Store *genericregistry.Store
}

var _ = rest.Connecter(&ResourceCallREST{})

var methods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}

func (r *ResourceCallREST) New() runtime.Object {
	return &example.ResourceCallOptions{}
}

func (r *ResourceCallREST) Destroy() {
}

func (r *ResourceCallREST) ConnectMethods() []string {
	return methods
}

func (r *ResourceCallREST) NewConnectOptions() (runtime.Object, bool, string) {
	return &example.ResourceCallOptions{}, true, "path"
}

func (r *ResourceCallREST) Connect(ctx context.Context, id string, opts runtime.Object, responder rest.Responder) (http.Handler, error) {
	resourceCallOpts, ok := opts.(*example.ResourceCallOptions)
	if !ok {
		return nil, fmt.Errorf("invalid options object: %#v", opts)
	}
	fmt.Println("ResourceCallREST.Connect() called with id:", id, "and opts:", resourceCallOpts)
	return &fakeHandler{}, nil
}

type fakeHandler struct {
}

func (f *fakeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Path:\n" + req.URL.Path + "\n"))
}

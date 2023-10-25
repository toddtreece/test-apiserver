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

// ResourceCallREST implements the proxy subresource for a Pod
type ResourceCallREST struct {
	Store *genericregistry.Store
}

// Implement Connecter
var _ = rest.Connecter(&ResourceCallREST{})

var proxyMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}

// New returns an empty podProxyOptions object.
func (r *ResourceCallREST) New() runtime.Object {
	return &example.ResourceCallOptions{}
}

// Destroy cleans up resources on shutdown.
func (r *ResourceCallREST) Destroy() {
	// Given that underlying store is shared with REST,
	// we don't destroy it here explicitly.
}

// ConnectMethods returns the list of HTTP methods that can be proxied
func (r *ResourceCallREST) ConnectMethods() []string {
	return proxyMethods
}

// NewConnectOptions returns versioned resource that represents proxy parameters
func (r *ResourceCallREST) NewConnectOptions() (runtime.Object, bool, string) {
	return &example.ResourceCallOptions{}, true, "path"
}

// Connect returns a handler for the pod proxy
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
	w.Write([]byte(req.URL.Path))
}

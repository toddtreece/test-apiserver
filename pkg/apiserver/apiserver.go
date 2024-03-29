/*
Copyright 2016 The Kubernetes Authors.

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

package apiserver

import (
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/apiserver/pkg/endpoints/discovery"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"

	"github.com/toddtreece/test-apiserver/pkg/apis/example"
	"github.com/toddtreece/test-apiserver/pkg/apis/example/install"
	examplestorage "github.com/toddtreece/test-apiserver/pkg/registry/example/example"
)

var (
	// Scheme defines methods for serializing and deserializing API objects.
	Scheme = runtime.NewScheme()
	// Codecs provides methods for retrieving codecs and serializers for specific
	// versions and content types.
	Codecs = serializer.NewCodecFactory(Scheme)
)

func init() {
	install.Install(Scheme)

	// we need to add the options to empty v1
	// TODO fix the server code to avoid this
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})

	// TODO: keep the generic API server from wanting this
	unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	Scheme.AddUnversionedTypes(unversioned,
		&metav1.Status{},
		&metav1.APIVersions{},
		&metav1.APIGroupList{},
		&metav1.APIGroup{},
		&metav1.APIResourceList{},
	)
}

// ExtraConfig holds custom apiserver config
type ExtraConfig struct {
	// Place you custom config here.
}

// Config defines the config for the apiserver
type Config struct {
	GenericConfig *genericapiserver.RecommendedConfig
	ExtraConfig   ExtraConfig
}

// ExampleServer contains state for a Kubernetes cluster master/api server.
type ExampleServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}

type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   *ExtraConfig
}

// CompletedConfig embeds a private pointer that cannot be instantiated outside of this package.
type CompletedConfig struct {
	*completedConfig
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (cfg *Config) Complete() CompletedConfig {
	c := completedConfig{
		cfg.GenericConfig.Complete(),
		&cfg.ExtraConfig,
	}

	c.GenericConfig.Version = &version.Info{
		Major: "1",
		Minor: "0",
	}

	return CompletedConfig{&c}
}

// New returns a new instance of ExampleServer from the given config.
func (c completedConfig) New() (*ExampleServer, error) {
	delegationTarget := genericapiserver.NewEmptyDelegate()
	genericServer, err := c.GenericConfig.New("test-apiserver", delegationTarget)
	if err != nil {
		return nil, err
	}

	s := &ExampleServer{
		GenericAPIServer: genericServer,
	}

	parameterCodec := runtime.NewParameterCodec(Scheme)

	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(example.GroupName, Scheme, parameterCodec, Codecs)

	v1beta1storage := map[string]rest.Storage{}
	restStorage, err := examplestorage.NewREST(Scheme, c.GenericConfig.RESTOptionsGetter)
	if err != nil {
		return nil, err
	}
	v1beta1storage["examples"] = restStorage.Example
	v1beta1storage["examples/resources"] = restStorage.ResourceCall
	v1beta1storage["examples/query"] = restStorage.Query
	apiGroupInfo.VersionedResourcesStorageMap["v1beta1"] = v1beta1storage

	if err := s.GenericAPIServer.InstallAPIGroup(&apiGroupInfo); err != nil {
		return nil, err
	}

	delegateHandler := delegationTarget.UnprotectedHandler()
	if delegateHandler == nil {
		delegateHandler = http.NotFoundHandler()
	}

	versionDiscovery := &versionDiscoveryHandler{
		discovery: map[schema.GroupVersion]*discovery.APIVersionHandler{},
		delegate:  delegateHandler,
	}

	groupDiscovery := &groupDiscoveryHandler{
		discovery: map[string]*discovery.APIGroupHandler{},
		delegate:  delegateHandler,
	}

	gv := schema.GroupVersion{Group: "empty.toddtreece.com", Version: "v1"}
	apiVersionsForDiscovery := []metav1.GroupVersionForDiscovery{}
	apiResourcesForDiscovery := []metav1.APIResource{}
	apiVersionsForDiscovery = append(apiVersionsForDiscovery, metav1.GroupVersionForDiscovery{
		GroupVersion: gv.String(),
		Version:      gv.Version,
	})
	apiGroup := metav1.APIGroup{
		Name:     gv.Group,
		Versions: apiVersionsForDiscovery,
		// the preferred versions for a group is the first item in
		// apiVersionsForDiscovery after it put in the right ordered
		PreferredVersion: apiVersionsForDiscovery[0],
	}
	groupDiscovery.setDiscovery(gv.Group, discovery.NewAPIGroupHandler(Codecs, apiGroup))
	versionDiscovery.setDiscovery(gv, discovery.NewAPIVersionHandler(Codecs, gv, discovery.APIResourceListerFunc(func() []metav1.APIResource {
		return apiResourcesForDiscovery
	})))

	s.GenericAPIServer.Handler.NonGoRestfulMux.Handle("/apis/empty.toddtreece.com", groupDiscovery)
	s.GenericAPIServer.Handler.NonGoRestfulMux.Handle("/apis/empty.toddtreece.com/v1", versionDiscovery)

	return s, nil
}

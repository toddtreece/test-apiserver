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

package server

import (
	"fmt"
	"io"
	"net"

	"github.com/spf13/cobra"
	"github.com/toddtreece/test-apiserver/pkg/apis/example/v1beta1"
	"github.com/toddtreece/test-apiserver/pkg/apiserver"

	exampleopenapi "github.com/toddtreece/test-apiserver/pkg/generated/openapi"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apiserver/pkg/endpoints/openapi"
	"k8s.io/apiserver/pkg/features"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/client-go/informers"
	netutils "k8s.io/utils/net"
)

const defaultEtcdPathPrefix = "/registry/example.toddtreece.com"

// ExampleServerOptions contains state for master/api server
type ExampleServerOptions struct {
	RecommendedOptions *genericoptions.RecommendedOptions

	SharedInformerFactory informers.SharedInformerFactory
	StdOut                io.Writer
	StdErr                io.Writer

	AlternateDNS []string
}

// NewExampleServerOptions returns a new ExampleServerOptions
func NewExampleServerOptions(out, errOut io.Writer) *ExampleServerOptions {
	o := &ExampleServerOptions{
		RecommendedOptions: genericoptions.NewRecommendedOptions(
			defaultEtcdPathPrefix,
			apiserver.Codecs.LegacyCodec(v1beta1.SchemeGroupVersion),
		),

		StdOut: out,
		StdErr: errOut,
	}
	o.RecommendedOptions.Etcd.StorageConfig.EncodeVersioner = runtime.NewMultiGroupVersioner(v1beta1.SchemeGroupVersion, schema.GroupKind{Group: v1beta1.GroupName})
	return o
}

// NewCommandStartExampleServer provides a CLI handler for 'start master' command
// with a default ExampleServerOptions.
func NewCommandStartExampleServer(defaults *ExampleServerOptions, stopCh <-chan struct{}) *cobra.Command {
	o := *defaults
	cmd := &cobra.Command{
		Short: "Launch a test API server",
		Long:  "Launch a test API server",
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(); err != nil {
				return err
			}
			if err := o.Validate(args); err != nil {
				return err
			}
			if err := o.RunExampleServer(stopCh); err != nil {
				return err
			}
			return nil
		},
	}

	flags := cmd.Flags()
	o.RecommendedOptions.AddFlags(flags)
	utilfeature.DefaultMutableFeatureGate.AddFlag(flags)

	return cmd
}

// Validate validates ExampleServerOptions
func (o ExampleServerOptions) Validate(args []string) error {
	errors := []error{}
	errors = append(errors, o.RecommendedOptions.Validate()...)
	return utilerrors.NewAggregate(errors)
}

// Complete fills in fields required to have valid data
func (o *ExampleServerOptions) Complete() error {
	return nil
}

// Config returns config for the api server given ExampleServerOptions
func (o *ExampleServerOptions) Config() (*apiserver.Config, error) {
	// TODO have a "real" external address
	if err := o.RecommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", o.AlternateDNS, []net.IP{netutils.ParseIPSloppy("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	o.RecommendedOptions.Etcd.StorageConfig.Paging = true

	serverConfig := genericapiserver.NewRecommendedConfig(apiserver.Codecs)

	serverConfig.OpenAPIConfig = genericapiserver.DefaultOpenAPIConfig(exampleopenapi.GetOpenAPIDefinitions, openapi.NewDefinitionNamer(apiserver.Scheme))
	serverConfig.OpenAPIConfig.Info.Title = "Example"
	serverConfig.OpenAPIConfig.Info.Version = "0.1"

	if utilfeature.DefaultFeatureGate.Enabled(features.OpenAPIV3) {
		serverConfig.OpenAPIV3Config = genericapiserver.DefaultOpenAPIV3Config(exampleopenapi.GetOpenAPIDefinitions, openapi.NewDefinitionNamer(apiserver.Scheme))
		serverConfig.OpenAPIV3Config.Info.Title = "Example"
		serverConfig.OpenAPIV3Config.Info.Version = "0.1"
	}

	if err := o.RecommendedOptions.ApplyTo(serverConfig); err != nil {
		return nil, err
	}

	config := &apiserver.Config{
		GenericConfig: serverConfig,
		ExtraConfig:   apiserver.ExtraConfig{},
	}
	return config, nil
}

// RunExampleServer starts a new ExampleServer given ExampleServerOptions
func (o ExampleServerOptions) RunExampleServer(stopCh <-chan struct{}) error {
	config, err := o.Config()
	if err != nil {
		return err
	}

	server, err := config.Complete().New()
	if err != nil {
		return err
	}

	return server.GenericAPIServer.PrepareRun().Run(stopCh)
}

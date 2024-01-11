package apiserver

import (
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/endpoints/discovery"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/kube-openapi/pkg/openapiconv"
	"k8s.io/kube-openapi/pkg/spec3"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func addQueryEndpoint(s *genericapiserver.GenericAPIServer) error {
	if err := addDiscovery(s); err != nil {
		return err
	}

	return nil
}

func addQueryOpenAPI(s *genericapiserver.GenericAPIServer) {
	orig := openapiconv.ConvertV2ToV3(s.StaticOpenAPISpec)
	gv := schema.GroupVersion{Group: "query.toddtreece.com", Version: "v1"}

	querySchema := spec.Schema{
		SchemaProps: spec.SchemaProps{
			Description: "Copied from: https://github.com/grafana/grafana/blob/main/pkg/api/dtos/models.go#L62",
			Type:        []string{"object"},
			Properties: map[string]spec.Schema{
				"from": {
					SchemaProps: spec.SchemaProps{
						Description: "From Start time in epoch timestamps in milliseconds or relative using Grafana time units. required: true example: now-1h",
						Default:     "",
						Type:        []string{"string"},
						Format:      "",
					},
				},
				"to": {
					SchemaProps: spec.SchemaProps{
						Description: "To End time in epoch timestamps in milliseconds or relative using Grafana time units. required: true example: now",
						Default:     "",
						Type:        []string{"string"},
						Format:      "",
					},
				},
				"debug": {
					SchemaProps: spec.SchemaProps{
						Description: "required: false",
						Type:        []string{"boolean"},
						Format:      "",
					},
				},
			},
			Required: []string{"from", "to", "queries"},
		},
	}

	copy := spec3.OpenAPI{
		Version: orig.Info.Version,
		Info:    &spec.Info{InfoProps: spec.InfoProps{Title: gv.String(), Version: "1.0.0"}},
		Paths: &spec3.Paths{
			Paths: map[string]*spec3.Path{
				"/apis/query.toddtreece.com/v1/query": &spec3.Path{
					PathProps: spec3.PathProps{
						Post: &spec3.Operation{
							OperationProps: spec3.OperationProps{
								Tags:        []string{"query"},
								Description: "query across multiple datasources with expressions.  This api matches the legacy /ds/query endpoint",
								Parameters: []*spec3.Parameter{
									{
										ParameterProps: spec3.ParameterProps{
											Name:        "namespace",
											Description: "object name and auth scope, such as for teams and projects",
											In:          "path",
											Required:    true,
											Schema:      spec.StringProperty(),
											Example:     "default",
										},
									},
								},
								RequestBody: &spec3.RequestBody{
									RequestBodyProps: spec3.RequestBodyProps{
										Required:    true,
										Description: "the query array",
										Content: map[string]*spec3.MediaType{
											"application/json": {
												MediaTypeProps: spec3.MediaTypeProps{
													Schema: &querySchema,
												},
											},
										},
									},
								},
								Responses: &spec3.Responses{
									ResponsesProps: spec3.ResponsesProps{
										StatusCodeResponses: map[int]*spec3.Response{
											200: {
												ResponseProps: spec3.ResponseProps{
													Description: "the query results",
													Content: map[string]*spec3.MediaType{
														"application/json": {
															MediaTypeProps: spec3.MediaTypeProps{
																Schema: spec.MapProperty(spec.StringProperty()),
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			VendorExtensible: spec.VendorExtensible{},
		},
		Servers:             orig.Servers,
		Components:          orig.Components,
		SecurityRequirement: []map[string][]string{},
		ExternalDocs:        orig.ExternalDocs,
	}

	s.OpenAPIV3VersionedService.UpdateGroupVersion("apis/"+gv.String(), &copy)
}

func addDiscovery(s *genericapiserver.GenericAPIServer) error {
	delegateHandler := genericapiserver.NewEmptyDelegate().UnprotectedHandler()
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

	gv := schema.GroupVersion{Group: "query.toddtreece.com", Version: "v1"}
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

	s.Handler.NonGoRestfulMux.Handle("/apis/query.toddtreece.com", groupDiscovery)
	s.Handler.NonGoRestfulMux.Handle("/apis/query.toddtreece.com/v1", versionDiscovery)

	return nil
}

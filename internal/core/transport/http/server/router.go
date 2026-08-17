package corre_http_server

import (
	"fmt"
	"net/http"
)

type ApiVersion string

var (
	ApiVersionV1 = ApiVersion("v1")
	ApiVersionV2 = ApiVersion("v2")
	ApiVersionV3 = ApiVersion("v3")
)

type APIVersionRouter struct {
	*http.ServeMux
	apiVersion ApiVersion
}

func NewAPIVersionRouter(apiVersion ApiVersion) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux:   http.NewServeMux(),
		apiVersion: apiVersion,
	}
}

func (r *APIVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		r.Handle(pattern, route.Handler)
	}
}

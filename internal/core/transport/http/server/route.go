package corre_http_server

import (
	"net/http"
	core_http_middleware "todolist/internal/core/transport/http/middleware"
)

type Route struct {
	Method     string
	Path       string
	Handler    http.HandlerFunc
	Middleware []core_http_middleware.Middleware
}

func (r *Route) WithMiddleware() http.Handler {
	return core_http_middleware.ChainMiddlewares(
		r.Handler,
		r.Middleware...,
	)
}

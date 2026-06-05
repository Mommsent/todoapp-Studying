package core_http_server

import (
	"net/http"

	core_http_middleware "github.com/Mommsent/todoapp-Studying.git/internal/core/transport/http/middleware"
)

type Route struct {
	Method     string
	Path       string
	Handler    http.HandlerFunc
	Middleware []core_http_middleware.Middleware
}

func (router *Route) WithMiddleware() http.Handler {
	return core_http_middleware.ChainMiddleware(
		router.Handler,
		router.Middleware...,
	)
}

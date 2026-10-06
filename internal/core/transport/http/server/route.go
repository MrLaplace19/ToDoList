package core_http_server

import (
	"net/http"

	core_http_middleware "github.com/MrLaplace19/ToDoList/internal/core/transport/http/middelware"
)

type Router struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
	Middleware []core_http_middleware.Middleware
}

func (r *Router) WithMiddleware() http.Handler{
	return core_http_middleware.ChainMiddleware(r.Handler,r.Middleware...)
}

func NewRoute(
	method string,
	path string,
	handler http.HandlerFunc,
) Router {
	return Router{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}

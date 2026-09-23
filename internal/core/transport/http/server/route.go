package core_http_server

import "net/http"

type Router struct{
	Method string
	Path string
	Handler http.HandlerFunc
}

func NewRoute(
	method string,
	path string,
	handler http.HandlerFunc,
) Router{
	return Router{
		Method: method,
		Path: path,
		Handler: handler,
	}
}
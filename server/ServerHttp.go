package server

import (
	"fmt"
	"net/http"
)

type httpHandler interface {
	RouteAble
	Start()
}

type SingUp func(ctx Context)

type HttpServer struct {
	Port int
	Handler
	Root SingUp
}

func (h HttpServer) Route(method, path string, fn SingUp) {
	h.Handler.Route(method, path, fn)
}

func (h HttpServer) Start() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		context := NewContext(writer, request)
		h.Root(context)
	})
	http.ListenAndServe(fmt.Sprintf(":%d", h.Port), nil)
}

var _ httpHandler = HttpServer{}

func NewServerHttp(port int, filters ...FilterBuild) httpHandler {
	//handler := NewHandler()
	handler := NewNodeHandler()
	root := func(ctx Context) {
		handler.ServeHTTP(ctx.Response, ctx.Request)
	}

	for i := len(filters) - 1; i >= 0; i-- {
		f := filters[i]
		root = f(root)
	}

	return HttpServer{
		Port:    port,
		Handler: handler,
		Root:    root,
	}
}

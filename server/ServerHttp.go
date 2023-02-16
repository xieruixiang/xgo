package server

import (
	"fmt"
	"log"
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
}

var _ httpHandler = HttpServer{}

func (h HttpServer) Route(method, path string, fn SingUp) {
	h.Handler.Route(method, path, fn)
}

func (h HttpServer) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", h.Port), h.Handler))
}

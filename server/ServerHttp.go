package server

import (
	"fmt"
	"log"
	"net/http"
)

type httpHandler interface {
	Route(path string)
	Start()
}

type SingUp func(ctx Context)

type HttpServer struct {
	Port int
	HandlerOnMap
}

func main() {
	fmt.Println(22)
}

func (h *HttpServer) Route(method, path string, fn SingUp) {
	key := h.HandlerOnMap.Key(method, path)
	h.HandlerMap[key] = fn
}

func (h *HttpServer) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", h.Port), h.HandlerOnMap))
}

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

type httpServer struct {
	Port int
}

func main() {
	fmt.Println(22)
}

func (h *httpServer) Route(path string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(path, handlerFunc)
}

func (h *httpServer) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", h.Port), nil))
}

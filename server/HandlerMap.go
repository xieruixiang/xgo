package server

import (
	"net/http"
)

type RouteAble interface {
	Route(method, path string, fn SingUp)
}

type Handler interface {
	http.Handler
	RouteAble
}

type HandlerOnMap struct {
	HandlerMap map[string]func(ctx Context)
}

var _ Handler = &HandlerOnMap{}

func (h *HandlerOnMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := h.Key(r.Method, r.URL.Path)
	if m, ok := h.HandlerMap[key]; ok {
		context := NewContext(w, r)
		m(context)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("未匹配到方法"))
	}
}

func (h *HandlerOnMap) Route(method, path string, fn SingUp) {
	key := h.Key(method, path)
	h.HandlerMap[key] = fn
}

func (h *HandlerOnMap) Key(method, pattern string) string {
	return method + "#" + pattern
}

func NewHandler() Handler {
	return &HandlerOnMap{
		HandlerMap: map[string]func(ctx Context){},
	}
}

package server

import "net/http"

type HandlerOnMap struct {
	HandlerMap map[string]func(ctx Context)
}

func (h HandlerOnMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := h.Key(r.Method, r.URL.Path)
	if m, ok := h.HandlerMap[key]; ok {
		context := NewContext(w, r)
		m(context)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("未匹配到方法"))
	}
}

func (h *HandlerOnMap) Key(method, pattern string) string {
	return method + "#" + pattern
}

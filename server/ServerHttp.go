package server

type httpServer interface {
	Route()
	Start()
}

package externalmodule

import (
	"net"
	"net/http"
	"time"
)

type Server struct {
	serveMux   http.ServeMux
	httpServer *http.Server
}

func NewServer(moduleHandlers ...*ModuleHandler) *Server {
	ms := Server{}

	for _, mh := range moduleHandlers {
		ms.serveMux.HandleFunc("/"+mh.path, mh.handleConnection)
	}

	return &ms
}

func (ms *Server) Serve(addrPort string) error {
	l, err := net.Listen("tcp", addrPort)
	if err != nil {
		return err
	}

	ms.httpServer = &http.Server{
		ReadHeaderTimeout: 2 * time.Second,
		Handler:           ms,
	}

	return ms.httpServer.Serve(l)
}

func (ms *Server) Stop() error {
	return ms.httpServer.Close()
}

func (ms *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ms.serveMux.ServeHTTP(w, r)
}

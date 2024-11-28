package externalmodule

import (
	"net"
	"net/http"
	"time"
)

// Server implements a HTTP server containing remote modules.
type Server struct {
	serveMux   http.ServeMux
	httpServer *http.Server
}

// NewServer returns a new Server containing the given moduleHandlers which define the modules operated by the server.
// Each handler associates a module with a path under which the module will be accessible.
func NewServer(moduleHandlers ...*ModuleHandler) *Server {
	ms := Server{}

	for _, mh := range moduleHandlers {
		ms.serveMux.HandleFunc("/"+mh.path, mh.handleConnection)
	}

	return &ms
}

// Serve starts the module server, making it listen to new connections
// at the given address and port (e.g., "0.0.0.0:8080").
// Serve blocks until the Stop method is called. It is therefore expected to call Serve in a separate goroutine.
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

// Stop stops the server, making the call to the Serve method return.
func (ms *Server) Stop() error {
	return ms.httpServer.Close()
}

// ServeHTTP is a wrapper around the HTTP serveMux's method of the same name, so it can be used as a http.ServeMux.
func (ms *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ms.serveMux.ServeHTTP(w, r)
}

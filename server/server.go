package server

import (
	"net/http"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) (*Server, error) {
	// Server Object
	server := &Server{
		listenAddr: listenAddr,
	}
	return server, nil
}

func (s *Server) StartServer() error {
	return http.ListenAndServe(s.listenAddr, GetRoutes())
}

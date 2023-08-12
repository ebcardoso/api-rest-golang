package server

import (
	"net/http"

	"github.com/ebcardoso/api-rest-golang/config"
)

type Server struct {
	listenAddr string
	config     *config.Config
}

func NewServer(listenAddr string, envFile string) (*Server, error) {
	// Server Configurations
	config, err := config.SetConfigs(envFile)
	if err != nil {
		return &Server{}, err
	}

	// Server Object
	server := &Server{
		listenAddr: listenAddr,
		config:     config,
	}
	return server, nil
}

func (s *Server) StartServer() error {
	return http.ListenAndServe(s.listenAddr, GetRoutes(s.config))
}

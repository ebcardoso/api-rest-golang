package server

import (
	"github.com/go-chi/chi/v5"
)

var (
	V1 = "/api/v1/"
)

func GetRoutes() *chi.Mux {

	router := chi.NewRouter()

	return router
}

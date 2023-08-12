package server

import (
	"time"

	"github.com/go-chi/chi/v5"
)

var (
	V1 = "/api/v1/"
)

func GetRoutes(configs *config.Config) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))

	return router
}

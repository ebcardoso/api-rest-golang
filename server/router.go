package server

import (
	"time"

	"github.com/ebcardoso/api-rest-golang/app/handlers"
	"github.com/ebcardoso/api-rest-golang/app/middlewares"
	"github.com/ebcardoso/api-rest-golang/config"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

var (
	V1 = "/api/v1/"
)

func GetRoutes(configs *config.Config) *chi.Mux {
	auth := handlers.NewAuth(configs)

	router := chi.NewRouter()
	protector := middlewares.NewRouterProtector(configs)

	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Group(func(protected chi.Router) {
		protected.Use(protector.Protect)
	})
	router.Post(V1+"auth/signup", auth.Signup)

	return router
}

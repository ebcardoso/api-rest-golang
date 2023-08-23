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
	users := handlers.NewUsers(configs)

	router := chi.NewRouter()
	protector := middlewares.NewRouterProtector(configs)

	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Group(func(protected chi.Router) {
		protected.Use(protector.Protect)
		protected.Post(V1+"auth/check_token", auth.CheckToken)
		protected.Get(V1+"users", users.GetListUsers)
		protected.Get(V1+"users/{id}", users.GetUserByID)
		protected.Patch(V1+"users/{id}", users.UpdateUser)
		protected.Delete(V1+"users/{id}", users.DestroyUser)
		protected.Post(V1+"users/{id}/block", users.Block)
		protected.Post(V1+"users/{id}/unblock", users.Unblock)
	})
	router.Post(V1+"auth/signup", auth.Signup)
	router.Post(V1+"auth/signin", auth.Signin)
	router.Post(V1+"auth/forgot_password_token", auth.ForgotPasswordToken)

	return router
}

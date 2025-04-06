package routes

import (
	"TodoList-Golang-Auth/delivery"
	"TodoList-Golang-Auth/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupUserRoutes(r chi.Router, userHandler *delivery.UserHandler, todoHandler *delivery.TodoHandler) {
	r.Route("/api/users", func(r chi.Router) {
		r.Post("/register", userHandler.RegisterUser)
		r.Post("/login", userHandler.Login)
	})
	r.Post("/todo", todoHandler.CreateTodo)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Go-Chi API"))
	})
	r.Route("/api", func(protected chi.Router) {
		protected.Use(middleware.JWTAuthMiddleware)

		protected.Get("/me", userHandler.Profile)
	})
}

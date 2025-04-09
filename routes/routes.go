package routes

import (
	"TodoList-Golang-Auth/delivery"
	"TodoList-Golang-Auth/middleware"
	"TodoList-Golang-Auth/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupUserRoutes(r chi.Router, userHandler *delivery.UserHandler, todoHandler *delivery.TodoHandler, blacklistRepo repository.BlacklistRepository) {
	r.Post("/register", userHandler.RegisterUser)
	r.Post("/login", userHandler.Login)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Go-Chi API"))
	})

	r.Route("/api", func(protected chi.Router) {
		protected.Use(middleware.JWTAuthMiddleware(blacklistRepo))

		protected.Post("/todo", todoHandler.CreateTodo)
		protected.Get("/me", userHandler.Profile)
		protected.Post("/logout", userHandler.Logout)
	})
}

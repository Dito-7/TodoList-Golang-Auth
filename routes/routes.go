package routes

import (
	"TodoList-Golang-Auth/delivery"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupUserRoutes(r chi.Router, handler *delivery.UserHandler) {
	r.Post("/register", handler.RegisterUser)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Go-Chi API"))
	})
}

package routes

import (
	"TodoList-Golang-Auth/delivery"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupUserRoutes(r chi.Router, userHandler *delivery.UserHandler, todoHandler *delivery.TodoHandler) {
	r.Post("/register", userHandler.RegisterUser)
	r.Post("/todo", todoHandler.CreateTodo)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Go-Chi API"))
	})
}

package delivery

import (
	"TodoList-Golang-Auth/models"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

type TodoHandler struct {
	todoUsecase models.TodoUsecase
}

func NewTodoHandler(todoUsecase models.TodoUsecase) *TodoHandler {
	return &TodoHandler{
		todoUsecase: todoUsecase,
	}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title       string `json:"title"`
		Completed   bool   `json:"completed"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	todo, err := h.todoUsecase.CreateTodo(context.Background(), req.Title, strconv.FormatBool(req.Completed), req.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

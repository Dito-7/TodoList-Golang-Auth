package usecase

import (
	"TodoList-Golang-Auth/models"
	"context"
)

type TodoUsecase interface {
	CreateTodo(ctx context.Context, title, completed, description string) (*models.Todo, error)
}

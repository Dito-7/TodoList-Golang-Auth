package repository

import (
	"TodoList-Golang-Auth/models"
	"context"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *models.Todo) (*models.Todo, error)
}

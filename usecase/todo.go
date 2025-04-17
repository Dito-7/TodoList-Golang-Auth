package usecase

import (
	"TodoList-Golang-Auth/models"
	"context"
)

type todoUsecaseImpl struct {
	todoRepo models.TodoRepository
}

func NewTodoUsecase(todoRepo models.TodoRepository) models.TodoUsecase {
	return &todoUsecaseImpl{todoRepo: todoRepo}
}

func (u *todoUsecaseImpl) CreateTodo(ctx context.Context, completed, title, description string) (*models.Todo, error) {
	todo := &models.Todo{
		Title:       title,
		Completed:   true,
		Description: description,
	}

	return u.todoRepo.CreateTodo(ctx, todo)
}

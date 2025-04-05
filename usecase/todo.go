package usecase

import (
	"TodoList-Golang-Auth/models"
	"TodoList-Golang-Auth/repository"
	"context"
)

type todoUsecaseImpl struct {
	todoRepo repository.TodoRepository
}

func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
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

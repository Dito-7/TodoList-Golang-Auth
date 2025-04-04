package usecase

import (
	"TodoList-Golang-Auth/models"
	"context"
)

type UserUsecase interface {
	RegisterUser(ctx context.Context, email, password string) (*models.User, error)
}

package repository

import (
	"TodoList-Golang-Auth/models"
	"context"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
}

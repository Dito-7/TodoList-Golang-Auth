package repository

import (
	"TodoList-Golang-Auth/models"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
}

package usecase

import (
	"TodoList-Golang-Auth/models"
	"TodoList-Golang-Auth/repository"
	"TodoList-Golang-Auth/utils"
	"context"
)

type userUsecaseImpl struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecaseImpl{userRepo: userRepo}
}

func (u *userUsecaseImpl) RegisterUser(ctx context.Context, email, password string) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    email,
		Password: hashedPassword,
	}

	return u.userRepo.CreateUser(ctx, user)
}

package usecase

import (
	"TodoList-Golang-Auth/models"
	"TodoList-Golang-Auth/utils"
	"context"
	"errors"
)

type userUsecaseImpl struct {
	userRepo models.UserRepository
}

func NewUserUsecase(userRepo models.UserRepository) models.UserUsecase {
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

func (u *userUsecaseImpl) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	user, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPassword(user.Password, password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

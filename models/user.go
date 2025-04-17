package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email    string             `bson:"email" json:"email" validate:"required,email"`
	Password string             `bson:"password" json:"password" validate:"required,min=6"`
}

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
}

type UserUsecase interface {
	RegisterUser(ctx context.Context, email, password string) (*User, error)
	LoginUser(ctx context.Context, email, password string) (*User, error)
}

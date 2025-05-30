package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title,omitempty" json:"title,omitempty"`
	Completed   bool               `bson:"completed,omitempty" json:"completed,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	UserID      primitive.ObjectID `bson:"userId,omitempty" json:"userId,omitempty"`
}

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *Todo) (*Todo, error)
}

type TodoUsecase interface {
	CreateTodo(ctx context.Context, title, completed, description string) (*Todo, error)
}

package mongodb

import (
	"TodoList-Golang-Auth/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type todoRepositoryImpl struct {
	collection *mongo.Collection
}

func NewTodoRepository(db *mongo.Database) *todoRepositoryImpl {
	return &todoRepositoryImpl{
		collection: db.Collection("todos"),
	}
}

func (r *todoRepositoryImpl) CreateTodo(ctx context.Context, todo *models.Todo) (*models.Todo, error) {
	todo.ID = primitive.NewObjectID()
	_, err := r.collection.InsertOne(ctx, todo)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

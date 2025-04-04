package mongodb

import (
	"TodoList-Golang-Auth/models"
	"TodoList-Golang-Auth/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository struct {
	Client mongo.Collection
}

func (r *TodoRepository) CreateTodo(ctx context.Context, todo models.Todo) (string, error) {
	result, err := r.Client.InsertOne(ctx, todo)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *TodoRepository) GetTodoByID(ctx context.Context, id string) (*models.Todo, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var todo models.Todo
	err = r.Client.FindOne(ctx, bson.M{"_id": objectID}).Decode(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *TodoRepository) UpdateTodo(ctx context.Context, id string, updateData bson.M) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.Client.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": updateData})
	return err
}

func (r *TodoRepository) DeleteTodo(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.Client.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

func NewTodo(client mongo.Collection) repository.TodoInterface {
	return TodoRepository{Client: client}
}

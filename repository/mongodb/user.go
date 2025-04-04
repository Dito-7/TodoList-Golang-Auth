package mongodb

import (
	"TodoList-Golang-Auth/models"
	"TodoList-Golang-Auth/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepositoryImpl struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) repository.UserRepository {
	return &userRepositoryImpl{
		collection: db.Collection("users"),
	}
}

func (r *userRepositoryImpl) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	user.ID = primitive.NewObjectID()
	_, err := r.collection.InsertOne(ctx, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

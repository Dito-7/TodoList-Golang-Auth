package mongodb

import (
	"TodoList-Golang-Auth/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepositoryImpl struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) models.UserRepository {
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

func (r *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	filter := bson.M{"email": email}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

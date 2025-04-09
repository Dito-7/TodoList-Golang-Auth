package mongodb

import (
	"TodoList-Golang-Auth/models"
	"TodoList-Golang-Auth/repository"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type blacklistRepoImpl struct {
	collection *mongo.Collection
}

func NewBlacklistRepository(db *mongo.Database) repository.BlacklistRepository {
	return &blacklistRepoImpl{
		collection: db.Collection("blacklist_tokens"),
	}
}

func (r *blacklistRepoImpl) AddToken(ctx context.Context, token string, expiresAt time.Time) error {
	_, err := r.collection.InsertOne(ctx, models.BlacklistedToken{
		Token:     token,
		ExpiresAt: expiresAt,
	})
	return err
}

func (r *blacklistRepoImpl) IsBlacklisted(ctx context.Context, token string) (bool, error) {
	filter := bson.M{"token": token}
	count, err := r.collection.CountDocuments(ctx, filter)
	return count > 0, err
}

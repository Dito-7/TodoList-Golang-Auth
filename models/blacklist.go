package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlacklistedToken struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Token     string             `bson:"token"`
	ExpiresAt time.Time          `bson:"expires_at"`
}

type BlacklistRepository interface {
	AddToken(ctx context.Context, token string, expiresAt time.Time) error
	IsBlacklisted(ctx context.Context, token string) (bool, error)
	EnsureTTLIndex(ctx context.Context) error
}

package repository

import (
	"context"
	"time"
)

type BlacklistRepository interface {
	AddToken(ctx context.Context, token string, expiresAt time.Time) error
	IsBlacklisted(ctx context.Context, token string) (bool, error)
}

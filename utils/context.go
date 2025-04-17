package utils

import "context"

type contextKey string

const userEmailKey contextKey = "userEmail"

func GetUserEmailFromContext(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(userEmailKey).(string)
	return email, ok
}

func SetUserEmailToContext(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, userEmailKey, email)
}

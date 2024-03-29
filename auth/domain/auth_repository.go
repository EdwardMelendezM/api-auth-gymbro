package domain

import "context"

type AuthRepository interface {
	CheckExistenceByUsername(ctx context.Context, username string) int
	VerifyPassword(ctx context.Context, body LoginBody) int
	CheckAccountStatus(ctx context.Context, body LoginBody) bool
}

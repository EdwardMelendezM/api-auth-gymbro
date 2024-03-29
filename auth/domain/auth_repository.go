package domain

import "context"

type AuthRepository interface {
	CheckExistenceByUsername(ctx context.Context, username string) bool
	VerifyPassword(ctx context.Context, body LoginBody) bool
	CheckAccountStatus(ctx context.Context, body LoginBody) bool
}

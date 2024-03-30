package domain

import "context"

type AuthRepository interface {
	CheckExistenceByUsername(ctx context.Context, username string) (countAccount *int, err error)
	VerifyPassword(ctx context.Context, body LoginBody) (countAccount *int, err error)
	CheckAccountStatus(ctx context.Context, body LoginBody) (statusAccount *int, err error)
}

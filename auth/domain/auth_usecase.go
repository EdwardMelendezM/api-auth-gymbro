package domain

import "context"

type AuthUseCase interface {
	Login(ctx context.Context, body LoginBody) (token string, err error)
}

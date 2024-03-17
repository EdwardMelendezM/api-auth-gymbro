package domain

type AuthUseCase interface {
	Login(body LoginBody) *string
}

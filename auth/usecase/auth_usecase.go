package usecase

import (
	"time"

	"github.com/EdwardMelendezM/api-auth/auth/domain"
)

type authUseCase struct {
	authRepository domain.AuthRepository
	contextTimeout time.Duration
}

func NewAuthUseCase(
	ur domain.AuthRepository,
	timeout time.Duration,
) domain.AuthUseCase {
	return &authUseCase{
		authRepository: ur,
		contextTimeout: timeout,
	}
}

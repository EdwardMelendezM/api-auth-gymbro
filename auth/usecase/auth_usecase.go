package usecase

import (
	errorLog "github.com/EdwardMelendezM/info-code-api-shared-v1/error-log"
	"time"

	"github.com/EdwardMelendezM/api-auth/auth/domain"
)

type authUseCase struct {
	authRepository domain.AuthRepository
	contextTimeout time.Duration
	err            *errorLog.CustomError
}

func NewAuthUseCase(
	ur domain.AuthRepository,
	timeout time.Duration,
) domain.AuthUseCase {
	return &authUseCase{
		authRepository: ur,
		contextTimeout: timeout,
		err:            errorLog.NewErr().SetLayer(errorLog.UseCase),
	}
}

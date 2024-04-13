package usecase

import (
	"time"

	authJwtRepo "github.com/EdwardMelendezM/api-info-shared/auth/domain"
	errorLog "github.com/EdwardMelendezM/api-info-shared/error-log"

	"github.com/EdwardMelendezM/api-auth/authentication/domain"
)

type authUseCase struct {
	authRepository domain.AuthRepository
	authJwtRepo    authJwtRepo.AuthRepository
	contextTimeout time.Duration
	err            *errorLog.CustomError
}

func NewAuthUseCase(
	ur domain.AuthRepository,
	timeout time.Duration,
	authJwtRepo authJwtRepo.AuthRepository,

) domain.AuthUseCase {
	return &authUseCase{
		authRepository: ur,
		contextTimeout: timeout,
		err:            errorLog.NewErr().SetLayer(errorLog.UseCase),
		authJwtRepo:    authJwtRepo,
	}
}

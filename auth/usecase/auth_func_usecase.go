package usecase

import (
	"context"

	"github.com/EdwardMelendezM/api-auth/auth/domain"

	"github.com/EdwardMelendezM/info-code-api-shared-v1/error-log"
)

func (u authUseCase) Login(
	ctx context.Context,
	body domain.LoginBody,
) (token string, err error) {
	defer errorLog.PanicRecovery(ctx, &err)

	var checkExistenceByUsername, errCheckExistenceByUsername = u.authRepository.CheckExistenceByUsername(ctx, body.Username)
	if errCheckExistenceByUsername != nil {
		return "", errCheckExistenceByUsername
	}
	if checkExistenceByUsername != 1 {
		return "", err
	}
	return "", nil
}

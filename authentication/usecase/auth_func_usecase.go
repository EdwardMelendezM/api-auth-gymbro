package usecase

import (
	"context"

	"github.com/EdwardMelendezM/api-auth/authentication/domain"

	"github.com/EdwardMelendezM/api-info-shared/error-log"
)

func (u authUseCase) Login(
	ctx context.Context,
	body domain.LoginBody,
) (token string, err error) {
	defer errorLog.PanicRecovery(&ctx, &err)
	var userIdAux string
	var tokenAux string

	var checkExistenceByUsername, errCheckExistenceByUsername = u.authRepository.CheckExistenceByUsername(ctx, body.Username)
	if errCheckExistenceByUsername != nil {
		return "", errCheckExistenceByUsername
	}
	if *checkExistenceByUsername != 1 {
		return "", u.err.Clone().CopyCodeDescription(domain.ErrNotFoundUsername).SetRaw(errCheckExistenceByUsername)
	}

	var userId, errCheckExistenceByPassword = u.authRepository.VerifyPassword(ctx, body)
	if errCheckExistenceByPassword != nil {
		return "", errCheckExistenceByPassword
	}

	userIdAux = *userId

	var tokenOfLogin, errTokenOfLogin = u.authJwtRepo.GenerateToken(userIdAux)
	if errTokenOfLogin != nil {
		return "", errTokenOfLogin
	}
	tokenAux = *tokenOfLogin

	return tokenAux, nil
}

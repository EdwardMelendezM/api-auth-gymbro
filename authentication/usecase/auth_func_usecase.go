package usecase

import (
	"context"

	"github.com/EdwardMelendezM/api-auth/authentication/domain"

	"github.com/EdwardMelendezM/info-code-api-shared-v1/error-log"
)

func (u authUseCase) Login(
	ctx context.Context,
	body domain.LoginBody,
) (token string, err error) {
	defer errorLog.PanicRecovery(&ctx, &err)

	var checkExistenceByUsername, errCheckExistenceByUsername = u.authRepository.CheckExistenceByUsername(ctx, body.Username)
	if errCheckExistenceByUsername != nil {
		return "", errCheckExistenceByUsername
	}
	if *checkExistenceByUsername != 1 {
		return "", u.err.Clone().CopyCodeDescription(domain.ErrNotFoundUsername).SetRaw(errCheckExistenceByUsername)
	}

	var checkExistenceByPassword, errCheckExistenceByPassword = u.authRepository.VerifyPassword(ctx, body)
	if errCheckExistenceByPassword != nil {
		return "", errCheckExistenceByPassword
	}
	if *checkExistenceByPassword != 1 {
		return "", u.err.Clone().CopyCodeDescription(domain.ErrNotFoundPassword).SetRaw(errCheckExistenceByPassword)
	}

	var checkUserStatus, errCheckUserStatus = u.authRepository.CheckAccountStatus(ctx, body)
	if errCheckUserStatus != nil {
		return "", errCheckUserStatus
	}

	if *checkUserStatus == 0 {
		return "", u.err.Clone().CopyCodeDescription(domain.ErrUserStatusNotActive).SetRaw(errCheckUserStatus)
	}

	return "TOKEN OF LOGIN", nil
}

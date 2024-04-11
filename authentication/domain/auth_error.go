package domain

import (
	"net/http"

	errLog "github.com/EdwardMelendezM/api-info-shared/error-log"
)

const (
	ErrNotFoundUsernameCode    = "ERR_NOT_FOUND_USERNAME"
	ErrNotFoundPasswordCode    = "ERR_NOT_FOUND_PASSWORD"
	ErrUserStatusNotActiveCode = "ERR_USER_STATUS_NOT_ACTIVE"
)

var (
	ErrNotFoundUsername = errLog.NewErr().
				SetCode(ErrNotFoundUsernameCode).
				SetDescription("USERNAME NOT FOUND").
				SetLevel(errLog.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errLog.UseCase).
				SetFunction("Login")

	ErrNotFoundPassword = errLog.NewErr().
				SetCode(ErrNotFoundPasswordCode).
				SetDescription("PASSWORD NOT FOUND").
				SetLevel(errLog.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errLog.UseCase).
				SetFunction("Login")

	ErrUserStatusNotActive = errLog.NewErr().
				SetCode(ErrUserStatusNotActiveCode).
				SetDescription("USER STATUS NOT ACTIVE").
				SetLevel(errLog.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errLog.UseCase).
				SetFunction("Login")
)

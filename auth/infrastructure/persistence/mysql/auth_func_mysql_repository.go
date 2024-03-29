package mysql

import (
	"auth-api/auth/domain"
	"auth-api/db"
	"context"
	_ "embed"

	"github.com/EdwardMelendezM/info-code-api-shared-v1/error-log"
)

//go:embed sql/check_existence_by_username.sql
var QueryCheckExistenceByUsername string

//go:embed sql/verify_password.sql
var QueryVerifyPassword string

func (r userMysqlRepo) CheckExistenceByUsername(
	ctx context.Context,
	username string,
) bool {
	var err error
	defer errorLog.PanicRecovery(ctx, &err)

	var totalTmp int
	err = db.Client.QueryRowContext(
		ctx,
		QueryCheckExistenceByUsername,
		username,
	).Scan(totalTmp)
	if err != nil {
		return false
	}
	return totalTmp > 0
}

func (r userMysqlRepo) VerifyPassword(
	ctx context.Context,
	body domain.LoginBody,
) bool {
	var err error
	defer errorLog.PanicRecovery(ctx, &err)

	var totalTmp int
	err = db.Client.QueryRowContext(
		ctx,
		QueryVerifyPassword,
		body.Username,
		body.Password,
	).Scan(totalTmp)
	if err != nil {
		return false
	}
	return totalTmp > 0
}

func (r userMysqlRepo) CheckAccountStatus(
	ctx context.Context,
	body domain.LoginBody,
) bool {
	var err error
	defer errorLog.PanicRecovery(ctx, &err)

	var statusTmp int
	err = db.Client.QueryRowContext(
		ctx,
		QueryVerifyPassword,
		body.Username,
		body.Password,
	).Scan(statusTmp)
	if err != nil {
		return false
	}
	return statusTmp > 0
}

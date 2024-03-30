package mysql

import (
	"context"
	_ "embed"

	"github.com/EdwardMelendezM/info-code-api-shared-v1/error-log"

	"github.com/EdwardMelendezM/api-auth/auth/domain"
	"github.com/EdwardMelendezM/api-auth/db"
)

//go:embed sql/check_existence_by_username.sql
var QueryCheckExistenceByUsername string

//go:embed sql/verify_password.sql
var QueryVerifyPassword string

func (r userMysqlRepo) CheckExistenceByUsername(
	ctx context.Context,
	username string,
) (countAccount int, err error) {
	defer errorLog.PanicRecovery(ctx, &err)

	err = db.Client.QueryRowContext(
		ctx,
		QueryCheckExistenceByUsername,
		username,
	).Scan(countAccount)
	if err != nil {
		return 0, err
	}
	return countAccount, err
}

func (r userMysqlRepo) VerifyPassword(
	ctx context.Context,
	body domain.LoginBody,
) (countAccount int, err error) {
	defer errorLog.PanicRecovery(ctx, &err)

	err = db.Client.QueryRowContext(
		ctx,
		QueryVerifyPassword,
		body.Username,
		body.Password,
	).Scan(countAccount)
	if err != nil {
		return 0, err
	}
	return countAccount, err
}

func (r userMysqlRepo) CheckAccountStatus(
	ctx context.Context,
	body domain.LoginBody,
) (statusAccount int, err error) {
	defer errorLog.PanicRecovery(ctx, &err)

	var statusTmp int
	err = db.Client.QueryRowContext(
		ctx,
		QueryVerifyPassword,
		body.Username,
		body.Password,
	).Scan(statusTmp)
	if err != nil {
		return -1, err
	}
	return statusTmp, err
}

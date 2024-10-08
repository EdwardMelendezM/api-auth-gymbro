package mysql

import (
	"context"
	_ "embed"

	"github.com/EdwardMelendezM/api-info-shared/db"
	errorLog "github.com/EdwardMelendezM/api-info-shared/error-log"

	"github.com/EdwardMelendezM/api-auth/authentication/domain"
)

//go:embed sql/check_existence_by_username.sql
var QueryCheckExistenceByUsername string

//go:embed sql/verify_password.sql
var QueryVerifyPassword string

//go:embed sql/check_account_status.sql
var QueryCheckAccountStatus string

func (r userMysqlRepo) CheckExistenceByUsername(
	ctx context.Context,
	username string,
) (countAccount *int, err error) {
	defer errorLog.PanicRecovery(&ctx, &err)

	var countAccountTmp int
	err = db.Client.QueryRowContext(
		ctx,
		QueryCheckExistenceByUsername,
		username,
	).Scan(&countAccountTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CheckExistenceByUsername").SetRaw(err)
	}
	countAccount = &countAccountTmp
	return countAccount, err
}

func (r userMysqlRepo) VerifyPassword(
	ctx context.Context,
	body domain.LoginBody,
) (userId *string, err error) {
	defer errorLog.PanicRecovery(&ctx, &err)

	var countAccountTmp string
	err = db.Client.QueryRowContext(
		ctx,
		QueryVerifyPassword,
		body.Username,
		body.Password,
	).Scan(&countAccountTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("VerifyPassword").SetRaw(err)
	}
	userId = &countAccountTmp
	return userId, err
}

func (r userMysqlRepo) CheckAccountStatus(
	ctx context.Context,
	body domain.LoginBody,
) (statusAccount *int, err error) {
	defer errorLog.PanicRecovery(&ctx, &err)

	var statusTmp int
	err = db.Client.QueryRowContext(
		ctx,
		QueryCheckAccountStatus,
		body.Username,
		body.Password,
	).Scan(&statusTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CheckAccountStatus").SetRaw(err)
	}
	statusAccount = &statusTmp
	return statusAccount, err
}

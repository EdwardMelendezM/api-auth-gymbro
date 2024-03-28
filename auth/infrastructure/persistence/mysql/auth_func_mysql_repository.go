package mysql

import (
	"auth-api/db"
	"context"
	_ "embed"

	"github.com/EdwardMelendezM/info-code-api-shared-v1/error-log"
)

//go:embed sql/check_existence_by_username.sql
var QueryCheckExistenceByUsername string

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

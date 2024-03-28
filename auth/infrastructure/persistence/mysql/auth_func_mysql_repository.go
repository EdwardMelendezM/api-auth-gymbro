package mysql

import (
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
}

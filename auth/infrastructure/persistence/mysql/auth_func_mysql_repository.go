package mysql

import (
	"context"
	_ "embed"

	"github.com/EdwardMelendezM/api-shared"
)

//go:embed sql/check_existence_by_username.sql
var QueryCheckExistenceByUsername string

func (r userMysqlRepo) CheckExistenceByUsername(
	ctx context.Context,
	username string,
) bool {
	defer errorLog.PanicRecovery(&ctx, &err)
}

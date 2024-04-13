package mysql

import (
	"time"

	"github.com/EdwardMelendezM/api-info-shared/clock"
	errorLog "github.com/EdwardMelendezM/api-info-shared/error-log"

	"github.com/EdwardMelendezM/api-auth/authentication/domain"
)

type userMysqlRepo struct {
	clock clock.Clock

	timeout time.Duration
	err     *errorLog.CustomError
}

func NewAuthRepository(
	clock clock.Clock,
	mongoTimeout int,
) domain.AuthRepository {
	rep := &userMysqlRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errorLog.NewErr().SetLayer(errorLog.Infra),
	}
	return rep
}

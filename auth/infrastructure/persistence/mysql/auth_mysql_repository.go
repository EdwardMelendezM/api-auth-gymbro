package mysql

import (
	"time"

	"auth-api/auth/domain"
	"github.com/EdwardMelendezM/info-code-api-shared-v1/clock"
)

type userMysqlRepo struct {
	clock   clock.Clock
	timeout time.Duration
}

func NewAuthRepository(
	clock clock.Clock,
	mongoTimeout int,
) domain.AuthRepository {
	rep := &userMysqlRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
	}
	return rep
}

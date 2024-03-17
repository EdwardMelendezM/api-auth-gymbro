package mysql

import (
	"auth-api/auth/domain"
	"time"

	"github.com/EdwardMelendezM/api-shared"
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

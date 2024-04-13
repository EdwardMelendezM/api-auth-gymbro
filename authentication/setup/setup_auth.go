package setup

import (
	"time"

	"github.com/gin-gonic/gin"

	authJwtRepo "github.com/EdwardMelendezM/api-info-shared/auth/infrastructure/jwt"
	"github.com/EdwardMelendezM/api-info-shared/clock"

	authDomain "github.com/EdwardMelendezM/api-auth/authentication/infrastructure/persistence/mysql"
	authHttpDelivery "github.com/EdwardMelendezM/api-auth/authentication/interfaces/rest"
	authUseCase "github.com/EdwardMelendezM/api-auth/authentication/usecase"
)

func LoadAuthentication(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	newClock := clock.NewClock()

	authRepository := authDomain.NewAuthRepository(newClock, 60)
	authJwtRepository := authJwtRepo.NewAuthRepository()
	authUCase := authUseCase.NewAuthUseCase(authRepository, timeoutContext, authJwtRepository)
	authHttpDelivery.NewAuthHandler(authUCase, router)
}

package setup

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/EdwardMelendezM/info-code-api-shared-v1/clock"

	authDomain "github.com/EdwardMelendezM/api-auth/authentication/infrastructure/persistence/mysql"
	authHttpDelivery "github.com/EdwardMelendezM/api-auth/authentication/interfaces"
	authUseCase "github.com/EdwardMelendezM/api-auth/authentication/usecase"
)

func LoadAuth(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	newClock := clock.NewClock()

	authRepository := authDomain.NewAuthRepository(newClock, 60)
	authUCase := authUseCase.NewAuthUseCase(authRepository, timeoutContext)
	authHttpDelivery.NewAuthHandler(authUCase, router)
}

package rest

import (
	"github.com/gin-gonic/gin"

	errorLog "github.com/EdwardMelendezM/info-code-api-shared-v1/error-log"

	authDomain "github.com/EdwardMelendezM/api-auth/authentication/domain"
)

type authHandler struct {
	authUseCase authDomain.AuthUseCase
	router      *gin.Engine
	err         *errorLog.CustomError
}

func NewAuthHandler(
	auth authDomain.AuthUseCase,
	router *gin.Engine,
) {
	handler := &authHandler{
		authUseCase: auth,
		router:      router,
		err:         errorLog.NewErr().SetLayer(errorLog.Interface),
	}

	apiAuth := router.Group("/api/v1/authentication")
	apiAuth.POST("/login", handler.Login)
}

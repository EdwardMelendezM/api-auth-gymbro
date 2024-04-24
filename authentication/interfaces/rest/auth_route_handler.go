package rest

import (
	authDomain "github.com/EdwardMelendezM/api-auth/authentication/domain"
	errorLog "github.com/EdwardMelendezM/api-info-shared/error-log"
	"github.com/gin-gonic/gin"
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

	apiAuth := router.Group("/api/v1/auth")
	apiAuth.POST("/login", handler.Login)
}

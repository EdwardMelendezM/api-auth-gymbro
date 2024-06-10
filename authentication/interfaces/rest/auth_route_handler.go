package rest

import (
	"github.com/gin-gonic/gin"

	errorLog "github.com/EdwardMelendezM/api-info-shared/error-log"

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

	apiAuth := router.Group("/api/v1/auth")
	apiAuth.POST("/login", handler.Login)
	apiAuth.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})
}

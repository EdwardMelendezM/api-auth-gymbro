package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	authDomain "github.com/EdwardMelendezM/api-auth/authentication/domain"

	restShared "github.com/EdwardMelendezM/info-code-api-shared-v1/rest"
)

func (h authHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var loginValidate LoginValidate
	if err := c.ShouldBindJSON(&loginValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("Login").SetRaw(errors.New("casting ValidationErrors"))
			restShared.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("Login").SetMessages(messagesErr)
		restShared.ErrJson(c, err)
		return
	}
	loginBody := authDomain.LoginBody{
		Username: loginValidate.Username,
		Password: loginValidate.Password,
	}

	token, err := h.authUseCase.Login(ctx, loginBody)
	if err != nil {
		restShared.ErrJson(c, err)
		return
	}

	res := LoginResult{
		Token:  &token,
		Status: http.StatusOK,
	}
	restShared.Json(c, http.StatusOK, res)
}

package api

import (
	"github.com/gin-gonic/gin"
	"go-ewallet/core/entity"
	"go-ewallet/core/module"
	"go-ewallet/pkg/http/response"
	"go-ewallet/pkg/validator"
	"net/http"
)

type authHandler struct {
	authUsecase module.AuthUsecase
}

type AuthHandler interface {
	AuthLogin(c *gin.Context)
	AuthLogout(c *gin.Context)
	Auth() gin.HandlerFunc
}

func NewAuthHandler(authUsecase module.AuthUsecase) AuthHandler {
	return &authHandler{authUsecase: authUsecase}
}

func (h *authHandler) AuthLogin(c *gin.Context) {
	ctx := c.Copy().Request.Context()
	var payload entity.AuthLoginRequest
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		response.Error(c, err)
		return
	}
	err = validator.Validate(payload)
	if err != nil {
		response.Error(c, err)
		return
	}

	resp, err := h.authUsecase.AuthLogin(ctx, payload)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, resp, http.StatusOK)
}

func (h *authHandler) AuthLogout(c *gin.Context) {
	ctx := c.Copy().Request.Context()
	user, err := FromCtx(c)
	if err != nil {
		response.Error(c, err)
		return
	}

	err = h.authUsecase.AuthLogout(ctx, user)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, nil, http.StatusOK)
}

func (h *authHandler) Auth() gin.HandlerFunc {
	return h.authUsecase.Auth()
}

func FromCtx(c *gin.Context) (*entity.JwtClaim, error) {
	tkn, ok := c.Get("user")
	if !ok {
		return nil, entity.ErrInvalidLogin
	}

	user, ok := tkn.(*entity.JwtClaim)
	if !ok || len(user.User.Email) < 0 {
		return nil, entity.ErrInvalidLogin
	}

	return user, nil
}

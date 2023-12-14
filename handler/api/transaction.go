package api

import (
	"github.com/gin-gonic/gin"
	"go-ewallet/core/entity"
	"go-ewallet/core/module"
	"go-ewallet/pkg/http/response"
	"go-ewallet/pkg/validator"
	"net/http"
)

type trxHandler struct {
	transactionUsecase module.TransactionUsecase
}

type TrxHandler interface {
	SendMoney(c *gin.Context)
}

func NewTrxHandler(transactionUsecase module.TransactionUsecase) TrxHandler {
	return &trxHandler{transactionUsecase: transactionUsecase}
}

func (t *trxHandler) SendMoney(c *gin.Context) {
	user, err := FromCtx(c)
	if err != nil {
		response.Error(c, err)
		return
	}

	payload := entity.SendMoneyRequest{
		SenderID:    user.User.ID,
		SenderEmail: user.User.Email,
	}

	err = c.ShouldBindJSON(&payload)
	if err != nil {
		response.Error(c, err)
		return
	}

	err = validator.Validate(&payload)
	if err != nil {
		response.Error(c, err)
		return
	}

	resp, err := t.transactionUsecase.SendMoney(c.Copy().Request.Context(), payload)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, resp, http.StatusCreated)
}

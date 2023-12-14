package entity

import (
	"go-ewallet/pkg/error"
	"net/http"
)

var (
	ErrPayloadInvalid      = error.NewErr(http.StatusBadRequest, "PAYLOAD_INVALID")
	ErrInvalidLogin        = error.NewErr(http.StatusUnauthorized, "INVALID_LOGIN")
	ErrInvalidTrxReceiver  = error.NewErr(http.StatusBadRequest, "INVALID_TRANSACTION_RECEIVER")
	ErrInsufficientBalance = error.NewErr(http.StatusBadRequest, "INSUFFICIENT_BALANCE")
	ErrUserNotFound        = error.NewErr(http.StatusNotFound, "USER_NOT_FOUND")
	ErrReceiverNotFound    = error.NewErr(http.StatusNotFound, "RECEIVER_NOT_FOUND")
	ErrTransactionFailed   = error.NewErr(http.StatusInternalServerError, "TRANSACTION_FAILED")
	ErrAuthNotFound        = error.NewErr(http.StatusUnauthorized, "AUTH_NOT_FOUND")
)

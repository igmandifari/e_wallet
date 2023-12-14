package repository

import (
	"context"
	"go-ewallet/core/entity"
)

type TransactionRepository interface {
	DoTransaction(ctx context.Context, req entity.DoTransactionRequest) error
}

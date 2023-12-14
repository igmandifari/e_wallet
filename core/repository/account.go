package repository

import (
	"context"
	"go-ewallet/core/entity"
)

type AccountRepository interface {
	GetAccountsByUserIDs(ctx context.Context, userIDs []string) ([]entity.Account, error)
}

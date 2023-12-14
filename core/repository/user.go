package repository

import (
	"context"
	"go-ewallet/core/entity"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}

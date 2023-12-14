package repository

import (
	"context"
	"go-ewallet/core/entity"
)

type AuthRepository interface {
	GetAuth(ctx context.Context, token string) (entity.Auth, error)
	CreateAuth(ctx context.Context, auth entity.Auth) error
	DeleteAuth(ctx context.Context, req *entity.JwtClaim) error
}

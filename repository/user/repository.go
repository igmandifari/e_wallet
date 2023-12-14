package user

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-ewallet/core/entity"
	"go-ewallet/core/repository"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	row, err := r.db.Query(ctx, "SELECT id, name, email, type, password FROM users WHERE email = $1", email)
	if err != nil {
		return entity.User{}, err
	}

	if err = pgxscan.ScanOne(&user, row); err != nil {
		if pgxscan.NotFound(err) {
			return entity.User{}, entity.ErrUserNotFound
		}
		return entity.User{}, err
	}

	return user, nil
}

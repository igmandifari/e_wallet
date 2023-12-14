package account

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
	"go-ewallet/core/entity"
	"go-ewallet/core/repository"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.AccountRepository {
	return &repo{db: db}
}

func (r *repo) GetAccountsByUserIDs(ctx context.Context, userIDs []string) ([]entity.Account, error) {
	var accounts []entity.Account

	rows, err := r.db.Query(ctx, "SELECT id, user_id, balance FROM accounts WHERE user_id = ANY($1)", pq.Array(userIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	err = pgxscan.ScanAll(&accounts, rows)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}


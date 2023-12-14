package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"go-ewallet/core/entity"
)

type ActivityHistoryRepository interface {
	CreateActivityHistoriesTx(ctx context.Context, tx pgx.Tx, activities []entity.ActivityHistory) (pgconn.CommandTag, error)
	CreateActivityHistory(ctx context.Context, activity entity.ActivityHistory) error
}

package activityhistory

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-ewallet/core/entity"
	"go-ewallet/core/repository"
	"go-ewallet/pkg/database"
	"strings"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.ActivityHistoryRepository {
	return &repo{db: db}
}

func (r *repo) CreateActivityHistoriesTx(ctx context.Context, tx pgx.Tx, activities []entity.ActivityHistory) (pgconn.CommandTag, error) {
	query := `INSERT INTO activity_histories (user_id, activity_type, description) VALUES `
	args := make([]interface{}, 0)
	bindings := make([]string, len(activities))
	for i, activity := range activities {
		bindings[i] = "(?, ?, ?)"
		args = append(args, activity.UserID, activity.ActivityType, activity.Description)
	}
	query = query + strings.Join(bindings, ",")
	query = database.Rebind(query)

	return tx.Exec(ctx, query, args...)
}

func (r *repo) CreateActivityHistory(ctx context.Context, activity entity.ActivityHistory) error {
	query := `INSERT INTO activity_histories (user_id, activity_type, description) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, activity.UserID, activity.ActivityType, activity.Description)
	return err
}

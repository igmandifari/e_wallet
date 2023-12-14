package auth

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-ewallet/core/entity"
	"go-ewallet/core/repository"
)

type repo struct {
	db           *pgxpool.Pool
	activityRepo repository.ActivityHistoryRepository
}

func NewRepository(db *pgxpool.Pool, activityRepo repository.ActivityHistoryRepository) repository.AuthRepository {
	return &repo{db: db, activityRepo: activityRepo}
}

func (r *repo) CreateAuth(ctx context.Context, auth entity.Auth) error {
	query := `INSERT INTO auths (user_id, uid, token, expired_at) VALUES ($1, $2, $3, $4)`
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	_, err = tx.Exec(ctx, query, auth.UserID, auth.UID, auth.Token, auth.ExpiredAt)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	_, err = r.activityRepo.CreateActivityHistoriesTx(ctx, tx, []entity.ActivityHistory{
		{
			UserID:       auth.UserID,
			ActivityType: entity.ActivityHistoryTypeLogin,
			Description:  "Login",
		},
	})
	if err != nil {
		return err
	}

	return tx.Commit(ctx)

}

func (r *repo) GetAuth(ctx context.Context, token string) (entity.Auth, error) {
	var auth entity.Auth
	query := `SELECT user_id, uid, token, expired_at FROM auths WHERE token = $1`
	row, err := r.db.Query(ctx, query, token)
	if err != nil {
		return auth, err
	}

	if err = pgxscan.ScanOne(&auth, row); err != nil {
		if pgxscan.NotFound(err) {
			return auth, entity.ErrAuthNotFound
		}
		return auth, err
	}
	return auth, err
}

func (r *repo) DeleteAuth(ctx context.Context, req *entity.JwtClaim) error {
	query := `DELETE FROM auths WHERE token = $1`

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	_, err = tx.Exec(ctx, query, req.Token)
	if err != nil {
		return err
	}
	_, err = r.activityRepo.CreateActivityHistoriesTx(ctx, tx, []entity.ActivityHistory{
		{
			UserID:       req.User.ID,
			ActivityType: entity.ActivityHistoryTypeLogout,
			Description:  "Logout",
		},
	})
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

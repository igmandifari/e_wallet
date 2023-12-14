package transaction

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-ewallet/core/entity"
	"go-ewallet/core/repository"
)

type repo struct {
	db                        *pgxpool.Pool
	activityHistoryRepository repository.ActivityHistoryRepository
}

func NewRepository(db *pgxpool.Pool, activityHistoryRepository repository.ActivityHistoryRepository) repository.TransactionRepository {
	return &repo{db: db, activityHistoryRepository: activityHistoryRepository}
}

func (r *repo) DoTransaction(ctx context.Context, req entity.DoTransactionRequest) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	decreaseBalanceSenderTx, err := tx.Exec(ctx, "UPDATE accounts SET balance = balance - $1 WHERE id = $2", req.Transaction.Amount, req.Transaction.AccountSenderID)
	if err != nil {
		return err
	}
	increaseBalanceReceiverTx, err := tx.Exec(ctx, "UPDATE accounts SET balance = balance + $1 WHERE id = $2", req.Transaction.Amount, req.Transaction.AccountReceiverID)
	if err != nil {
		return err
	}

	transactionTx, err := tx.Exec(ctx, "INSERT INTO transactions (id, sender_account_id, receiver_account_id, amount, status) VALUES ($1, $2, $3, $4, $5)",
		req.Transaction.ID, req.Transaction.AccountSenderID, req.Transaction.AccountReceiverID, req.Transaction.Amount, req.Transaction.Status)
	if err != nil {
		return err
	}

	historiesTx, err := r.activityHistoryRepository.CreateActivityHistoriesTx(ctx, tx, req.ActivityHistories)
	if err != nil {
		return err
	}

	if decreaseBalanceSenderTx.RowsAffected() == 0 || increaseBalanceReceiverTx.RowsAffected() == 0 || transactionTx.RowsAffected() == 0 || historiesTx.RowsAffected() == 0 {
		return entity.ErrTransactionFailed
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

package module

import (
	"context"
	"fmt"
	"github.com/oklog/ulid/v2"
	"go-ewallet/core/entity"
	"go-ewallet/core/repository"
)

type transactionUsecase struct {
	userRepo    repository.UserRepository
	trxRepo     repository.TransactionRepository
	accountRepo repository.AccountRepository
}

type TransactionUsecase interface {
	SendMoney(ctx context.Context, req entity.SendMoneyRequest) (entity.SendMoneyResponse, error)
}

func NewTransactionUsecase(userRepo repository.UserRepository, trxRepo repository.TransactionRepository, accountRepo repository.AccountRepository) TransactionUsecase {
	return &transactionUsecase{userRepo: userRepo, trxRepo: trxRepo, accountRepo: accountRepo}
}

func (t *transactionUsecase) SendMoney(ctx context.Context, req entity.SendMoneyRequest) (entity.SendMoneyResponse, error) {
	if req.SenderEmail == req.To {
		return entity.SendMoneyResponse{}, entity.ErrInvalidTrxReceiver
	}

	receiver, err := t.userRepo.GetUserByEmail(ctx, req.To)
	if err != nil {
		if err == entity.ErrUserNotFound {
			return entity.SendMoneyResponse{}, entity.ErrReceiverNotFound
		}
		return entity.SendMoneyResponse{}, err
	}

	accounts, err := t.accountRepo.GetAccountsByUserIDs(ctx, []string{req.SenderID, receiver.ID})
	if err != nil {
		return entity.SendMoneyResponse{}, err
	}

	if len(accounts) != 2 {
		return entity.SendMoneyResponse{}, entity.ErrInvalidTrxReceiver
	}

	trxReq := entity.Transaction{
		ID:         ulid.Make().String(),
		SenderID:   req.SenderID,
		ReceiverID: receiver.ID,
		Amount:     req.Amount,
		Status:     entity.TransactionDescriptionSuccess,
	}

	var activityHistory []entity.ActivityHistory
	for _, account := range accounts {
		if account.UserID == req.SenderID {
			if account.Balance < req.Amount {
				return entity.SendMoneyResponse{}, entity.ErrInsufficientBalance
			}

			trxReq.AccountSenderID = account.ID

			activityHistory = append(activityHistory, entity.ActivityHistory{
				UserID:       req.SenderID,
				ActivityType: entity.ActivityHistorySendMoney,
				Description:  fmt.Sprintf("Send money %v to %s ", req.Amount, req.To),
			})
		} else if account.UserID == receiver.ID {
			trxReq.AccountReceiverID = account.ID
			activityHistory = append(activityHistory, entity.ActivityHistory{
				UserID:       receiver.ID,
				ActivityType: entity.ActivityHistoryReceiveMoney,
				Description:  fmt.Sprintf("Receive money %v from  %s", req.Amount, req.SenderEmail),
			})
		}
	}

	err = t.trxRepo.DoTransaction(ctx, entity.DoTransactionRequest{
		Transaction:       trxReq,
		ActivityHistories: activityHistory,
	})
	if err != nil {
		return entity.SendMoneyResponse{}, err
	}

	return entity.SendMoneyResponse{
		TransactionID: trxReq.ID,
	}, nil
}

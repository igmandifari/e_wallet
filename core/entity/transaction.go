package entity

type TransactionDescription string

const (
	TransactionDescriptionPending   TransactionDescription = "pending"
	TransactionDescriptionSuccess   TransactionDescription = "completed"
	TransactionDescriptionFailed    TransactionDescription = "failed"
	TransactionDescriptionCancelled TransactionDescription = "cancelled"
)

type SendMoneyRequest struct {
	SenderID    string  `validate:"required"`
	SenderEmail string  `validate:"required,email"`
	To          string  `json:"to" validate:"required,email"`
	Amount      float64 `json:"amount" validate:"required,numeric,min=1"`
}

type SendMoneyResponse struct {
	TransactionID string `json:"transaction_id"`
}

type Transaction struct {
	ID                string
	SenderID          string
	ReceiverID        string
	AccountSenderID   string
	AccountReceiverID string
	Amount            float64
	Status            TransactionDescription
}

type DoTransactionRequest struct {
	Transaction       Transaction
	ActivityHistories []ActivityHistory
}

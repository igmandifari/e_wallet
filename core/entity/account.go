package entity

type Account struct {
	ID      string  `json:"id"`
	UserID  string  `json:"user_id"`
	Balance float64 `json:"balance"`
}

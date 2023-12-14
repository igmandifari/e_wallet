package entity

type UserType string

const (
	CustomerUser UserType = "customer"
	MerchantUser UserType = "merchant"
)

type User struct {
	ID       string   `json:"id" db:"id"`
	Name     string   `json:"name" db:"name"`
	Email    string   `json:"email" db:"email"`
	Type     UserType `json:"type" db:"type"`
	Password string   `json:"-" db:"password"`
}

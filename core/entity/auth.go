package entity

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthLoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type Auth struct {
	UserID    string    `db:"user_id"`
	UID       string    `db:"uid"`
	Token     string    `db:"token"`
	ExpiredAt time.Time `db:"expired_at"`
}

type JwtClaim struct {
	User  User   `json:"user"`
	Token string `json:"token"`
	jwt.StandardClaims
}

package module

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-ewallet/core/entity"
	"go-ewallet/core/repository"
	"go-ewallet/pkg/crypt"
	"go-ewallet/pkg/http/response"
	"strings"
	"time"
)

type authUsecase struct {
	userRepo  repository.UserRepository
	authRepo  repository.AuthRepository
	secretKey string
	expTime   time.Duration
}

type AuthUsecase interface {
	AuthLogin(ctx context.Context, req entity.AuthLoginRequest) (entity.AuthLoginResponse, error)
	AuthLogout(ctx context.Context, req *entity.JwtClaim) error
	Auth() gin.HandlerFunc
}

func NewAuthUsecase(userRepo repository.UserRepository, authRepo repository.AuthRepository, secretKey string, expTime time.Duration) AuthUsecase {
	return &authUsecase{userRepo: userRepo, secretKey: secretKey, expTime: expTime, authRepo: authRepo}
}

func (a *authUsecase) AuthLogin(ctx context.Context, req entity.AuthLoginRequest) (entity.AuthLoginResponse, error) {
	user, err := a.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == entity.ErrUserNotFound {
			return entity.AuthLoginResponse{}, entity.ErrInvalidLogin
		}
		return entity.AuthLoginResponse{}, err
	}

	if !crypt.Compare([]byte(req.Password), []byte(user.Password)) {
		return entity.AuthLoginResponse{}, entity.ErrInvalidLogin
	}
	claims, err := a.generateToken(user)
	if err != nil {
		return entity.AuthLoginResponse{}, err
	}
	err = a.authRepo.CreateAuth(ctx, entity.Auth{
		UserID:    user.ID,
		Token:     claims.Token,
		UID:       claims.Id,
		ExpiredAt: time.Unix(claims.ExpiresAt, 0),
	})

	if err != nil {
		return entity.AuthLoginResponse{}, err
	}

	return entity.AuthLoginResponse{
		Token: claims.Token,
		User:  user,
	}, nil
}

func (a *authUsecase) generateToken(user entity.User) (*entity.JwtClaim, error) {

	expireAt := time.Now().Add(a.expTime)
	claims := &entity.JwtClaim{
		User:  user,
		Token: "",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.secretKey))

	claims.Token = tokenString

	return claims, err
}

func (a *authUsecase) AuthLogout(ctx context.Context, req *entity.JwtClaim) error {
	return a.authRepo.DeleteAuth(ctx, req)
}

func (a *authUsecase) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.Request.Header.Get("Authorization")
		if bearToken == "" {
			response.Error(c, entity.ErrInvalidLogin)
			return
		}

		strArr := strings.Split(bearToken, " ")
		if len(strArr) != 2 {
			response.Error(c, entity.ErrInvalidLogin)
			return
		}
		token := strArr[1]

		_, err := a.authRepo.GetAuth(c.Copy().Request.Context(), token)
		if err != nil {
			if err == entity.ErrAuthNotFound {
				response.Error(c, entity.ErrInvalidLogin)
				return
			}
			response.Error(c, err)
			return
		}

		tk, err := a.verifyToken(token)
		if err != nil {
			response.Error(c, entity.ErrInvalidLogin)
			return

		}

		if !tk.Valid {
			response.Error(c, entity.ErrInvalidLogin)
			return
		}

		user, isTokenClaim := tk.Claims.(*entity.JwtClaim)
		if !isTokenClaim {
			response.Error(c, entity.ErrInvalidLogin)
			return
		}

		user.Token = strArr[1]
		c.Set("user", user)
		c.Next()
	}
}

func (a *authUsecase) verifyToken(tokenString string) (*jwt.Token, error) {
	claims := &entity.JwtClaim{Token: tokenString}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

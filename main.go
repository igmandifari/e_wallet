package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-ewallet/config"
	"go-ewallet/core/module"
	"go-ewallet/handler/api"
	"go-ewallet/pkg/database"
	"go-ewallet/repository/account"
	activityhistory "go-ewallet/repository/activity-history"
	"go-ewallet/repository/auth"
	"go-ewallet/repository/transaction"
	"go-ewallet/repository/user"
	"log"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewDatabase(cfg.DbURL)
	if err != nil {
		log.Fatal(err)
	}

	activityRepo := activityhistory.NewRepository(db)
	authRepo := auth.NewRepository(db, activityRepo)
	userRepo := user.NewRepository(db)
	trxRepo := transaction.NewRepository(db, activityRepo)
	accountRepo := account.NewRepository(db)
	authUsecase := module.NewAuthUsecase(userRepo, authRepo, cfg.SecretKey, cfg.DurationExpire)
	trxUsecase := module.NewTransactionUsecase(userRepo, trxRepo, accountRepo)
	authHandler := api.NewAuthHandler(authUsecase)
	trxHandler := api.NewTrxHandler(trxUsecase)

	router := gin.Default()

	router.POST("/auth/login", authHandler.AuthLogin)
	router.Use(authHandler.Auth())
	router.POST("/auth/logout", authHandler.AuthLogout)
	router.POST("/transaction/send-money", trxHandler.SendMoney)
	err = router.Run(fmt.Sprintf(":%s", cfg.HTTPPort))
	if err != nil {
		log.Fatal(err)
	}
}

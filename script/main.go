package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oklog/ulid/v2"
	"go-ewallet/config"
	"go-ewallet/core/entity"
	"go-ewallet/pkg/crypt"
	"go-ewallet/pkg/database"
	"log"
	"strings"
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

	users, accounts := getPayload()

	err = insert(context.Background(), db, users, accounts)
	if err != nil {
		log.Fatal(err)
	}
}

func getPayload() ([]entity.User, []entity.Account) {
	passwordHash, err := crypt.Generate("1234567890")
	if err != nil {
		log.Fatal(err)
	}
	users := []entity.User{
		{
			ID:       ulid.Make().String(),
			Email:    "igman.difari@gmail.com",
			Name:     "Igman Difari",
			Password: string(passwordHash),
			Type:     entity.CustomerUser,
		},
		{
			ID:       ulid.Make().String(),
			Email:    "geprek_nadila@gmail.com",
			Name:     "ayam geprek nadila",
			Password: string(passwordHash),
			Type:     entity.MerchantUser,
		},
		{
			ID:       ulid.Make().String(),
			Email:    "irfan1234@gmail.com",
			Name:     "irfan",
			Password: string(passwordHash),
			Type:     entity.CustomerUser,
		},
	}

	var accounts []entity.Account
	for _, user := range users {
		accounts = append(accounts, entity.Account{
			ID:      ulid.Make().String(),
			UserID:  user.ID,
			Balance: 1000000,
		})
	}
	return users, accounts

}

func insert(ctx context.Context, db *pgxpool.Pool, users []entity.User, accounts []entity.Account) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	argsUser := make([]interface{}, 0)
	bindingUser := make([]string, 0)
	for _, user := range users {
		argsUser = append(argsUser, user.ID, user.Email, user.Name, user.Password, user.Type)
		bindingUser = append(bindingUser, "(?, ?, ?, ?, ?)")
	}
	queryInsertUsers := "INSERT INTO users (id, email, name, password, type) VALUES "
	queryInsertUsers = queryInsertUsers + strings.Join(bindingUser, ",")
	queryInsertUsers = database.Rebind(queryInsertUsers)
	_, err = tx.Exec(ctx, queryInsertUsers, argsUser...)
	if err != nil {
		return err
	}

	argsAccount := make([]interface{}, 0)
	bindingAccount := make([]string, 0)
	for _, account := range accounts {
		argsAccount = append(argsAccount, account.ID, account.UserID, account.Balance)
		bindingAccount = append(bindingAccount, "(?, ?, ?)")
	}
	queryInsertAccounts := "INSERT INTO accounts (id, user_id, balance) VALUES "
	queryInsertAccounts = queryInsertAccounts + strings.Join(bindingAccount, ",")
	queryInsertAccounts = database.Rebind(queryInsertAccounts)
	_, err = tx.Exec(ctx, queryInsertAccounts, argsAccount...)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

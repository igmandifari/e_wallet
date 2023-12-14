package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"strconv"
	"strings"
)

func NewDatabase(dbUrl string) (*pgxpool.Pool, error) {

	var config *pgxpool.Config
	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func Rebind(query string) string {
	rqb := make([]byte, 0, len(query)+10)

	var i, j int

	for i = strings.Index(query, "?"); i != -1; i = strings.Index(query, "?") {
		rqb = append(rqb, query[:i]...)

		rqb = append(rqb, '$')

		j++
		rqb = strconv.AppendInt(rqb, int64(j), 10)

		query = query[i+1:]
	}

	return string(append(rqb, query...))
}

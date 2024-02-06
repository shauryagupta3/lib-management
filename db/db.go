package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbConn *pgxpool.Pool

func ConnectPostgres(url string) error {
	conn, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return err
	}
	dbConn = conn
	return nil
}

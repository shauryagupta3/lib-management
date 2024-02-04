package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type dbPgx struct {
	conn *pgxpool.Pool
}

func initPostgres() (*dbPgx, error) {
	err := godotenv.Load()
	conn, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))

	if err != nil {
		return nil, err
	}

	return &dbPgx{
		conn: conn,
	}, nil
}

func (s *dbPgx) createTables() error {

	if err := s.createMembersTable(); err != nil {
		return err
	}
	if err := s.createBooksTable(); err != nil {
		return err
	}
	if err := s.createAuthorsTable(); err != nil {
		return err
	}
	return nil
}

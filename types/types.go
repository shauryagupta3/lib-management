package types

import "github.com/jackc/pgx/v5/pgxpool"

type DB struct {
	PostgresConnection *pgxpool.Pool
}

type Member struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Status string `json:"status"`
}

type Book struct {
	Id      int      `json:"id"`
	Title   string   `json:"title"`
	Year    int      `json:"year"`
	Genre   string   `json:"genre"`
	Authors []Author `json:"authors"`
}

type Author struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

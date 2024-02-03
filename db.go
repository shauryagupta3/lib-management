package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type dbPgx struct {
	conn pgx.Conn
}

func initPostgres() (*dbPgx, error) {
	err := godotenv.Load()
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))

	if err != nil {
		return nil, err
	}

	return &dbPgx{
		conn: *conn,
	}, nil
}

func (s *dbPgx) getMembers() ([]Member, error) {
	members := []Member{}
	rows, err := s.conn.Query(context.Background(), "select id,full_name,phone_number,current_status from Members")
	if err != nil {
		return nil, err
	}

	var Name string
	var Phone string
	var id int
	var status string

	for rows.Next() {
		err := rows.Scan(&id, &Name, &Phone, &status)
		if err != nil {
			return nil, err
		}
		members = append(members, Member{id, Name, Phone, status})
	}
	return members, nil
}

func (s *dbPgx) getMemberById(id int) (*Member, error) {
	var a Member

	err := s.conn.QueryRow(context.Background(), "select id,full_name,phone_number,current_status from Members where id=$1", id).Scan(&a.Id, &a.Name, &a.Phone, &a.Status)
	if err != nil {
		return nil, err
	}

	return &a, err
}

func (s *dbPgx) createMemeber(a Member) error {

	_, err := s.conn.Query(context.Background(), "insert into Members(full_name,phone_number) values ($1,$2)", a.Name, a.Phone)
	if err != nil {
		return err
	}

	return nil
}

package main

import "context"

func (s *dbPgx) createAuthor(a *Author) error {
	_, err := s.conn.Query(context.Background(), "insert into authors(full_name,date_of_birth) values ($1,$2)", a.Name, a.DOB)
	if err != nil {
		return err
	}

	return nil

}

func (s *dbPgx) createAuthorsTable() error {
	_, err := s.conn.Query(context.Background(), "create table if not exists authors(id serial primary key, full_name varchar(255) not null, date_of_birth date)")
	if err != nil {
		return err
	}
	return nil
}

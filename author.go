package main

import "context"

func (s *dbPgx) createAuthor(a *Author) error {
	_, err := s.conn.Query(context.Background(), "insert into authors(full_name) values ($1)", a.Name)
	if err != nil {
		return err
	}

	return nil

}

func (s *dbPgx) createAuthorsTable() error {
	_, err := s.conn.Query(context.Background(), "create table if not exists authors(id serial primary key, full_name varchar(255) unique not null)")
	if err != nil {
		return err
	}
	return nil
}

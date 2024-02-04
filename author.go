package main

import "context"

func (s *dbPgx) createAuthor(a *Author) error {
	err := s.conn.QueryRow(context.Background(), "insert into authors(full_name) values ($1) returning id", a.Name).Scan(&a.Id)
	if err != nil {
		if err := s.conn.QueryRow(context.Background(), "select id from authors where full_name=$1", a.Name).Scan(&a.Id); err != nil {
			return err
		}
	}

	return nil
}

func (s *dbPgx) getAuthorByID(id int) (*Author, error) {
	var a Author

	err := s.conn.QueryRow(context.Background(), "select id,full_name from books where id=$1", id).Scan(&a.Id, &a.Name)
	if err != nil {
		return nil, err
	}

	return &a, err
}

func (s *dbPgx) createAuthorsTable() error {
	_, err := s.conn.Query(context.Background(), "create table if not exists authors(id serial primary key, full_name varchar(255) unique not null)")
	if err != nil {
		return err
	}
	return nil
}

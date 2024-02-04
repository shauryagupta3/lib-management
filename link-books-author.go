package main

import "context"

func (s *dbPgx) createLinkBooksAuthorTable() error {
	_, err := s.conn.Query(context.Background(), "create table if not exists linkBooksToAuthors(book_id int references books (id) on update cascade, authors_id int references authors (id) on update cascade,primary key (book_id,authors_id))")
	if err != nil {
		return err
	}
	return nil
}

func (s *dbPgx) getAuthorsByBookID(a *Book) error {
	authors := []Author{}
	rows, err := s.conn.Query(context.Background(), "select * from linkBooksToAuthors where book_id=$1", a.Id)
	if err != nil {
		return err
	}

	for rows.Next() {
		var author Author
		if err := rows.Scan(&author); err != nil {
			return nil
		}
		authors = append(authors, author)
	}

	a.Authors = authors
	return nil
}

func (s *dbPgx) linkBookToAuthors(a *Book) error {
	for _, author := range a.Authors {
		if _, err := s.conn.Query(context.Background(), "insert into linkBooksToAuthors(book_id,authors_id) values ($1,$2)", a.Id, author.Id); err != nil {
			return err
		}
	}

	return nil
}

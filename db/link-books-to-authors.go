package db

import "context"

func LinkBookToAuthor(author_id int, book_id int) error {
	err := dbConn.QueryRow(context.Background(), "insert into linkbooktoauthors(book_id,author_id) values($1,$2)", book_id, author_id).Scan()
	if err != nil {
		return err
	}
	return nil
}

package db

import (
	"context"
	"libraryManagement/types"

	"github.com/jackc/pgx/v5"
)

func LinkBookToAuthor(author_id int, book_id int) error {
	_, err := dbConn.Exec(context.Background(), "insert into linkbookstoauthors(books_id,authors_id) values($1,$2)", book_id, author_id)
	if err != nil {
		return err
	}
	return nil
}

func GetAuthorsOfBook(a *types.Book) error {
	var name string
	var id int

	rows, _ := dbConn.Query(context.Background(), "select id,full_name from linkbookstoauthors as a inner join authors as b on a.authors_id=b.id where a.books_id=$1", a.Id)
	_, err := pgx.ForEachRow(rows, []any{&id, &name}, func() error {
		x := types.Author{Name: name, Id: id}
		a.Authors = append(a.Authors, x)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func GetBooksOfAuthor(id int) (*[]types.Book, error) {
	var books []types.Book

	var title string
	var Bid int
	var year int
	var genre string

	rows, _ := dbConn.Query(context.Background(), "select id,title,genre,release_year from linkbookstoauthors as a inner join books as b on a.books_id=b.id where a.authors_id=$1", id)
	_, err := pgx.ForEachRow(rows, []any{&Bid, &title, &genre, &year}, func() error {
		x := types.Book{Title: title, Year: year, Id: Bid, Genre: genre}
		books = append(books, x)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &books, nil
}

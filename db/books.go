package db

import (
	"context"
	"libraryManagement/types"
)

func GetBook(id int) (*types.Book, error) {
	var a types.Book

	err := dbConn.QueryRow(context.Background(), "select id,title,release_year,genre from books where id=$1", id).Scan(&a.Id, &a.Title, &a.Year, &a.Genre)
	if err != nil {
		return nil, err
	}

	err = GetAuthorsOfBook(&a)
	if err != nil {
		return nil, err
	}
	return &a, nil

}

func CreateBook(a *types.Book) error {

	err := dbConn.QueryRow(context.Background(), "insert into books(title,release_year,genre) values($1,$2,$3) returning id", a.Title, a.Year, a.Genre).Scan(&a.Id)
	if err != nil {
		return err
	}
	for index := range a.Authors {
		err := createAuthorIfNotExists(&a.Authors[index])
		if err != nil {
			return err
		}
		err = LinkBookToAuthor(a.Authors[index].Id, a.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetBooksByAuthorID(id int) (*[]types.Book, error) {
	books, err := GetBooksOfAuthor(id)
	if err != nil {
		return nil, err
	}
	return books, nil
}

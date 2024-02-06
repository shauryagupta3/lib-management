package db

import (
	"context"
	"fmt"
	"libraryManagement/types"
)

func GetBook(id int) (*types.Book, error) {
	var a types.Book

	err := dbConn.QueryRow(context.Background(), "select id,title,release_year,genre from books where id=$1", id).Scan(&a.Id, &a.Title, &a.Year, &a.Genre)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &a, nil

}

func CreateBook(a *types.Book) error {
	for index := range a.Authors {
		err:=createAuthorIfNotExists(&a.Authors[index])
		if err!=nil {
			return err
		}
	}
	return nil
}

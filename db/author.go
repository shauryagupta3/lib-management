package db

import (
	"context"
	"libraryManagement/types"
)

func getAuthorByName(a *types.Author) error {
	err := dbConn.QueryRow(context.Background(), "select id,full_name from authors where full_name=$1", a.Name).Scan(&a.Id, &a.Name)
	if err != nil {
		return err
	}
	return nil
}

func createAuthor(a *types.Author) error {
	err := dbConn.QueryRow(context.Background(), "insert into authors(full_name) values ($1) returning id", a.Name).Scan(&a.Id)
	if err != nil {
		return err
	}
	return nil
}

func createAuthorIfNotExists(a *types.Author) error {

	err := getAuthorByName(a)
	if err != nil {
		err := createAuthor(a)
		if err != nil {
			return err
		}
	}
	return nil
}

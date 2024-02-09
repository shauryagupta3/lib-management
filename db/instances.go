package db

import (
	"context"
	"libraryManagement/types"
)

func CreateNewInstance(a *types.Instance) error {
	if err := dbConn.QueryRow(context.Background(), "insert into instances(book_id) values($1) returning id", a.BookId).Scan(&a.ID); err != nil {
		return err
	}
	return nil
}

func DeleteInstance(a *types.Instance) error {
	if err := dbConn.QueryRow(context.Background(), "delete from instances where id=$1 returning book_id", a.ID).Scan(&a.BookId); err != nil {
		return err
	}
	return nil
}

func UpdateAvailableStatus(id int) error {
	if _, err := dbConn.Exec(context.Background(), "update instances set available = not available where id=$1", id); err != nil {
		return err
	}
	return nil
}

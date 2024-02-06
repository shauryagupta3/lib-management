package db

import (
	"context"
	"libraryManagement/types"
)

func GetMember(id int) (*types.Member, error) {
	var a types.Member

	err := dbConn.QueryRow(context.Background(), "select id,full_name,phone_number,current_status from members where id=$1", id).Scan(&a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

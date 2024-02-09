package db

import (
	"context"
	"fmt"
	// "errors"
	"libraryManagement/types"
)

func GetMember(id int) (*types.Member, error) {
	var a types.Member

	err := dbConn.QueryRow(context.Background(), "select id,full_name,phone_number,joined_at,active from members where id=$1", id).Scan(&a.ID, &a.Name, &a.Phone, &a.JoinedAt, &a.Active)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func CreateNewMember(a *types.Member) error {
	// if err := CheckMemberExistsByPhone(a); err == nil {
	// 	return errors.New("user already exists")
	// }

	if err := dbConn.QueryRow(context.Background(), "insert into members(full_name,phone_number) values ($1,$2) returning id,joined_at,active,expires_at", a.Name, a.Phone).Scan(&a.ID, &a.JoinedAt,&a.Active,&a.ExpiresAt); err != nil {
		fmt.Println(err,"i am here")
		return err
	}
	return nil
}

func CheckMemberExistsByPhone(a *types.Member) error {
	if err := dbConn.QueryRow(context.Background(), "select id,full_name,joined_at,active from members where phone_number=$1", a.Phone).Scan(&a.ID, &a.Name, a.JoinedAt, &a.Active); err != nil {
		return err
	}
	return nil
}

func CheckMemberActive(id int) (bool, error) {
	var check bool
	if err := dbConn.QueryRow(context.Background(), "select active from members where id=$1", id).Scan(&check); err != nil {
		return false, err
	}
	return check, nil
}

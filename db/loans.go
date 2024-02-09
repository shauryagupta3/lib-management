package db

import (
	"context"
	"errors"
	"libraryManagement/types"
)

func CreateNewLoan(a *types.Loan) error {

	if check, _ := CheckMemberActive(a.MemberID); check != true {
		return errors.New("Error, Please check member again")
	}

	if err := dbConn.QueryRow(context.Background(), "insert into loans(instance_id,member_id) values ($1,$2) returning id,issued_at", a.InstanceID, a.MemberID).Scan(&a.ID, &a.IssuedAt); err != nil {
		return err
	}

	if err := UpdateAvailableStatus(a.InstanceID); err != nil {
		return err
	}
	return nil
}

func CompleteLoan(a *types.Loan) error {
	if err:=dbConn.QueryRow(context.Background(),"update loans set pending = not pending,returned_at = CURRENT_DATE where id=$1 returning returned_at,pending",a.ID).Scan(&a.ReturnedAt,&a.Pending);err!=nil {
		return err
	}
	
	if err:=UpdateAvailableStatus(a.InstanceID);err!=nil{
		return err
	}

	return nil
}

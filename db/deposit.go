package db

import (
	"github.com/deposinator/models"
)

func DepositCreate(issuer int, members []int, amount int, description string) error {
	q := "INSERT INTO deposits (issuer, amount, description) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err := db.QueryRow(q, issuer, amount, description).Scan(&id)
	if err != nil {
		return err
	}
	return insertDepositMembers(id, members)
}

func DepositUpdate(id int, issuer int, members []int, amount int, description string) error {
	q := "UPDATE deposits SET issuer = $1, amount = $2, description = $3 where id = $4"
	_, err := db.Exec(q, issuer, amount, description, id)
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM deposit_members where deposit_id = $1", id)
	if err != nil {
		return err
	}
	return insertDepositMembers(id, members)
}

func DepositDelete(id int) error {
	q := "DELETE FROM deposits where id = $1"
	_, err := db.Exec(q, id)
	return err
}

func DepositFromId(id int) *models.Deposit {
	var deposit models.Deposit
	err := db.Get(&deposit, "SELECT * FROM deposits WHERE id = $1", id)
	if err != nil {
		return &models.Deposit{}
	}
	return &deposit
}

func insertDepositMembers(deposit_id int, members []int) error {
	for _, member_id := range members {
		q := "INSERT INTO deposit_members (deposit_id, user_id) VALUES ($1, $2)"
		_, err := db.Exec(q, deposit_id, member_id)
		if err != nil {
			return err
		}
	}
	return nil
}

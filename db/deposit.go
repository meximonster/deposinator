package db

import (
	"github.com/deposinator/models"
	"github.com/deposinator/utils"
)

func GetDeposits(query string, args ...interface{}) ([]models.Deposit, error) {
	var deposits []models.Deposit
	rows, err := db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows
	for rows.Next() {
		var d models.Deposit
		var membersStr string

		// Scan the row into the struct and the temporary members string
		err := rows.Scan(
			&d.Id,
			&d.Issuer,
			&membersStr, // Scan into string
			&d.Amount,
			&d.Description,
			&d.Created_at,
		)
		if err != nil {
			return nil, err
		}

		// Parse the PostgreSQL array string into []int
		d.Members, err = utils.ParseArray(membersStr)
		if err != nil {
			return nil, err
		}

		deposits = append(deposits, d)
	}

	return deposits, nil
}

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

package models

import "time"

type Withdraw struct {
	Id          int
	Issuer      string
	Deposit_id  int
	Amount      int
	Description string
	Created_at  time.Time
}

func WithdrawCreate(issuer string, deposit_id int, amount int, description string) (int, error) {
	q := "INSERT INTO withdraws (issuer, deposit_id, amount, description) VALUES ($1, $2, $3, $4) RETURNING id"
	var id int
	err := db.QueryRow(q, issuer, deposit_id, amount, description).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func WithdrawUpdate(id int, deposit_id int, issuer string, amount int, description string) error {
	q := "UPDATE withdraws SET issuer = $1, deposit_id = $2, amount = $3, description = $4 where id = $5"
	_, err := db.Exec(q, issuer, deposit_id, amount, description, id)
	return err
}

func WithdrawDelete(id int) error {
	q := "DELETE FROM withdraws where id = $1"
	_, err := db.Exec(q, id)
	return err
}

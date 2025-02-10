package db

func WithdrawCreate(issuer int, deposit_id int, amount int, description string) error {
	q := "INSERT INTO withdraws (issuer, deposit_id, amount, description) VALUES ($1, $2, $3, $4) RETURNING id"
	_, err := db.Exec(q, issuer, deposit_id, amount, description)
	if err != nil {
		return err
	}
	return nil
}

func WithdrawUpdate(id int, issuer int, deposit_id int, amount int, description string) error {
	q := "UPDATE withdraws SET issuer = $1, deposit_id = $2, amount = $3, description = $4 where id = $5"
	_, err := db.Exec(q, issuer, deposit_id, amount, description, id)
	return err
}

func WithdrawDelete(id int) error {
	q := "DELETE FROM withdraws where id = $1"
	_, err := db.Exec(q, id)
	return err
}

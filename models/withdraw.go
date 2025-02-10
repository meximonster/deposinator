package models

import (
	"errors"
	"time"
)

type Withdraw struct {
	Id          int       `db:"id" json:"id,omitempty"`
	Issuer      int       `db:"issuer" json:"issuer"`
	Deposit_id  int       `db:"deposit_id" json:"deposit_id"`
	Amount      int       `db:"amount" json:"amount"`
	Description string    `db:"description" json:"description"`
	Created_at  time.Time `db:"created_at" json:"created_at,omitempty"`
}

func (w *Withdraw) Validate() error {
	if len(w.Description) > 100 {
		return errors.New("invalid description length")
	}
	return nil
}

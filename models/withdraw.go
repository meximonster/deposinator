package models

import (
	"errors"
	"time"
)

type Withdraw struct {
	Id          int
	Issuer      string
	Deposit_id  int
	Amount      int
	Description string
	Created_at  time.Time
}

func (w *Withdraw) Validate() error {
	if len(w.Description) > 100 {
		return errors.New("invalid description length")
	}
	return nil
}

package models

import (
	"errors"
	"time"
)

type Withdraw struct {
	Id          int       `json:"id,omitempty"`
	Issuer      int       `json:"issuer"`
	Deposit_id  int       `json:"deposit_id"`
	Amount      int       `json:"amount"`
	Description string    `json:"description"`
	Created_at  time.Time `json:"created_at,omitempty"`
}

func (w *Withdraw) Validate() error {
	if len(w.Description) > 100 {
		return errors.New("invalid description length")
	}
	return nil
}

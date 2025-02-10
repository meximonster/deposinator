package models

import (
	"errors"
	"time"
)

type Deposit struct {
	Id          int       `json:"id,omitempty"`
	Issuer      string    `json:"issuer"`
	Members     []int     `json:"members"`
	Amount      int       `json:"amount"`
	Description string    `json:"description"`
	Created_at  time.Time `json:"created_at,omitempty"`
}

func (d *Deposit) Validate() error {
	if len(d.Description) > 100 {
		return errors.New("invalid description length")
	}
	return nil
}

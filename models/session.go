package models

import (
	"errors"
	"time"
)

type Session struct {
	Id             int       `db:"id" json:"id,omitempty"`
	Issuer         int       `db:"issuer" json:"issuer"`
	Members        []int     `db:"members" json:"members"`
	Amount         int       `db:"amount" json:"amount"`
	WithdrawAmount int       `db:"withdraw_amount" json:"withdraw_amount"`
	Description    string    `db:"description" json:"description"`
	Created_at     time.Time `db:"created_at" json:"created_at,omitempty"`
}

func (s *Session) Validate() error {
	if len(s.Description) > 100 {
		return errors.New("invalid description length")
	}
	return nil
}

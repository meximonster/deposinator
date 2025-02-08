package models

import "time"

type Deposit struct {
	Id          int       `json:"id,omitempty"`
	Issuer      string    `json:"issuer"`
	Amount      int       `json:"amount"`
	Description string    `json:"description"`
	Created_at  time.Time `json:"created_at,omitempty"`
}

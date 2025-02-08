package models

import "time"

type Withdraw struct {
	Id          int       `json:"id,omitempty"`
	Issuer      string    `json:"issuer"`
	Deposit     int       `json:"deposit"`
	Amount      int       `json:"amount"`
	Description string    `json:"description"`
	Created_at  time.Time `json:"created_at,omitempty"`
}

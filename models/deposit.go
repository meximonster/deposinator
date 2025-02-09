package models

import "time"

type Deposit struct {
	Id          int
	Issuer      string
	Members     []int
	Amount      int
	Description string
	Created_at  time.Time
}

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

package serializers

import "time"

type Member struct {
	Id       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
}

type SessionSerializer struct {
	Id             int       `db:"id" json:"id,omitempty"`
	IssuerID       int       `db:"issuer_id" json:"issuer_id"`
	IssuerUsername string    `db:"issuer_username" json:"issuer_username"`
	Members        []Member  `db:"members" json:"members"`
	Amount         int       `db:"amount" json:"amount"`
	WithdrawAmount int       `db:"withdraw_amount" json:"withdraw_amount"`
	Description    string    `db:"description" json:"description"`
	Created_at     time.Time `db:"created_at" json:"created_at,omitempty"`
}

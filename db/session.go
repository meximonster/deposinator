package db

import (
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/deposinator/models"
	"github.com/deposinator/serializers"
)

func GetSessions(query string, args ...interface{}) ([]serializers.SessionSerializer, error) {
	sessions := []serializers.SessionSerializer{}
	rows, err := db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows
	for rows.Next() {
		var s serializers.SessionSerializer
		var membersStr string

		// Scan the row into the struct and the temporary members string
		err := rows.Scan(
			&s.Id,
			&s.IssuerID,
			&s.IssuerUsername,
			&membersStr, // Scan into string
			&s.Amount,
			&s.WithdrawAmount,
			&s.Description,
			&s.Created_at,
		)
		if err != nil {
			return nil, err
		}

		var members []serializers.Member
		err = json.Unmarshal([]byte(membersStr), &members)
		if err != nil {
			return nil, err
		}
		s.Members = members
		sessions = append(sessions, s)
	}

	return sessions, nil
}

func SessionCreate(issuer int, members []int, amount int, withdraw_amount int, description string) error {
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	var id int
	q := "INSERT INTO sessions (issuer, amount, withdraw_amount, description) VALUES ($1, $2, $3, $4) RETURNING id"
	err = tx.QueryRow(q, issuer, amount, withdraw_amount, description).Scan(&id)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = insertSessionMembers(tx, id, members)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func SessionUpdate(id int, issuer int, members []int, amount int, withdraw_amount int, description string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	q := "UPDATE sessions SET issuer = $1, amount = $2, withdraw_amount = $3, description = $4 WHERE id = $5"
	_, err = tx.Exec(q, issuer, amount, withdraw_amount, description, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec("DELETE FROM session_members where session_id = $1", id)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = insertSessionMembers(tx, id, members)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func SessionFromId(id int) (*models.Session, error) {
	var session models.Session
	err := db.Get(&session, "SELECT * FROM sessions WHERE id = $1", id)
	if err == sql.ErrNoRows {
		return nil, errors.New("session does not exist")
	}
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func SessionDelete(id int) error {
	q := "DELETE FROM sessions where id = $1"
	_, err := db.Exec(q, id)
	return err
}

func insertSessionMembers(tx *sql.Tx, session_id int, members []int) error {
	for _, member_id := range members {
		q := "INSERT INTO session_members (session_id, user_id) VALUES ($1, $2)"
		_, err := tx.Exec(q, session_id, member_id)
		if err != nil {
			return err
		}
	}
	return nil
}

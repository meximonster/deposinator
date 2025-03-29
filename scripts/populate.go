package main

import (
	"fmt"
	"log"

	"github.com/deposinator/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type SessionMember struct {
	SessionID int `db:"session_id"`
	UserID    int `db:"user_id"`
}

func main() {
	db, err := sqlx.Connect("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	tx := db.MustBegin()

	users := []models.User{
		{Username: "john_doe", Password: "password123", Email: "john.doe@example.com"},
		{Username: "jane_smith", Password: "password456", Email: "jane.smith@example.com"},
		{Username: "alice_jones", Password: "password789", Email: "alice.jones@example.com"},
		{Username: "bob_brown", Password: "password101", Email: "bob.brown@example.com"},
		{Username: "charlie_black", Password: "password102", Email: "charlie.black@example.com"},
		{Username: "david_white", Password: "password103", Email: "david.white@example.com"},
		{Username: "eve_green", Password: "password104", Email: "eve.green@example.com"},
		{Username: "frank_blue", Password: "password105", Email: "frank.blue@example.com"},
		{Username: "grace_red", Password: "password106", Email: "grace.red@example.com"},
		{Username: "henry_yellow", Password: "password107", Email: "henry.yellow@example.com"},
	}

	userIDs := make([]int, len(users))
	for i, user := range users {
		var id int
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		err = tx.QueryRowx(`INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id`, user.Username, string(hashedPassword), user.Email).Scan(&id)
		if err != nil {
			tx.Rollback()
			log.Fatalln(err)
		}
		userIDs[i] = id
	}

	sessions := []models.Session{
		{Issuer: userIDs[0], Amount: 100, WithdrawAmount: 50, Description: "Session 1"},
		{Issuer: userIDs[1], Amount: 200, WithdrawAmount: 100, Description: "Session 2"},
		{Issuer: userIDs[2], Amount: 300, WithdrawAmount: 150, Description: "Session 3"},
		{Issuer: userIDs[3], Amount: 400, WithdrawAmount: 200, Description: "Session 4"},
		{Issuer: userIDs[4], Amount: 500, WithdrawAmount: 250, Description: "Session 5"},
		{Issuer: userIDs[5], Amount: 600, WithdrawAmount: 300, Description: "Session 6"},
		{Issuer: userIDs[6], Amount: 700, WithdrawAmount: 350, Description: "Session 7"},
		{Issuer: userIDs[7], Amount: 800, WithdrawAmount: 400, Description: "Session 8"},
		{Issuer: userIDs[8], Amount: 900, WithdrawAmount: 450, Description: "Session 9"},
		{Issuer: userIDs[9], Amount: 1000, WithdrawAmount: 500, Description: "Session 10"},
	}

	sessionIDs := make([]int, len(sessions))
	for i, session := range sessions {
		var id int
		err := tx.QueryRowx(`INSERT INTO sessions (issuer, amount, withdraw_amount, description) VALUES ($1, $2, $3, $4) RETURNING id`, session.Issuer, session.Amount, session.WithdrawAmount, session.Description).Scan(&id)
		if err != nil {
			tx.Rollback()
			log.Fatalln(err)
		}
		sessionIDs[i] = id
	}

	sessionMembers := []SessionMember{
		{SessionID: sessionIDs[0], UserID: userIDs[0]},
		{SessionID: sessionIDs[0], UserID: userIDs[1]},
		{SessionID: sessionIDs[1], UserID: userIDs[2]},
		{SessionID: sessionIDs[1], UserID: userIDs[3]},
		{SessionID: sessionIDs[2], UserID: userIDs[4]},
		{SessionID: sessionIDs[2], UserID: userIDs[5]},
		{SessionID: sessionIDs[3], UserID: userIDs[6]},
		{SessionID: sessionIDs[3], UserID: userIDs[7]},
		{SessionID: sessionIDs[4], UserID: userIDs[8]},
		{SessionID: sessionIDs[4], UserID: userIDs[9]},
		{SessionID: sessionIDs[5], UserID: userIDs[0]},
		{SessionID: sessionIDs[5], UserID: userIDs[1]},
		{SessionID: sessionIDs[6], UserID: userIDs[2]},
		{SessionID: sessionIDs[6], UserID: userIDs[3]},
		{SessionID: sessionIDs[7], UserID: userIDs[4]},
		{SessionID: sessionIDs[7], UserID: userIDs[5]},
		{SessionID: sessionIDs[8], UserID: userIDs[6]},
		{SessionID: sessionIDs[8], UserID: userIDs[7]},
		{SessionID: sessionIDs[9], UserID: userIDs[8]},
		{SessionID: sessionIDs[9], UserID: userIDs[9]},
	}

	for _, sessionMember := range sessionMembers {
		_, err := tx.NamedExec(`INSERT INTO session_members (session_id, user_id) VALUES (:session_id, :user_id)`, &sessionMember)
		if err != nil {
			tx.Rollback()
			log.Fatalln(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Sample data inserted successfully")
}

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	ID        int       `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
}

type Deposit struct {
	ID          int       `db:"id"`
	Issuer      int       `db:"issuer"`
	Amount      int       `db:"amount"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}

type DepositMember struct {
	DepositID int `db:"deposit_id"`
	UserID    int `db:"user_id"`
}

type Withdrawal struct {
	ID          int       `db:"id"`
	Issuer      int       `db:"issuer"`
	DepositID   int       `db:"deposit_id"`
	Amount      int       `db:"amount"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}

func main() {
	// Connect to the database
	db, err := sqlx.Connect("postgres", "postgres://localhost:5432/postgres?sslmode=disable&user=postgres&password=postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Sample users
	users := []User{
		{Username: "alice", Password: "alice123", Email: "alice@example.com", CreatedAt: time.Now()},
		{Username: "bob", Password: "bob123", Email: "bob@example.com", CreatedAt: time.Now()},
		{Username: "charlie", Password: "charlie123", Email: "charlie@example.com", CreatedAt: time.Now()},
	}

	// Sample deposits
	deposits := []Deposit{
		{Issuer: 1, Amount: 1000, Description: "Initial deposit", CreatedAt: time.Now()},
		{Issuer: 2, Amount: 500, Description: "Savings deposit", CreatedAt: time.Now()},
		{Issuer: 3, Amount: 750, Description: "Investment deposit", CreatedAt: time.Now()},
	}

	// Sample withdrawals
	withdrawals := []Withdrawal{
		{Issuer: 1, DepositID: 1, Amount: 100, Description: "Withdrawal for expenses", CreatedAt: time.Now()},
		{Issuer: 2, DepositID: 2, Amount: 50, Description: "Withdrawal for savings", CreatedAt: time.Now()},
		{Issuer: 3, DepositID: 3, Amount: 75, Description: "Withdrawal for investment", CreatedAt: time.Now()},
	}

	// Insert users
	for _, user := range users {
		query := `
			INSERT INTO users (username, password, email, created_at)
			VALUES (:username, :password, :email, :created_at)
			RETURNING id
		`
		rows, err := db.NamedQuery(query, user)
		if err != nil {
			log.Fatalf("Failed to insert user: %v", err)
		}

		var userID int
		if rows.Next() {
			rows.Scan(&userID)
		}
		rows.Close()

		fmt.Printf("Inserted user with ID: %d\n", userID)
	}

	// Insert deposits
	for _, deposit := range deposits {
		query := `
				INSERT INTO deposits (issuer, amount, description, created_at)
				VALUES (:issuer, :amount, :description, :created_at)
				RETURNING id
			`
		rows, err := db.NamedQuery(query, deposit)
		if err != nil {
			log.Fatalf("Failed to insert deposit: %v", err)
		}

		var depositID int
		if rows.Next() {
			rows.Scan(&depositID)
		}
		rows.Close()

		fmt.Printf("Inserted deposit with ID: %d\n", depositID)

		// Add the issuer as a member of the deposit
		member := DepositMember{DepositID: depositID, UserID: deposit.Issuer}
		_, err = db.NamedExec(`
				INSERT INTO deposit_members (deposit_id, user_id)
				VALUES (:deposit_id, :user_id)
			`, member)
		if err != nil {
			log.Fatalf("Failed to insert deposit member: %v", err)
		}
	}

	// Insert withdrawals
	for _, withdrawal := range withdrawals {
		query := `
			INSERT INTO withdraws (issuer, deposit_id, amount, description, created_at)
			VALUES (:issuer, :deposit_id, :amount, :description, :created_at)
			RETURNING id
		`
		rows, err := db.NamedQuery(query, withdrawal)
		if err != nil {
			log.Fatalf("Failed to insert withdrawal: %v", err)
		}

		var withdrawalID int
		if rows.Next() {
			rows.Scan(&withdrawalID)
		}
		rows.Close()

		fmt.Printf("Inserted withdrawal with ID: %d\n", withdrawalID)
	}

}

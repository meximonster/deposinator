package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	Username string
	Password string
	Email    string
}

func main() {
	db, err := sqlx.Connect("postgres", "postgres://localhost:5432/postgres?sslmode=disable&user=postgres&password=postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	sampleUsers := []User{
		{Username: "john_doe", Password: "password123", Email: "john@example.com"},
		{Username: "jane_doe", Password: "password456", Email: "jane@example.com"},
		{Username: "alice", Password: "password789", Email: "alice@example.com"},
		{Username: "bob", Password: "password101", Email: "bob@example.com"},
	}

	for _, user := range sampleUsers {
		id, err := insertUser(db, user)
		if err != nil {
			log.Printf("Failed to insert user %s: %v\n", user.Username, err)
			continue
		}
		fmt.Printf("Inserted user %s with ID: %d\n", user.Username, id)
	}

}

func insertUser(db *sqlx.DB, user User) (int, error) {
	q := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err := db.QueryRow(q, user.Username, user.Password, user.Email).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

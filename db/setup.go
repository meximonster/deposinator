package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func NewDB(host string, user string, password string) (*sql.DB, error) {
	maxRetries := 10
	retryDelay := 5 * time.Second

	var err error
	for i := 0; i < maxRetries; i++ {
		database, connectErr := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s/postgres?sslmode=disable&user=%s&password=%s", host, user, password))
		if connectErr == nil {
			// Connection successful
			database.SetMaxOpenConns(25)
			database.SetMaxIdleConns(25)
			database.SetConnMaxLifetime(5 * time.Minute)

			db = database
			return database.DB, nil
		}

		// Log the error and retry
		err = connectErr
		log.Printf("Failed to connect to database (attempt %d/%d): %v\n", i+1, maxRetries, connectErr)
		time.Sleep(retryDelay)
	}

	// If all retries fail, return the last error
	return nil, fmt.Errorf("failed to connect to database after %d retries: %w", maxRetries, err)
}

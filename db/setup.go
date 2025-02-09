package db

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func NewDB(host string, user string, password string) error {
	database, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s/postgres?sslmode=disable&user=%s&password=%s", host, user, password))
	if err != nil {
		return err
	}
	database.SetMaxOpenConns(25)
	database.SetMaxIdleConns(25)
	database.SetConnMaxLifetime(5 * time.Minute)

	db = database
	return nil
}

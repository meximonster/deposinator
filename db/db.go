package db

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func NewDB(host string, user string, password string) error {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s/postgres?sslmode=disable&user=%s&password=%s&timezone=Europe/Athens", host, user, password))
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	return nil
}

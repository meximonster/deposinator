package models

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64  `json:"id,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserExists(email string) (bool, error) {
	var user User
	err := db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	// user does not exist
	if err == sql.ErrNoRows {
		return false, nil
	}
	// error getting user
	if err != nil {
		return false, err
	}
	// user exists
	return true, nil
}

func UserCreate(username string, email string, password string) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	q := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err = db.QueryRow(q, username, email, hashedPassword).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func MatchUserPassword(email string, password string) *User {
	var user User
	err := db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err == sql.ErrNoRows {
		return &User{}
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return &User{}
	}
	return &user
}

func UserFromId(id int) *User {
	var user User
	err := db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return &User{}
	}
	return &user
}

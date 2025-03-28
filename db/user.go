package db

import (
	"database/sql"

	"github.com/deposinator/models"
	"golang.org/x/crypto/bcrypt"
)

func UserExists(username string, email string) (bool, error) {
	var user models.User
	err := db.Get(&user, "SELECT * FROM users WHERE username = $1 OR email = $2", username, email)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func UserCreate(username string, email string, password string) (int, error) {
	q := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err := db.QueryRow(q, username, email, password).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func MatchUserPassword(email string, password string) *models.User {
	var user models.User
	err := db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err == sql.ErrNoRows {
		return &models.User{}
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return &models.User{}
	}
	return &user
}

func UserFromId(id int) *models.User {
	var user models.User
	err := db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return &models.User{}
	}
	return &user
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	err := db.Select(&users, "SELECT id, username FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

package models

import (
	"errors"
	"time"
)

type User struct {
	Id         int       `json:"id,omitempty"`
	Username   string    `json:"username,omitempty"`
	Email      string    `json:"email" binding:"required,email"`
	Password   string    `json:"password" binding:"required"`
	Created_at time.Time `json:"-"`
}

func (u *User) Validate() error {
	if u.Username == "" || len(u.Username) > 50 {
		return errors.New("invalid username length")
	}
	if u.Password == "" || len(u.Password) < 5 {
		return errors.New("invalid password length")
	}
	return nil
}

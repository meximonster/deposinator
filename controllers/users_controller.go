package controllers

import (
	"log"
	"net/http"

	"github.com/deposinator/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type userData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(c *gin.Context) {
	var data userData
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	exists, err := db.UserExists(data.Username, data.Email)
	if err != nil {
		log.Println("error checking user status: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if exists {
		log.Println("user exists: ", data.Username)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userID, err := db.UserCreate(data.Username, data.Email, data.Password)
	if err != nil {
		log.Printf("error creating user %s: %s", data.Username, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	session := sessions.Default(c)
	session.Set("userID", userID)
	session.Save()
	c.Status(http.StatusOK)
}

func Login(c *gin.Context) {
	var data userData
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := db.MatchUserPassword(data.Email, data.Password)
	if user.Id == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	session := sessions.Default(c)
	session.Set("userID", user.Id)
	session.Save()
	c.Status(http.StatusOK)
}

func Logout(c *gin.Context) {
	// Delete the session
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Status(http.StatusOK)
}

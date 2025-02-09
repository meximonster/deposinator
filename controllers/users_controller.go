package controllers

import (
	"log"
	"net/http"

	"github.com/deposinator/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type userData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(c *gin.Context) {
	var data userData
	c.Bind(&data)

	exists, err := db.UserExists(data.Username, data.Email)
	if err != nil {
		log.Println("error checking user status: ", err)
		c.Render(http.StatusInternalServerError, render.Data{})
		return
	}

	if exists {
		log.Println("user exists: ", data.Username)
		c.Render(http.StatusBadRequest, render.Data{})
		return
	}

	userID, err := db.UserCreate(data.Username, data.Email, data.Password)
	if err != nil {
		log.Printf("error creating user %s: %s", data.Username, err.Error())
		c.Render(http.StatusInternalServerError, render.Data{})
		return
	}

	session := sessions.Default(c)
	session.Set("userID", userID)
	session.Save()
	c.Status(http.StatusOK)
}

func Login(c *gin.Context) {
	var data userData
	c.Bind(&data)

	user := db.MatchUserPassword(data.Email, data.Password)
	if user.Id == 0 {
		c.Render(http.StatusUnauthorized, render.Data{})
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

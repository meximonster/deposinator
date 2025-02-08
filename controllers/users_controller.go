package controllers

import (
	"net/http"

	"github.com/deposinator/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type formData struct {
	Username string `form:"username"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func Signup(c *gin.Context) {

	var data formData
	c.Bind(&data)

	exists, err := models.UserExists(data.Email)
	if err != nil {
		c.Render(http.StatusInternalServerError, render.Data{})
		return
	}

	if exists {
		c.Render(http.StatusBadRequest, render.Data{})
		return
	}

	userID, err := models.UserCreate(data.Username, data.Email, data.Password)
	if err != nil {
		c.Render(http.StatusInternalServerError, render.Data{})
		return
	}

	session := sessions.Default(c)
	session.Set("userID", userID)
	session.Save()
	c.Status(http.StatusAccepted)

}

func Login(c *gin.Context) {

	var data formData
	c.Bind(&data)

	user := models.MatchUserPassword(data.Email, data.Password)
	if user.Id == 0 {
		c.Render(http.StatusUnauthorized, render.Data{})
		return
	}

	session := sessions.Default(c)
	session.Set("userID", user.Id)
	session.Save()
	c.Status(http.StatusAccepted)

}

func Logout(c *gin.Context) {

	// Delete the session
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Status(http.StatusAccepted)

}

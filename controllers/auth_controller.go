package controllers

import (
	"log"
	"net/http"

	"github.com/deposinator/db"
	"github.com/deposinator/models"
	"github.com/deposinator/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Println("error binding user: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateJSONResponse("error", err.Error()))
		return
	}

	err = user.Validate()
	if err != nil {
		log.Println("error validating user: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateJSONResponse("error", err.Error()))
		return
	}

	exists, err := db.UserExists(user.Username, user.Email)
	if err != nil {
		log.Println("error checking user status: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateJSONResponse("error", err.Error()))
		return
	}

	if exists {
		log.Println("user exists: ", user.Username)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateJSONResponse("error", "user already exists"))
		return
	}

	user.Id, err = db.UserCreate(user.Username, user.Email, user.Password)
	if err != nil {
		log.Printf("error creating user %s: %s", user.Username, err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateJSONResponse("error", err.Error()))
		return
	}

	session := sessions.Default(c)
	session.Set("userID", user.Id)
	session.Save()
	c.JSON(http.StatusOK, utils.GenerateJSONResultResponse("success", "OK", user))
}

func Login(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Println("error binding user: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateJSONResponse("error", err.Error()))
		return
	}

	u := db.MatchUserPassword(user.Email, user.Password)
	if u.Id == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.GenerateJSONResponse("error", "user not found"))
		return
	}

	session := sessions.Default(c)
	session.Set("userID", u.Id)
	session.Save()
	c.JSON(http.StatusOK, utils.GenerateJSONResultResponse("success", "OK", u))
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, utils.GenerateJSONResponse("success", "OK"))
}

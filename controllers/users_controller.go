package controllers

import (
	"net/http"

	"github.com/deposinator/db"
	"github.com/deposinator/utils"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := db.GetUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateJSONResponse("error", err.Error()))
	}
	c.JSON(http.StatusOK, utils.GenerateJSONResultResponse("success", "OK", users))
}

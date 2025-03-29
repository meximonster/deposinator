package controllers

import (
	"net/http"
	"strconv"

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

func GetUserDashboard(c *gin.Context) {
	id := c.Param("user_id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateJSONResponse("error", "User ID is required"))
		return
	}
	// Convert userID to int
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateJSONResponse("error", "Invalid User ID"))
		return
	}
	// Fetch dashboard data
	dashboardData, err := db.GetUserDashboard(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateJSONResponse("error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.GenerateJSONResultResponse("success", "OK", dashboardData))
}

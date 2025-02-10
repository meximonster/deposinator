package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/deposinator/db"
	"github.com/deposinator/models"
	"github.com/deposinator/utils"
	"github.com/gin-gonic/gin"
)

func WithdrawCreate(c *gin.Context) {
	var withdraw *models.Withdraw
	if err := c.BindJSON(&withdraw); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateJSONResponse("error", err.Error()))
		return
	}

	err := db.WithdrawCreate(withdraw.Issuer, withdraw.Deposit_id, withdraw.Amount, withdraw.Description)
	if err != nil {
		log.Println("error creating withdraw: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateJSONResponse("error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.GenerateJSONResponse("success", "OK"))
}

func WithdrawUpdate(c *gin.Context) {
	var withdraw *models.Withdraw
	if err := c.BindJSON(&withdraw); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateJSONResponse("error", err.Error()))
		return
	}

	err := db.WithdrawUpdate(withdraw.Id, withdraw.Issuer, withdraw.Deposit_id, withdraw.Amount, withdraw.Description)
	if err != nil {
		log.Println("error updating withdraw: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateJSONResponse("error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.GenerateJSONResponse("success", "OK"))
}

func WithdrawDelete(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		log.Println("id parameter not found")
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateJSONResponse("error", "missing id parameter"))
		return
	}
	deposit_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println("invalid deposit id: ", id)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateJSONResponse("error", err.Error()))
		return
	}
	err = db.WithdrawDelete(deposit_id)
	if err != nil {
		log.Println("error deleting withdraw: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateJSONResponse("error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.GenerateJSONResponse("success", "OK"))
}

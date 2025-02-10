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

func DepositCreate(c *gin.Context) {
	var deposit *models.Deposit
	if err := c.BindJSON(&deposit); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateJSONResponse("error", err.Error()))
		return
	}

	err := db.DepositCreate(deposit.Issuer, deposit.Members, deposit.Amount, deposit.Description)
	if err != nil {
		log.Println("error creating deposit: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateJSONResponse("error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.GenerateJSONResponse("success", "OK"))
}

func DepositUpdate(c *gin.Context) {
	var deposit *models.Deposit
	if err := c.BindJSON(&deposit); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateJSONResponse("error", err.Error()))
		return
	}

	err := db.DepositUpdate(deposit.Id, deposit.Issuer, deposit.Members, deposit.Amount, deposit.Description)
	if err != nil {
		log.Println("error updating deposit: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateJSONResponse("error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.GenerateJSONResponse("success", "OK"))
}

func DepositDelete(c *gin.Context) {
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
	err = db.DepositDelete(deposit_id)
	if err != nil {
		log.Println("error deleting deposit: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateJSONResponse("error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.GenerateJSONResponse("success", "OK"))
}

package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/deposinator/db"
	"github.com/gin-gonic/gin"
)

type withdrawData struct {
	Id          int       `json:"id,omitempty"`
	Issuer      string    `form:"issuer"`
	Deposit_id  int       `json:"deposit_id"`
	Amount      int       `form:"amount"`
	Description string    `form:"description"`
	Created_at  time.Time `json:"created_at,omitempty"`
}

func WithdrawCreate(c *gin.Context) {
	var data withdrawData
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := db.WithdrawCreate(data.Issuer, data.Deposit_id, data.Amount, data.Description)
	if err != nil {
		log.Println("error creating withdraw: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func WithdrawUpdate(c *gin.Context) {
	var data withdrawData
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := db.WithdrawUpdate(data.Id, data.Issuer, data.Deposit_id, data.Amount, data.Description)
	if err != nil {
		log.Println("error updating withdraw: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func WithdrawDelete(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		log.Println("id parameter not found")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	deposit_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println("invalid deposit id: ", id)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = db.WithdrawDelete(deposit_id)
	if err != nil {
		log.Println("error deleting withdraw: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

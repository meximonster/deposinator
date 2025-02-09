package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/deposinator/db"
	"github.com/gin-gonic/gin"
)

type depositData struct {
	Id          int       `json:"id,omitempty"`
	Issuer      string    `json:"issuer"`
	Members     []int     `json:"members"`
	Amount      int       `json:"amount"`
	Description string    `json:"description"`
	Created_at  time.Time `json:"created_at,omitempty"`
}

func DepositCreate(c *gin.Context) {
	var data depositData
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := db.DepositCreate(data.Issuer, data.Members, data.Amount, data.Description)
	if err != nil {
		log.Println("error creating deposit: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func DepositUpdate(c *gin.Context) {
	var data depositData
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := db.DepositUpdate(data.Id, data.Issuer, data.Members, data.Amount, data.Description)
	if err != nil {
		log.Println("error updating deposit: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func DepositDelete(c *gin.Context) {
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
	err = db.DepositDelete(deposit_id)
	if err != nil {
		log.Println("error deleting deposit: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

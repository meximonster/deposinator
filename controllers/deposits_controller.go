package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/deposinator/db"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
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
	c.Bind(&data)

	err := db.DepositCreate(data.Issuer, data.Members, data.Amount, data.Description)
	if err != nil {
		log.Println("error creating deposit: ", err)
		c.Render(http.StatusInternalServerError, render.Data{})
		return
	}
	c.Status(http.StatusOK)
}

func DepositUpdate(c *gin.Context) {
	var data depositData
	c.Bind(&data)
	err := db.DepositUpdate(data.Id, data.Issuer, data.Members, data.Amount, data.Description)
	if err != nil {
		log.Println("error updating deposit: ", err)
		c.Render(http.StatusInternalServerError, render.Data{})
		return
	}
	c.Status(http.StatusOK)
}

func DepositDelete(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		log.Println("id parameter not found")
		c.Render(http.StatusBadRequest, render.Data{})
		return
	}
	deposit_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println("invalid deposit id: ", id)
		c.Render(http.StatusBadRequest, render.Data{})
		return
	}
	err = db.DepositDelete(deposit_id)
	if err != nil {
		log.Println("error deleting deposit: ", err)
		c.Render(http.StatusInternalServerError, render.Data{})
		return
	}
	c.Status(http.StatusOK)
}

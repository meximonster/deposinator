package controllers

import (
	"github.com/gin-gonic/gin"
)

type depositData struct {
	Issuer      string `form:"issuer"`
	Amount      int    `form:"amount"`
	Description string `form:"description"`
}

func DepositCreate(c *gin.Context) {}

func DepositUpdate(c *gin.Context) {}

func DepositDelete(c *gin.Context) {}

package controllers

import "github.com/gin-gonic/gin"

type withdrawData struct {
	Issuer      string `form:"issuer"`
	Amount      int    `form:"amount"`
	Description string `form:"description"`
}

func WithdrawCreate(c *gin.Context) {}

func WithdrawUpdate(c *gin.Context) {}

func WithdrawDelete(c *gin.Context) {}

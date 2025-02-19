package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/deposinator/db"
	"github.com/deposinator/models"
	"github.com/deposinator/utils"
	"github.com/gin-gonic/gin"
)

func GetWithdrawals(c *gin.Context) {
	// Parse query parameters
	issuer := c.Query("issuer")
	depositID := c.Query("deposit_id")
	minAmount := c.Query("min_amount")
	maxAmount := c.Query("max_amount")
	description := c.Query("description")
	createdAfter := c.Query("created_after")
	createdBefore := c.Query("created_before")
	sortBy := c.DefaultQuery("sort_by", "created_at")
	sortOrder := c.DefaultQuery("sort_order", "desc")
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")

	query := `
		SELECT id, issuer, deposit_id, amount, description, created_at
		FROM withdraws
		WHERE 1=1
	`
	var args []interface{}
	argIndex := 1

	if issuer != "" {
		query += fmt.Sprintf(" AND issuer = $%d", argIndex)
		args = append(args, issuer)
		argIndex++
	}
	if depositID != "" {
		query += fmt.Sprintf(" AND deposit_id = $%d", argIndex)
		args = append(args, depositID)
		argIndex++
	}
	if minAmount != "" {
		query += fmt.Sprintf(" AND amount >= $%d", argIndex)
		args = append(args, minAmount)
		argIndex++
	}
	if maxAmount != "" {
		query += fmt.Sprintf(" AND amount <= $%d", argIndex)
		args = append(args, maxAmount)
		argIndex++
	}
	if description != "" {
		query += fmt.Sprintf(" AND description ILIKE $%d", argIndex)
		args = append(args, "%"+description+"%")
		argIndex++
	}
	if createdAfter != "" {
		query += fmt.Sprintf(" AND created_at >= $%d", argIndex)
		args = append(args, createdAfter)
		argIndex++
	}
	if createdBefore != "" {
		query += fmt.Sprintf(" AND created_at <= $%d", argIndex)
		args = append(args, createdBefore)
		argIndex++
	}

	query += fmt.Sprintf(" ORDER BY %s %s LIMIT $%d OFFSET $%d", sortBy, sortOrder, argIndex, argIndex+1)
	args = append(args, limit, offset)

	withdraws, err := db.GetWithdrawals(query, args...)
	if err != nil {
		log.Printf("error getting deposits. query: %s, error %s\n", query, err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateJSONResponse("error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.GenerateJSONResultResponse("success", "OK", withdraws))
}

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
	id := c.Param("id")
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

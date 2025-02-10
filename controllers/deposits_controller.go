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

func GetDeposits(c *gin.Context) {
	// Parse query parameters
	issuer := c.Query("issuer")
	member := c.Query("member")
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
		SELECT d.id, d.issuer, d.amount, d.description, d.created_at
		FROM deposits d
		LEFT JOIN deposit_members dm ON d.id = dm.deposit_id
		WHERE 1=1
	`
	var args []interface{}
	argIndex := 1

	if issuer != "" {
		query += fmt.Sprintf(" AND d.issuer = $%d", argIndex)
		args = append(args, issuer)
		argIndex++
	}
	if member != "" {
		query += fmt.Sprintf(" AND dm.user_id = $%d", argIndex)
		args = append(args, member)
		argIndex++
	}
	if minAmount != "" {
		query += fmt.Sprintf(" AND d.amount >= $%d", argIndex)
		args = append(args, minAmount)
		argIndex++
	}
	if maxAmount != "" {
		query += fmt.Sprintf(" AND d.amount <= $%d", argIndex)
		args = append(args, maxAmount)
		argIndex++
	}
	if description != "" {
		query += fmt.Sprintf(" AND d.description ILIKE $%d", argIndex)
		args = append(args, "%"+description+"%")
		argIndex++
	}
	if createdAfter != "" {
		query += fmt.Sprintf(" AND d.created_at >= $%d", argIndex)
		args = append(args, createdAfter)
		argIndex++
	}
	if createdBefore != "" {
		query += fmt.Sprintf(" AND d.created_at <= $%d", argIndex)
		args = append(args, createdBefore)
		argIndex++
	}

	query += fmt.Sprintf(" ORDER BY %s %s LIMIT $%d OFFSET $%d", sortBy, sortOrder, argIndex, argIndex+1)
	args = append(args, limit, offset)

	deposits, err := db.GetDeposits(query, args...)
	if err != nil {
		log.Printf("error getting deposits. query: %s, error %s\n", query, err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateJSONResponse("error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.GenerateJSONResultResponse("success", "OK", deposits))
}

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
	err = db.DepositDelete(deposit_id)
	if err != nil {
		log.Println("error deleting deposit: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateJSONResponse("error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.GenerateJSONResponse("success", "OK"))
}

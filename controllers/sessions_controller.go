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

func GetSessions(c *gin.Context) {
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
		SELECT 
			s.id, 
			s.issuer, 
			COALESCE(ARRAY_AGG(sm.user_id), '{}') AS members,
			s.amount,
			s.withdraw_amount,
			s.description, 
			s.created_at
		FROM 
			sessions s
		LEFT JOIN 
			session_members sm 
		ON 
			s.id = sm.session_id
		GROUP BY 
			s.id, s.issuer, s.amount, s.description, s.created_at
	`

	var args []interface{}
	argIndex := 1

	if issuer != "" {
		query += fmt.Sprintf(" AND s.issuer = $%d", argIndex)
		args = append(args, issuer)
		argIndex++
	}
	if member != "" {
		query += fmt.Sprintf(" AND sm.user_id = $%d", argIndex)
		args = append(args, member)
		argIndex++
	}
	if minAmount != "" {
		query += fmt.Sprintf(" AND s.amount >= $%d", argIndex)
		args = append(args, minAmount)
		argIndex++
	}
	if maxAmount != "" {
		query += fmt.Sprintf(" AND s.amount <= $%d", argIndex)
		args = append(args, maxAmount)
		argIndex++
	}
	if description != "" {
		query += fmt.Sprintf(" AND s.description ILIKE $%d", argIndex)
		args = append(args, "%"+description+"%")
		argIndex++
	}
	if createdAfter != "" {
		query += fmt.Sprintf(" AND s.created_at >= $%d", argIndex)
		args = append(args, createdAfter)
		argIndex++
	}
	if createdBefore != "" {
		query += fmt.Sprintf(" AND s.created_at <= $%d", argIndex)
		args = append(args, createdBefore)
		argIndex++
	}

	query += fmt.Sprintf(" ORDER BY %s %s LIMIT $%d OFFSET $%d", sortBy, sortOrder, argIndex, argIndex+1)
	args = append(args, limit, offset)

	sessions, err := db.GetSessions(query, args...)
	if err != nil {
		log.Printf("error getting sessions. query: %s, error %s\n", query, err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateJSONResponse("error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.GenerateJSONResultResponse("success", "OK", sessions))
}

func SessionCreate(c *gin.Context) {
	var session *models.Session
	if err := c.BindJSON(&session); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateJSONResponse("error", err.Error()))
		return
	}

	err := db.SessionCreate(session.Issuer, session.Members, session.Amount, session.WithdrawAmount, session.Description)
	if err != nil {
		log.Println("error creating session: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateJSONResponse("error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.GenerateJSONResponse("success", "OK"))
}

func SessionUpdate(c *gin.Context) {
	var session *models.Session
	if err := c.BindJSON(&session); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateJSONResponse("error", err.Error()))
		return
	}

	err := db.SessionUpdate(session.Id, session.Issuer, session.Members, session.Amount, session.WithdrawAmount, session.Description)
	if err != nil {
		log.Println("error updating session: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateJSONResponse("error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.GenerateJSONResponse("success", "OK"))
}

func SessionDelete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		log.Println("id parameter not found")
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateJSONResponse("error", "missing id parameter"))
		return
	}
	session_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println("invalid session id: ", id)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateJSONResponse("error", err.Error()))
		return
	}
	err = db.SessionDelete(session_id)
	if err != nil {
		log.Println("error deleting session: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateJSONResponse("error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.GenerateJSONResponse("success", "OK"))
}

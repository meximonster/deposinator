package server

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var tokens []string

func Run(accounts map[string]string) error {
	r := gin.Default()
	r.POST("/login", gin.BasicAuth(accounts), func(c *gin.Context) {
		token, _ := randomHex(20)
		tokens = append(tokens, token)

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	})
	r.GET("/resource", func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		reqToken := strings.Split(bearerToken, " ")[1]
		for _, token := range tokens {
			if token == reqToken {
				c.JSON(http.StatusOK, gin.H{
					"data": "resource data",
				})
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
	})
	if err := r.Run(); err != nil {
		return err
	}
	return nil
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

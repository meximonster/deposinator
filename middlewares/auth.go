package middlewares

import (
	"net/http"

	"github.com/deposinator/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		session := sessions.Default(c)
		sessionID := session.Get("userID")
		if sessionID == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		userId := sessionID.(int)
		// Check if the user exists
		user := db.UserFromId(userId)
		if user.Id == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("userID", user.Id)

		c.Next()
	}
}

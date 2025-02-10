package server

import (
	"net/http"
	"time"

	"github.com/deposinator/controllers"
	"github.com/deposinator/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func Run(env string, port string, storeKey string) {

	r := gin.Default()

	store := memstore.NewStore([]byte(storeKey))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   21600,
		Secure:   false, // Set to true in production with HTTPS
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	r.Use(sessions.Sessions("deposinator", store))
	if env == "dev" {
		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"}, // Frontend URL
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			AllowCredentials: true,           // Allow cookies and credentials
			MaxAge:           12 * time.Hour, // Cache preflight requests
		}))
	}

	users := r.Group("/users")
	{
		users.POST("/signup", controllers.Signup)
		users.POST("/login", controllers.Login)
		users.DELETE("/logout", controllers.Logout)
	}

	deposits := r.Group("/deposit")
	{
		deposits.POST("/", middlewares.AuthMiddleware(), controllers.DepositCreate)
		deposits.PUT("/", middlewares.AuthMiddleware(), controllers.DepositUpdate)
		deposits.DELETE("/", middlewares.AuthMiddleware(), controllers.DepositDelete)
	}

	withdraws := r.Group("/withdraw")
	{
		withdraws.POST("/", middlewares.AuthMiddleware(), controllers.WithdrawCreate)
		withdraws.PUT("/", middlewares.AuthMiddleware(), controllers.WithdrawUpdate)
		withdraws.DELETE("/", middlewares.AuthMiddleware(), controllers.WithdrawDelete)
	}

	r.Run(":" + port)
}

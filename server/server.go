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
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	r.Use(sessions.Sessions("deposinator", store))
	if env == "dev" {
		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	}

	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.DELETE("/logout", controllers.Logout)
	}

	users := r.Group("/users")
	{
		users.GET("", middlewares.AuthMiddleware(), controllers.GetUsers)
	}

	deposits := r.Group("/deposits")
	{
		deposits.GET("", middlewares.AuthMiddleware(), controllers.GetDeposits)
		deposits.POST("", middlewares.AuthMiddleware(), controllers.DepositCreate)
		deposits.PUT("", middlewares.AuthMiddleware(), controllers.DepositUpdate)
		deposits.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DepositDelete)
	}

	withdrawals := r.Group("/withdrawals")
	{
		withdrawals.GET("", middlewares.AuthMiddleware(), controllers.GetWithdrawals)
		withdrawals.POST("", middlewares.AuthMiddleware(), controllers.WithdrawCreate)
		withdrawals.PUT("", middlewares.AuthMiddleware(), controllers.WithdrawUpdate)
		withdrawals.DELETE("/:id", middlewares.AuthMiddleware(), controllers.WithdrawDelete)
	}

	r.Static("/swagger-ui", "./swagger-ui")
	r.StaticFile("/docs/swagger.yml", "./docs/swagger.yml")

	r.Run(":5000")
}

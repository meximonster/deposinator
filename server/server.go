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
		deposits.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DepositDelete)
	}
	r.GET("/deposits", middlewares.AuthMiddleware(), controllers.GetDeposits)

	withdraws := r.Group("/withdraw")
	{
		withdraws.GET("/", middlewares.AuthMiddleware(), controllers.GetWithdraws)
		withdraws.POST("/", middlewares.AuthMiddleware(), controllers.WithdrawCreate)
		withdraws.PUT("/", middlewares.AuthMiddleware(), controllers.WithdrawUpdate)
		withdraws.DELETE("/:id", middlewares.AuthMiddleware(), controllers.WithdrawDelete)
	}
	r.GET("/withdraws", middlewares.AuthMiddleware(), controllers.GetWithdraws)

	r.Static("/swagger-ui", "./swagger-ui")
	r.StaticFile("/docs/swagger.yml", "./docs/swagger.yml")

	r.Run(":" + port)
}

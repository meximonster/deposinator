package server

import (
	"github.com/deposinator/controllers"
	"github.com/deposinator/middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func Run(port string, storeKey string) {

	r := gin.Default()

	store := memstore.NewStore([]byte(storeKey))
	store.Options(sessions.Options{
		MaxAge: 21600,
	})
	r.Use(sessions.Sessions("deposinator", store))

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

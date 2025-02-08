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

	r.POST("/deposit", middlewares.AuthMiddleware(), controllers.Deposit)
	r.POST("/withdraw", middlewares.AuthMiddleware(), controllers.Withdraw)

	r.Run(":" + port)
}

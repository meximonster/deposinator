package server

import (
	"github.com/deposinator/controllers"
	"github.com/gin-gonic/gin"
)

func Run() error {

	r := gin.Default()
	r.Use(gin.Logger())
	r.POST("/users/signup", controllers.Signup)
	r.POST("/users/login", controllers.Login)
	r.DELETE("/users/logout", controllers.Logout)

	if err := r.Run(":5000"); err != nil {
		return err
	}
	return nil
}

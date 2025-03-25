package server

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/deposinator/controllers"
	"github.com/deposinator/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/gin"
)

func Run(db *sql.DB, env string, port string, storeKey string) error {

	r := gin.Default()

	store, err := postgres.NewStore(db, []byte(storeKey))
	if err != nil {
		return err
	}
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

	sessions := r.Group("/sessions")
	{
		sessions.GET("", middlewares.AuthMiddleware(), controllers.GetSessions)
		sessions.GET("/:id", middlewares.AuthMiddleware(), controllers.SessionById)
		sessions.POST("", middlewares.AuthMiddleware(), controllers.SessionCreate)
		sessions.PUT("", middlewares.AuthMiddleware(), controllers.SessionUpdate)
		sessions.DELETE("/:id", middlewares.AuthMiddleware(), controllers.SessionDelete)
	}

	r.Static("/swagger-ui", "./swagger-ui")
	r.StaticFile("/docs/swagger.yml", "./docs/swagger.yml")

	return r.Run(":5000")
}

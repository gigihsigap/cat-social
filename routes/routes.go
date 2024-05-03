package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"

	// "cat-social/auth"
	// "cat-social/repositories"

	_ "github.com/lib/pq"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	// Version 1 router group
	// v1 := r.Group("/v1")

	// "v1" group
	// v1 := r.Group("/v1")

	// "user" group
	// userRepo := repositories.NewUserRepository(db)
	// authHandler := auth.NewHandler(userRepo)

	// user := v1.Group("/user")
	// user.POST("/register", authHandler.Register)
	// user.POST("/login", authHandler.Login)

	return r
}

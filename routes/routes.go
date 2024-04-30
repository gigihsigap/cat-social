package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"

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

	return r
}

// backend/main.go
package main

import (
	"github.com/gin-gonic/gin"
	"yatte-mi/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.POST("/plan", handlers.HandlePlan)

	router.Run(":8080")
}

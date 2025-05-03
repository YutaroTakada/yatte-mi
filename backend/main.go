// backend/main.go
package main

import (
	"main/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.POST("/plan", handlers.HandlePlan)

	router.Run(":8080")
}

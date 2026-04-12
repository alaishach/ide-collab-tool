package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")

	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"message": "healthy",
		})
	})

	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}

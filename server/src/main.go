package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server/src/db"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")

	api.GET("/health", func(c *gin.Context) {
		if err := pg.DB.Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Database failed to init: " + err.Error(),
			})
			panic("Failed to init db")
		}
		c.JSON(http.StatusOK, map[string]string{
			"message": "healthy",
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: " + err.Error())
		panic("Failed to start server: " + err.Error())
	}
}

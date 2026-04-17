package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server/src/db/pg"
	"server/src/db/red"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")

	api.GET("/health", func(c *gin.Context) {
		// Postgres health check
		if err := pg.DB.Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "PostgreSql Database failed to init: " + err.Error(),
			})
			panic("Failed to init db")
		}
		// Redis health check
		if err := red.Client.Ping(red.Ctx).Err(); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Redis failed to init: " + err.Error(),
			})
			panic("Failed to init db")
		}
		// Success
		c.JSON(http.StatusOK, map[string]string{
			"message": "Success",
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: " + err.Error())
		panic("Failed to start server: " + err.Error())
	}
}

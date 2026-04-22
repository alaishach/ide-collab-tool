// Package gin
package gin

import (
	"github.com/gin-gonic/gin"
	"log"
	"server/internal/api/handlers/auth"
	"server/internal/api/handlers/health"
)

var Gin *gin.Engine

func checkRunning() {
	if err := Gin.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: " + err.Error())
		panic("Failed to start server: " + err.Error())
	}
}

func Run() {
	Gin = gin.Default()
	api := Gin.Group("/api")

	// Health
	api.GET("/health", health.Health)

	// Auth
	api.POST("/signup", auth.Signup)
	api.POST("/login", auth.PostLogin)

	checkRunning()
}

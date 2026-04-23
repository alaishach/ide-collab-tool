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

	// TODO rename
	// Auth
	api.POST("/users", auth.Signup)
	api.POST("/sessions", auth.PostLogin)
	api.PATCH("/sessions", auth.GetLogin)
	api.DELETE("/sessions", auth.Logout)

	checkRunning()
}

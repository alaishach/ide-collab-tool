// Package main
package main

import (
	"net/http"
	"server/internal/api/handlers/auth"
	"server/internal/api/handlers/health"
	"server/internal/consts"
	"server/internal/utils/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run() {
	api := chi.NewRouter()
	api.Use(middleware.Logger)

	// Health
	api.Get("/api/health", health.Health)

	// Auth
	api.Post("/api/users", auth.Signup)
	api.Post("/api/sessions", auth.PostLogin)
	api.Patch("/api/sessions", auth.GetLogin)
	api.Delete("/api/sessions", auth.Logout)

	logger.Logger.Info("Server init successful if no error")
	if err := http.ListenAndServe(":"+consts.SERVER_PORT, api); err != nil {
		panic("Server Error: " + err.Error())
	}
}

func main() {
	Run()
}

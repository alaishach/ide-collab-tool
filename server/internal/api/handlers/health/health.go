// Package health
package health

import (
	"net/http"
	"server/internal/db/pg"
	"server/internal/db/red"
	"server/internal/utils/resps"
)

func Health(w http.ResponseWriter, req *http.Request) {
	// Postgres health check
	if err := pg.DB.Ping(); err != nil {
		resps.RespJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "PostgreSql Database failed to init: " + err.Error(),
		})
		panic("Failed to init db")
	}
	// Redis health check
	if err := red.Client.Ping(red.Ctx).Err(); err != nil {
		resps.RespJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "Redis failed to init: " + err.Error(),
		})
		panic("Failed to init db")
	}
	// Success
	resps.RespJSON(w, http.StatusOK, map[string]string{
		"message": "Success",
	})
}

package reqs

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
	"net/http"
	"server/internal/consts"
	"server/internal/utils/logger"
)

const halfYear = 3600 * 24 * 180

func LogBody(c *gin.Context) {
	// 1. Read the body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Logger.Error("Failed to read request body", "error", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Request.Body = io.NopCloser(bytes.NewReader(body))

	slog.Info("Request body", "body", string(body))
}

func SimpleMessage(msg string) map[string]string {
	return map[string]string{"message": msg}
}

func SetServerCookie(c *gin.Context, name string, value string) {
	if consts.ENV == "dev" {
		c.SetCookie(name, value, halfYear, "/api", consts.SERVER_DOMAIN, false, true)
	} else {
		http.SetCookie(c.Writer, &http.Cookie{Name: name, Value: value, Domain: consts.SERVER_DOMAIN, Path: "/", MaxAge: halfYear, HttpOnly: true, Secure: true, SameSite: http.SameSiteDefaultMode})
	}
}

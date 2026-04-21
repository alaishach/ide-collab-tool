package reqs

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
	"net/http"
	"server/internal/utils/logger"
)

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

func SimpleResponseMessage(msg string) map[string]string {
	return map[string]string{"message": msg}
}

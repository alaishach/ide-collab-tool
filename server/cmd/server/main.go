// Package main
package main

import (
	"server/internal/api/gin"
	"server/internal/utils/logger"
)

func main() {
	logger.Logger.Debug("TEST Debug1")
	logger.Logger.Debug("TEST Debug2")
	gin.Run()
}

// Package auth
package auth

import (
	"server/internal/db/pg"

	"github.com/gin-gonic/gin"

	// "server/internal/db/red"
	// "server/internal/utils/funcs"
	"server/internal/err/pgerrors"
	"server/internal/utils/logger"
)

// GET login
// POST login
// POST Signup
// DELETE logout

type signupData struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Signup(c *gin.Context) {
	// TODO add field validation, create auth tokens, hash password
	var data signupData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, "")
	}
	logger.Logger.Debug("received signup request data: ", "username", data.Username, "email", data.Email, "password", data.Password)
	_, err := pg.CreateUser(data.Username, data.Email, data.Password)
	statusCode, marsh := pgerrors.GetDbErrorResp(err)
	c.JSON(statusCode, marsh)
}

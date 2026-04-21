// Package auth
package auth

import (
	"server/internal/db/pg"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	// "server/internal/db/red"
	// "server/internal/utils/funcs"
	"server/internal/err/panics"
	"server/internal/err/pgerrors"
	"server/internal/utils/logger"
	"server/internal/utils/reqs"
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
		c.JSON(400, reqs.SimpleResponseMessage("invalid request format"))
		return
	}
	logger.Logger.Debug("received signup request data: ", "username", data.Username, "email", data.Email, "password", data.Password)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	panics.PanicErr("Bcrypt failed to generated password", err)
	err = pg.CreateUser(data.Username, data.Email, passwordHash)
	if err != nil {
		statusCode, marsh := pgerrors.GetDbErrorResp(err)
		c.JSON(statusCode, marsh)
		return
	}
	c.JSON(201, reqs.SimpleResponseMessage("Account created"))
}

func PostLogin(c *gin.Context) {

}

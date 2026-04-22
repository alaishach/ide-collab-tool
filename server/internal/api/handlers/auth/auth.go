// Package auth
package auth

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"server/internal/db/pg"
	// "server/internal/db/red"
	// "server/internal/utils/funcs"
	"server/internal/err/errpg"
	errgl "server/internal/err/global"
	"server/internal/err/panics"
	"server/internal/utils/logger"
	"server/internal/utils/reqs"
)

// GET login check in redis and then in postgres and then add to redis if in postgres
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
		statusCode, marsh := errpg.GetDbErrorResp(err)
		c.JSON(statusCode, marsh)
		return
	}
	c.JSON(201, reqs.SimpleResponseMessage("Account created"))
}

type loginData struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func PostLogin(c *gin.Context) {
	println("!!!!! received post login")
	var data loginData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, reqs.SimpleResponseMessage("invalid request format"))
		return
	}
	// validation
	logger.Logger.Debug("received signup request data: ", "email", data.Email, "password", data.Password)
	userTable, error := pg.ValidCredentials(data.Email, data.Password)
	var errMessage *errgl.ErrMessage
	if errors.As(error, &errMessage) {
		c.JSON(500, errMessage.Message)
		return
	}

	sessionToken := uuid.New()
	if err := pg.CreateSession(userTable.UserID, sessionToken); err != nil {
		var pgErr *errpg.PgErr
		if errors.As(err, &pgErr) {
			c.JSON(500, "Failed to create session")
			return
		}
		panics.PanicDB("CreateSession", err)
		return
	}
	reqs.SetServerCookie(c, "sessionToken", sessionToken.String())
	c.JSON(201, reqs.SimpleResponseMessage("Session created"))
}

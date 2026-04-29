// Package auth
package auth

import (
	"errors"
	"net/http"
	"server/internal/api/schemas"
	"server/internal/db/pg"
	"server/internal/db/red"
	"server/internal/err/errpg"
	errgl "server/internal/err/global"
	"server/internal/err/panics"
	"server/internal/utils/logger"
	"server/internal/utils/reqs"
	"server/internal/utils/resps"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// GET login check in redis and then in postgres and then add to redis if in postgres
// POST login
// POST Signup
// DELETE logout

func Signup(w http.ResponseWriter, req *http.Request) {
	// TODO add field validation, create auth tokens, hash password
	var data schemas.SignupData
	if err := reqs.ParseBodyJSON(w, req, &data); err != nil {
		resps.RespMessage(w, 400, "invalid request format")
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(*data.Password), 14)
	panics.PanicErr("Bcrypt failed to generated password", err)
	err = pg.CreateUser(*data.Username, *data.Email, passwordHash)
	if err != nil {
		statusCode, marsh := errpg.GetDBErrorResp(err)
		resps.RespJSON(w, statusCode, marsh)
		return
	}
	resps.RespMessage(w, 201, "Account created")
}

func PostLogin(w http.ResponseWriter, req *http.Request) {
	var data schemas.LoginData
	// TODO do a function that json marshals and writes to http.ResponseWriter if there is an error
	// it should be able to check which params are missing from loginData
	if err := reqs.ParseBodyJSON(w, req, &data); err != nil {
		resps.RespMessage(w, 400, "invalid request format")
		return
	}

	// validation
	logger.Logger.Debug("received signup request data: ", "email", data.Email, "password", data.Password)
	userTable, error := pg.ValidCredentials(*data.Email, *data.Password)
	var errMessage *errgl.ErrMessage
	if errors.As(error, &errMessage) && errMessage != nil && errMessage.Type == "401" {
		resps.RespMessage(w, 401, (errMessage.Message))
		return
	} else if error != nil {
		resps.RespMessage(w, 500, errMessage.Message)
		return
	}

	// creating session
	sessionToken := uuid.New()
	if err := pg.CreateSession(userTable, sessionToken); err != nil {
		var pgErr *errpg.PgErr
		if errors.As(err, &pgErr) {
			resps.RespMessage(w, 500, "Failed to create session")
			return
		}
		panics.PanicDB("CreateSession", err)
		return
	}
	session := pg.SessionData{
		SessionToken: sessionToken,
		UserID:       userTable.UserID,
		Username:     userTable.Username,
	}
	red.AddSession(session)
	resps.SetServerCookie(w, "sessionToken", sessionToken.String())
	resps.RespMessage(w, 201, "Session created")
}

func GetLogin(w http.ResponseWriter, req *http.Request) {
	sessionToken, err := reqs.GetCookieValue(req, "sessionToken")
	if errors.Is(err, http.ErrNoCookie) {
		resps.RespMessage(w, 401, "Not authorized")
		return
	}
	session := red.GetSession(sessionToken)
	if session == nil {
		session = pg.GetSessionByToken(sessionToken)
		if session == nil || session.UserID == 0 {
			resps.RespMessage(w, 401, "Not authorized")
			return
		}
	}
	red.AddSession(*session)
	resps.RespMessage(w, 200, "success")
}

func Logout(w http.ResponseWriter, req *http.Request) {
	sessionToken, err := reqs.GetCookieValue(req, "sessionToken")
	if errors.Is(err, http.ErrNoCookie) {
		resps.RespMessage(w, 401, "Unauthorized")
	}
	red.DeleteSession(sessionToken)
	if err := pg.DeleteSession(sessionToken); err != nil {
		if errors.Is(err, errgl.ErrNotAuthorized) {
			resps.RespMessage(w, 401, "Unauthorized")
		} else {
			resps.RespMessage(w, 500, "Server side error")
		}
	}
}

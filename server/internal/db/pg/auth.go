// Package pg is dope
package pg

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"server/internal/err/errpg"
	errgl "server/internal/err/global"
	"server/internal/err/panics"
)

func CreateUser(username string, email string, password []byte) error {
	_, err := DB.Exec("insert into users (username, email, password) values ($1, $2, $3)", username, email, password)
	if err != nil {
		return errpg.NewPgError(err)
	}
	return nil
}

func CreateSession(user UserTable, sessionToken uuid.UUID) error {
	_, err := DB.Exec("insert into sessions (user_id, username, session_token) values ($1, $2, $3)", user.UserID, user.Username, sessionToken)
	if err != nil {
		return errpg.NewPgError(err)
	}
	return nil
}

// ValidCredentials POST login
func ValidCredentials(email string, password string) (UserTable, error) {
	var userTable UserTable
	err := DB.Get(&userTable, "select id, username, password from users where email=$1", email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userTable, errgl.NewErrMessage("No account tied to this email", "401")
		}
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			panics.PanicDB("ValidCredentials", err)
		}
	}
	if e := bcrypt.CompareHashAndPassword(userTable.PasswordHash, []byte(password)); e != nil {
		return userTable, errgl.NewErrMessage("Password is incorrect", "401")
	}
	return userTable, err
}

// GetSessionByToken GET login
func GetSessionByToken(sessionToken string) *SessionData {
	var session SessionData
	err := DB.Get(&session, "select user_id from sessions where session_token=$1", sessionToken)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil
	} else if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			errpg.NewPgError(pqErr)
		} else {
			panics.PanicDB("GetSessionByToken", err)
		}
	}
	return &session
}

func DeleteSession(sessionToken string) error {
	res, err := DB.Exec("delete from sessions where session_token=$1", sessionToken)
	if err != nil {
		return errpg.NewPgError(err)
	}
	if num, err := res.RowsAffected(); num == 0 || err != nil {
		if num == 0 {
			return errgl.ErrNotAuthorized
		} else if err != nil {
			return err
		}
	}
	return nil
}

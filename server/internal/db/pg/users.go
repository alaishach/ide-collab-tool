// Package pg is dope
package pg

import (
	"server/internal/err/errpg"
	errgl "server/internal/err/global"
	"server/internal/err/panics"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username string, email string, password []byte) error {
	_, err := DB.Exec("insert into users (username, email, password) values ($1, $2, $3)", username, email, password)
	if err != nil {
		return errpg.NewPgError(err)
	}
	return nil
}

func ReadUserCredentials(username string) (*PasswordSchema, error) {
	var passwordSchema *PasswordSchema
	err := DB.Get(passwordSchema, "select password, id from users where username=$1", username)
	if err != nil {
		return nil, errpg.NewPgError(err)
	}
	return passwordSchema, nil
}

func CreateSession(userID int, sessionToken uuid.UUID) error {
	_, err := DB.Exec("insert into sessions (user_id, session_token) values ($1, $2)", userID, sessionToken)
	if err != nil {
		return errpg.NewPgError(err)
	}
	return nil
}

func GetUsername(session Session) error {
	if session.UserID == 0 {
		panics.PanicMisuse("GetUsername", "UserId inside session object not set")
	}
	return DB.Get(&session, "select username from users where user_id=$1", session.UserID)
}

// POST login
func ValidCredentials(email string, password string) (UserTable, error) {
	var userTable UserTable
	err := DB.Get(&userTable, "select id, username, password from users where email=$1", email)
	if err != nil {
		panics.PanicDB("ValidCredentials", err)
	}
	if e := bcrypt.CompareHashAndPassword(userTable.PasswordHash, []byte(password)); e != nil {
		return userTable, errgl.NewErrMessage("Password is incorred")
	}
	return userTable, err
}

// GET login
func GetSessionByToken(sessionToken string) (*Session, error) {
	var session Session
	err := DB.Get(&session, "select user_id, from sessions where session_token=$1", sessionToken)
	if err != nil {
		return nil, errpg.NewPgError(err)
	}
	return &session, nil
}

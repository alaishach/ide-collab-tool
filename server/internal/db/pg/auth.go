// Package pg is dope
package pg

import (
	"database/sql"
	"errors"
	"server/internal/err/errpg"
	errgl "server/internal/err/global"
	"server/internal/err/panics"
	"strconv"

	"github.com/google/uuid"
	"github.com/lib/pq"
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

func CreateSession(user UserTable, sessionToken uuid.UUID) error {
	_, err := DB.Exec("insert into sessions (user_id, username, session_token) values ($1, $2, $3)", user.UserID, user.Username, sessionToken)
	if err != nil {
		return errpg.NewPgError(err)
	}
	return nil
}

func GetUsername(userID int) error {
	var username string
	if userID <= 0 {
		panics.PanicMisuse("GetUsername", "userId has invalid value: "+strconv.Itoa(userID))
	}
	return DB.Get(&username, "select username from users where id=$1", userID)
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
			println("IS PQ ERROR!!!!")
			println("Code: ", pqErr.Code)
			panics.PanicDB("ValidCredentials", err)
		}
	}
	if e := bcrypt.CompareHashAndPassword(userTable.PasswordHash, []byte(password)); e != nil {
		return userTable, errgl.NewErrMessage("Password is incorrect", "401")
	}
	return userTable, err
}

// GetSessionByToken GET login
func GetSessionByToken(sessionToken string) (*SessionData, error) {
	var session SessionData
	err := DB.Get(&session, "select user_id, from sessions where session_token=$1", sessionToken)
	if err != nil {
		return nil, errpg.NewPgError(err)
	}
	return &session, nil
}

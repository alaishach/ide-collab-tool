// Package pg is dope
package pg

import (
	"server/internal/err/pgerrors"
)

func CreateUser(username string, email string, password []byte) error {
	_, err := DB.Exec("insert into users (username, email, password) values ($1, $2, $3)", username, email, password)
	if err != nil {
		return pgerrors.NewPgError(err)
	}
	return nil
}

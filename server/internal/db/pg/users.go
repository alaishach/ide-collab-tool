// Package pg is dope
package pg

import (
	"server/internal/err/pgerrors"
)

func CreateUser(username string, email string, password string) (int64, error) {
	res, err := DB.Exec("insert into users (username, email, password) values ($1, $2, $3)", username, email, password)
	if err != nil {
		return 0, pgerrors.NewPgError(err)
	}
	return res.LastInsertId()
}

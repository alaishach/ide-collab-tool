// Package pg is dope
package pg

import "github.com/google/uuid"

type PasswordSchema struct {
	ID       int
	Password []byte
}

type SessionData struct {
	UserID       int       `db:"user_id" json:"user_id"`
	Username     string    `db:"username" json:"username"`
	SessionToken uuid.UUID `db:"session_token" json:"session_token"`
}

type UserTable struct {
	UserID       int    `db:"id"`
	Username     string `db:"username"`
	PasswordHash []byte `db:"password"`
}

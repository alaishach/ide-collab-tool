// Package pg is dope
package pg

type PasswordSchema struct {
	ID       int
	Password []byte
}

type Session struct {
	UserID   int    `db:"user_id"`
	Username string `db:"username"`
}

type UserTable struct {
	UserID       int    `db:"id"`
	Username     string `db:"username"`
	PasswordHash []byte `db:"password"`
}

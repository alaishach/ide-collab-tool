// Package errpg
package errpg

import (
	"errors"
	"reflect"
	"server/internal/utils/reqs"
	"strconv"
	"strings"

	"github.com/lib/pq"
)

// Postgres error code

type PgErr struct {
	Code    int    `binding:"required"`
	Message string `binding:"required"`
}

func (e *PgErr) Error() string {
	return e.Message
}

type PgErrDuplicate struct {
	Code    int    `binding:"required"`
	Message string `binding:"required"`
	Column  string `binding:"required"`
}

func (e *PgErrDuplicate) Error() string {
	return e.Message
}

type PgErrUnknown struct {
	Code    int    `binding:"required"`
	Message string `binding:"required"`
}

func (e *PgErrUnknown) Error() string {
	return e.Message
}

func NewPgErrUnknown(code int, message string) *PgErrUnknown {
	return &PgErrUnknown{
		Code:    code,
		Message: message,
	}
}

func handleDuplicate(code int, pgError *pq.Error) error {
	detail := pgError.Detail
	col := "some column"
	if detail[0:4] == "Key " {
		i := strings.Index(detail, "(")
		j := strings.Index(detail, ")")
		col = detail[i+1 : j]
	}
	return &PgErrDuplicate{
		Code:    code,
		Message: pgError.Message,
		Column:  col,
	}
}

func NewPgError(err error) error {
	var pgError *pq.Error
	if !errors.As(err, &pgError) {
		panic("Wrong error type passed to NewPgError" + reflect.TypeOf(err).Name())
	}
	code, _ := strconv.Atoi(string(pgError.Code))
	switch code {
	case 23505:
		return handleDuplicate(code, pgError)
	}
	return NewPgErrUnknown(code, pgError.Message)
}

func GetDbErrorResp(err error) (int, map[string]string) {
	var pgErrDuplicate *PgErrDuplicate
	if errors.As(err, &pgErrDuplicate) {
		return 409, reqs.SimpleResponseMessage(pgErrDuplicate.Column + " is already taken")
	}
	return 200, reqs.SimpleResponseMessage("")
}

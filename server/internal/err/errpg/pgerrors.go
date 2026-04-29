// Package errpg
package errpg

import (
	"errors"
	"github.com/lib/pq"
	"reflect"
	"server/internal/err/panics"
	"server/internal/utils/logger"
	"strings"
)

// Postgres error code

type PgErr struct {
	Code    string `binding:"required"`
	Message string `binding:"required"`
}

func (e *PgErr) Error() string {
	return e.Message
}

type PgErrDuplicate struct {
	Code    string `binding:"required"`
	Message string `binding:"required"`
	Column  string `binding:"required"`
}

func (e *PgErrDuplicate) Error() string {
	return e.Message
}

type PgErrUnknown struct {
	Code    string `binding:"required"`
	Message string `binding:"required"`
}

func (e *PgErrUnknown) Error() string {
	return e.Message
}

func NewPgErrUnknown(code string, message string) *PgErrUnknown {
	return &PgErrUnknown{
		Code:    code,
		Message: message,
	}
}

func handleDuplicate(code string, pgError *pq.Error) error {
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

type PgErrInvalidInput struct {
	Code    string `binding:"required"`
	Message string `binding:"required"`
}

func (e *PgErrInvalidInput) Error() string {
	return e.Message
}

// TODO make sure to have proper error message with column not just table name
func handleInvalidInput(code string, pgError *pq.Error) error {
	detail := pgError.Detail
	println("!!!!!!!!!!!: ", detail)
	return &PgErrInvalidInput{
		Code:    code,
		Message: "Invalid input" + pgError.Table,
	}

}

func NewPgError(err error) error {
	var pgError *pq.Error
	if !errors.As(err, &pgError) {
		panics.PanicMisuse("NewPgError", "Wrong error type passed to NewPgError"+reflect.TypeOf(err).Name())
	}
	code := string(pgError.Code)
	switch code {
	case "23505":
		return handleDuplicate(code, pgError)
	case "22P02":
		return handleInvalidInput(code, pgError)
	}
	return NewPgErrUnknown(code, pgError.Message)
}

func GetDBErrorResp(err error) (int, map[string]string) {
	var pgErrDuplicate *PgErrDuplicate
	if errors.As(err, &pgErrDuplicate) {
		return 409, map[string]string{"message": pgErrDuplicate.Column + " is already taken"}
	}
	logger.Logger.Error("Unhandled error: " + err.Error())
	return 500, map[string]string{"message": pgErrDuplicate.Column + "Unhandled Error"}
}

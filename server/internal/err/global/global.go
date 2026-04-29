// Package errgl
package errgl

import "errors"

type ErrMessage struct {
	Message string `binding:"required"`
	Type    string `binding:"required"`
}

func (e *ErrMessage) Error() string {
	return e.Message
}

func NewErrMessage(msg string, t string) *ErrMessage {
	return &ErrMessage{
		Message: msg,
		Type:    t,
	}
}

var ErrNotAuthorized = errors.New("not authorized")

type HTTPError struct {
	Code    int    `binding:"required"`
	Message string `binding:"required"`
}

func (e *HTTPError) Error() string {
	return e.Message
}

func NewHTTPError(code int, message string) *HTTPError {
	return &HTTPError{
		Code:    code,
		Message: message,
	}
}

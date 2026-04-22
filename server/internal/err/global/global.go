// Package errgl
package errgl

type ErrMessage struct {
	Message string `binding:"required"`
}

func (e *ErrMessage) Error() string {
	return e.Message
}

func NewErrMessage(msg string) *ErrMessage {
	return &ErrMessage{
		Message: msg,
	}
}

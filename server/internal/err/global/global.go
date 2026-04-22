// Package errgl
package errgl

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

type NotAuthorized struct {
	Message string `binding:"required"`
}

func (e *NotAuthorized) Error() string {
	return e.Message
}

func NewNotAuthorized(msg string) *ErrMessage {
	return &ErrMessage{
		Message: msg,
	}
}

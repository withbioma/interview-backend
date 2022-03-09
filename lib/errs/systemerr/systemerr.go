package systemerr

import "fmt"

// SystemError canonical struct for displaying system-specific errors.
//
// To be differentiated vs. ClientError. SystemError is only logged, and sent to Sentry but will
// not be shown to users (i.e. sent back as a HTTP response).
type SystemError interface {
	Error() string
}

type DefaultSystemError struct {
	Message string `json:"message"`
}

func (e DefaultSystemError) Error() string {
	return e.Message
}

// New creates new error object.
func New(message string) DefaultSystemError {
	return DefaultSystemError{Message: message}
}

// Newf creates new error object with formatted string.
func Newf(format string, args ...interface{}) DefaultSystemError {
	return DefaultSystemError{Message: fmt.Sprintf(format, args...)}
}

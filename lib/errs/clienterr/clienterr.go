package clienterr

import (
	"github.com/withbioma/interview-backend/lib/errs/systemerr"
)

// ClientError user-friendly error interface to be shown to users.
//
// There is a difference between ClientError and regular Error (found in errs/system.go). Whereas
// Error is for system-specific errors (i.e. things that will result in 5xx http error codes),
// ClientError is for 4xx http error codes.
type ClientError interface {
	// Error - to conform with the standard Go error interface. This method returns the error
	// message/payload.
	Error() string
	// Code - non-changing or seldom-changing string to be treated as keys. The purpose of having
	// this code is so that the frontend does not have to match server-sent errors based on the
	// message, which is often changing.
	ErrorCode() string
	// OriginalError - the original error that should not be returned back to the client. This error
	// should only be logged or sent to Sentry.
	OriginalError() error
	Resolve() APIError
}

// APIError the resolved error to be sent back to clients.
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// DefaultAPIError struct for standard API error.
//
// You should use this struct if you are not dealing with implementation-specific errors. For things
// like coordinators that require errors to be recognized properly, consider creating another
// specific error struct for it. Otherwise, please use this struct.
type DefaultAPIError struct {
	Code   Code          `json:"code"`
	Values []interface{} `json:"-"`

	SystemError systemerr.SystemError `json:"-"`
}

func (e DefaultAPIError) Error() string {
	return string(GenerateMessage(e.Code, e.Values...))
}

func (e DefaultAPIError) ErrorCode() string {
	return string(e.Code)
}

func (e DefaultAPIError) OriginalError() error {
	return e.SystemError
}

func (e DefaultAPIError) Resolve() APIError {
	return APIError{
		Code:    e.ErrorCode(),
		Message: GenerateMessage(e.Code, e.Values...),
	}
}

// New creates new api error object.
func New(code Code) DefaultAPIError {
	return DefaultAPIError{Code: code}
}

// NewWithError creates new error object with original system error.
func NewWithError(code Code, err systemerr.SystemError) DefaultAPIError {
	return DefaultAPIError{Code: code, SystemError: err}
}

// Newf creates new api error object with arguments.
func Newf(code Code, args ...interface{}) DefaultAPIError {
	return DefaultAPIError{Code: code, Values: args}
}

// NewfWithError creates new api error object with arguments with original system error.
func NewfWithError(
	code Code,
	err systemerr.SystemError,
	args ...interface{},
) DefaultAPIError {
	return DefaultAPIError{Code: code, Values: args, SystemError: err}
}

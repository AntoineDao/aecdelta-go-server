package store

import "errors"

type ConnectionError interface {
	error
	StatusCode() int
}

// errorString is a trivial implementation of error.
type ErrorHTTP struct {
	Code int `json:"status_code"`
	Err  error
}

func (err *ErrorHTTP) StatusCode() int {
	return err.Code
}

func (e *ErrorHTTP) Error() string {
	return e.Err.Error()
}

func HTTPError(statusCode int, message string) error {
	return &ErrorHTTP{Code: statusCode, Err: errors.New(message)}
}

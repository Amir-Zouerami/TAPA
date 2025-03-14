package errors

import (
	"fmt"
)

type ErrorCode int

type AppError struct {
	Code    ErrorCode
	Message string
	cause   error
}

func (e *AppError) Error() string {
	if e.cause == nil {
		return fmt.Sprintf("[%d] %s", e.Code, e.Message)
	}

	return fmt.Sprintf("[%d] %s: %v", e.Code, e.Message, e.cause)
}

func (e *AppError) Unwrap() error {
	return e.cause
}

var (
	ErrAppIconRead         = &AppError{Code: 1000, Message: "App icon read error"}
	ErrAppConfigGeneration = &AppError{Code: 1001, Message: "App config generation error"}
	ErrEmbeddedFileRead    = &AppError{Code: 1002, Message: "Embedded file read error"}
)

func Wrap(err *AppError, cause error) error {
	return &AppError{
		Code:    err.Code,
		Message: err.Message,
		cause:   cause,
	}
}

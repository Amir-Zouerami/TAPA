package errors

import (
	"fmt"
)

type ErrorCode int

type TapaError struct {
	Code    ErrorCode
	Message string
	cause   error
}

func (e *TapaError) Error() string {
	if e.cause == nil {
		return fmt.Sprintf("[%d] %s", e.Code, e.Message)
	}

	return fmt.Sprintf("[%d] %s: %v", e.Code, e.Message, e.cause)
}

func (e *TapaError) Unwrap() error {
	return e.cause
}

func Wrap(err *TapaError, cause error) error {
	return &TapaError{
		Code:    err.Code,
		Message: err.Message,
		cause:   cause,
	}
}

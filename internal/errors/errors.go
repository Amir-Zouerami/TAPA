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

func Wrap(err *AppError, cause error) error {
	return &AppError{
		Code:    err.Code,
		Message: err.Message,
		cause:   cause,
	}
}

// ------------- GENERAL APP ERRORS (1000)
var (
	ErrAppIconRead            = &AppError{Code: 1000, Message: "App icon read error \n"}
	ErrAppConfigGeneration    = &AppError{Code: 1001, Message: "App config generation error \n"}
	ErrEmbeddedFileRead       = &AppError{Code: 1002, Message: "Embedded file read error \n"}
	ErrGetUserConfigDirectory = &AppError{Code: 1003, Message: "Getting user config directory failed \n"}
	ErrCreateAppDirectory     = &AppError{Code: 1004, Message: "Failed creating app directory \n"}
)

// ------------- DATABASE ERRORS (2000)
var (
	ErrSchemaRead          = &AppError{Code: 2001, Message: "Database schema read error \n"}
	ErrSchemaCreation      = &AppError{Code: 2002, Message: "Database schema creation error \n"}
	ErrOpeningDatabaseFile = &AppError{Code: 2003, Message: "Opening database file failed \n"}
	ErrConnectingDatabase  = &AppError{Code: 2004, Message: "Connecting to database failed \n"}
)

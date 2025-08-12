package errors

import (
	"fmt"
)

// ErrorCode represents a specific error type
type ErrorCode string

const (
	ErrCodeInvalidInput   ErrorCode = "INVALID_INPUT"
	ErrCodeModelLoad      ErrorCode = "MODEL_LOAD"
	ErrCodeProcessing     ErrorCode = "PROCESSING"
	ErrCodeIO             ErrorCode = "IO"
	ErrCodeConfig         ErrorCode = "CONFIG"
	ErrCodeMemory         ErrorCode = "MEMORY"
	ErrCodeGPU            ErrorCode = "GPU"
	ErrCodeUnsupported    ErrorCode = "UNSUPPORTED"
)

// AppError represents an application error with context
type AppError struct {
	Code    ErrorCode
	Message string
	Cause   error
	Context map[string]interface{}
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// Unwrap returns the underlying error
func (e *AppError) Unwrap() error {
	return e.Cause
}

// New creates a new application error
func New(code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Context: make(map[string]interface{}),
	}
}

// Wrap wraps an existing error with application context
func Wrap(err error, code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Cause:   err,
		Context: make(map[string]interface{}),
	}
}

// WithContext adds context to the error
func (e *AppError) WithContext(key string, value interface{}) *AppError {
	e.Context[key] = value
	return e
}

// Is checks if the error matches a specific code
func Is(err error, code ErrorCode) bool {
	appErr, ok := err.(*AppError)
	if !ok {
		return false
	}
	return appErr.Code == code
}
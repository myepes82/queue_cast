package errors

import "fmt"

type ApplicationError struct {
	Message string
}

func (e *ApplicationError) Error() string {
	return fmt.Sprintf("Application error: %s", e.Message)
}

const (
	ErrInvalidLogLevelMessage = "Invalid log level"
)

var (
	ErrInvalidLogLevel = &ApplicationError{Message: ErrInvalidLogLevelMessage}
)

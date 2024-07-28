package errors

import "fmt"

type ConfigurationError struct {
	Message string
}

func (e *ConfigurationError) Error() string {
	return fmt.Sprintf("Configuration error: %s", e.Message)
}

const (
	errConfigFileNotFoundMessage = "Config file not found"
	errConfigMissingValueMessage = "Missing value in config"
)

var (
	ErrConfigFileNotFound = &ConfigurationError{Message: errConfigFileNotFoundMessage}
	ErrConfigMissingValue = &ConfigurationError{Message: errConfigMissingValueMessage}
)

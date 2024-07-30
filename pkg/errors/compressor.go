package errors

import "fmt"

type CompressorError struct {
	Message string
}

func (e *CompressorError) Error() string {
	return fmt.Sprintf("Compressor error: %s", e.Message)
}

const (
	errCompressionFailedMessage   = "compression failed"
	errDecompressionFailedMessage = "decompression failed"
)

var (
	ErrCompressionFailed   = &CompressorError{Message: errCompressionFailedMessage}
	ErrDecompressionFailed = &CompressorError{Message: errDecompressionFailedMessage}
)

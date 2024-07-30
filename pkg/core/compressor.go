package core

import (
	"bytes"
	"compress/gzip"
	"io"
)

type MessageCompressor struct {
	logger *Logger
}

func NewMessageCompressor(logger *Logger) *MessageCompressor {
	logger.Info("Creating new message compressor")
	defer logger.Info("Message compressor created")
	return &MessageCompressor{
		logger: logger,
	}
}

func (c *MessageCompressor) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	if _, err := writer.Write(data); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *MessageCompressor) Decompress(data []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return io.ReadAll(reader)
}

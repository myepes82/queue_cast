package service

import (
	"bytes"
	"compress/gzip"
	"io"

	"go.uber.org/zap"
)

type MessageCompressor struct {
	logger *zap.Logger
}

func NewMessageCompressor(logger *zap.Logger) *MessageCompressor {
	logger.Info("Creating new message compressor")
	defer logger.Info("Message compressor created")
	return &MessageCompressor{
		logger: logger,
	}
}

func (c *MessageCompressor) Compress(data string) ([]byte, error) {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)

	_, err := writer.Write([]byte(data))
	if err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		c.logger.Error("Error closing gzip writer", zap.Error(err))
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *MessageCompressor) Decompress(data []byte) (string, error) {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	defer func(rd *gzip.Reader) {
		if err := rd.Close(); err != nil {
			c.logger.Error("Error closing reader", zap.Error(err))
		}
	}(reader)

	decompressed, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(decompressed), nil
}

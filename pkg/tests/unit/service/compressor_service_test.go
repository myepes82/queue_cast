package service_test

import (
	"queuecast/pkg/tests"
	test_data "queuecast/pkg/tests/data"
	"testing"

	"github.com/stretchr/testify/require"
	"queuecast/pkg/core/service"
)

func TestCompressorMessageService(t *testing.T) {
	logger := tests.GetTestLogger()

	messageCompressor := service.NewMessageCompressor(logger)
	require.NotNil(t, messageCompressor)

	t.Run("given a simple test message, compress and decompress it", func(t *testing.T) {
		// Arrange
		originalData := "hello world"

		// Act - Compress
		compressedData, err := messageCompressor.Compress(originalData)
		require.NoError(t, err)
		require.NotEmpty(t, compressedData)

		// Act - Decompress
		decompressedData, err := messageCompressor.Decompress(compressedData)
		require.NoError(t, err)

		// Assert
		require.NotEmpty(t, decompressedData)
		require.Equal(t, originalData, decompressedData)
	})

	t.Run("compress and decompress an empty message", func(t *testing.T) {
		// Arrange
		emptyData := ""

		// Act - Compress
		compressedData, err := messageCompressor.Compress(emptyData)
		require.NoError(t, err)
		require.NotEmpty(t, compressedData)

		// Act - Decompress
		decompressedData, err := messageCompressor.Decompress(compressedData)
		require.NoError(t, err)

		// Assert
		require.Empty(t, decompressedData)
		require.Equal(t, emptyData, decompressedData)
	})

	t.Run("attempt to decompress invalid data", func(t *testing.T) {
		// Arrange
		invalidData := []byte("This is not compressed data")

		// Act
		_, err := messageCompressor.Decompress(invalidData)

		// Assert
		require.Error(t, err)
	})

	t.Run("check if compressed size is smaller than original", func(t *testing.T) {
		// Arrange
		originalData := test_data.ComplexJsonObject

		// Act - Compress
		compressedData, err := messageCompressor.Compress(originalData)
		require.NoError(t, err)

		// Assert
		t.Logf("Original size: %d bytes", len(originalData))
		t.Logf("Compressed size: %d bytes", len(compressedData))
		require.Less(t, len(compressedData), len(originalData), "Compressed data should be smaller than original data")

		// Decompress to verify data integrity
		decompressedData, err := messageCompressor.Decompress(compressedData)
		require.NoError(t, err)
		require.Equal(t, originalData, decompressedData, "Decompressed data should match original data")
	})

}

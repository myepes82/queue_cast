package definitions

type Compressor interface {
	Compress(data string) ([]byte, error)
	Decompress(data []byte) (string, error)
}

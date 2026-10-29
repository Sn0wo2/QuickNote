package compress

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io"
)

const maxDecompressedSize = 64 << 20

func FlateCompress(data []byte) ([]byte, error) {
	var b bytes.Buffer
	// level: flate.DefaultCompression, flate.BestSpeed, etc.
	writer, err := flate.NewWriter(&b, flate.BestSpeed)
	if err != nil {
		return nil, err
	}

	_, err = writer.Write(data)
	if err != nil {
		return nil, err
	}

	_ = writer.Close()

	return b.Bytes(), nil
}

func FlateDecompress(compressed []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(compressed))
	defer func() {
		_ = reader.Close()
	}()

	var out bytes.Buffer

	_, err := io.Copy(&out, io.LimitReader(reader, maxDecompressedSize+1))
	if err != nil {
		return nil, err
	}

	if out.Len() > maxDecompressedSize {
		return nil, fmt.Errorf("decompressed data exceeds %d bytes", maxDecompressedSize)
	}

	return out.Bytes(), nil
}

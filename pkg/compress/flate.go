package compress

import (
	"bytes"
	"compress/flate"
	"io"
)

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

	_, err := io.Copy(&out, reader)
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

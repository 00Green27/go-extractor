package extractor

import (
	"bytes"
	"net/http"
	"os"
)

func mimeType(src string) (string, error) {
	f, err := os.Open(src)
	if err != nil {
		return "", err
	}

	defer f.Close()
	data := make([]byte, 512)

	_, err = f.Read(data)
	if err != nil {
		return "", err
	}

	// Check for errors in http.DetectContentType
	if m := match(data); m != "" {
		return m, nil
	}

	return http.DetectContentType(data), nil
}

func match(data []byte) string {
	if bytes.HasPrefix(data, []byte("\x52\x61\x72\x21\x1A\x07\x00")) {
		return "application/x-rar-compressed"
	}
	return ""
}

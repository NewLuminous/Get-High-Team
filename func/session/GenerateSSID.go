package session

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

func GenerateSSID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

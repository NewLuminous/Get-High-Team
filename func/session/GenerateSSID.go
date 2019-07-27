package session

import (
    "io"
    "encoding/base64"
    "crypto/rand"
)

func GenerateSSID() string {
        b := make([]byte, 32)
        if _, err := io.ReadFull(rand.Reader, b); err != nil {
            return ""
        }
        return base64.URLEncoding.EncodeToString(b)
 }

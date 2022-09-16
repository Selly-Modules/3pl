package onpoint

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func hashSHA256(data, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func hashSHA256AndUppercase(data, key string) string {
	return strings.ToUpper(hashSHA256(data, key))
}

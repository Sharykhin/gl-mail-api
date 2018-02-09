package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func Hmac256(src string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(src))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func CompareHmac(message string, messageHmac string, secret string) bool {
	return messageHmac == Hmac256(message, secret)
}

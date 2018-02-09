package jwt

import (
	"encoding/base64"
	"fmt"
)

func Base64Encode(src string) string {
	data := []byte(src)
	str := base64.StdEncoding.EncodeToString(data)
	return str
}

func Base64Decode(src string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		errMsg := fmt.Errorf("Decoding Error %s", err)
		return "", errMsg
	}
	return string(decoded), nil
}

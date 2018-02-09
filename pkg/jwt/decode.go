package jwt

import (
	"strings"
	"fmt"
	"encoding/json"
	"time"
	"errors"
)

func Decode(jwt string, secret string) (interface{}, error) {
	token := strings.Split(jwt, ".")

	// check if the jwt token contains
	// header, payload and token
	if len(token) != 3 {
		splitErr := errors.New("Invalid token: token should contain header, payload and secret")
		return nil, splitErr
	}
	// decode payload
	decodedPayload, PayloadErr := Base64Decode(token[1])
	if PayloadErr != nil {
		return nil, fmt.Errorf("Invalid payload: %s", PayloadErr.Error())
	}
	payload := Payload{}

	// parses payload from string to a struct
	ParseErr := json.Unmarshal([]byte(decodedPayload), &payload)
	if ParseErr != nil {
		return nil, fmt.Errorf("Invalid payload: %s", ParseErr.Error())
	}

	if payload.Exp != 0 && time.Now().Unix() > payload.Exp {
		return nil, errors.New("Expired token: token has expired")
	}

	signatureValue := token[0] + "." + token[1]

	// verifies if the header and signature is exactly whats in
	// the signature
	if CompareHmac(signatureValue, token[2], secret) == false {
		return nil, errors.New("Invalid token")
	}

	return payload, nil
}

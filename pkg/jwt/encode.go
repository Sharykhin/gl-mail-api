package jwt

import "encoding/json"

func Encode(payload Payload, secret string) string {

	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	str, _ := json.Marshal(header)
	headerString := Base64Encode(string(str))
	encodedPayload, _ := json.Marshal(payload)
	signatureValue := headerString + "." +
		Base64Encode(string(encodedPayload))
	return signatureValue + "." + Hmac256(signatureValue, secret)
}

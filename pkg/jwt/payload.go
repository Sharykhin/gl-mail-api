package jwt

import "time"

type Payload struct {
	Sub    string      `json:"sub,omitempty"`
	Exp    int64       `json:"exp,omitempty"`
	Iss    string      `json:"iss,omitempty"`
	Aud    string      `json:"aud,omitempty"`
	Public interface{} `json:"public,omitempty"`
}

func NewPayload(c map[string]int) Payload {
	var payload = Payload{
		"api",
		time.Now().Add(time.Minute * MINUTES_EXP).Unix(),
		"localhost:8004",
		"",
		c,
	}
	return payload
}

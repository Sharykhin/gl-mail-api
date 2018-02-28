package entity

import (
	"encoding/json"
	"time"
)

// FailMail is a basis entity
type FailMail struct {
	ID        int64     `json:"id"`
	Action    string    `json:"action"`
	Payload   Payload   `json:"payload"`
	Reason    string    `json:"reason"`
	CreatedAt JSONTime  `json:"created_at"`
	DeletedAt *JSONTime `json:"deleted_at"`
}

// JSONTime represents time format that should be returned to a client
type JSONTime time.Time

// MarshalJSON implements common interface for changing marshaling
func (t JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format(time.RFC1123) + `"`), nil
}

// Payload is a specific time for json struct of a payload
type Payload json.RawMessage

// MarshalJSON implements general interface to change representation of payload field
func (p Payload) MarshalJSON() ([]byte, error) {
	return []byte(string(p)), nil
}

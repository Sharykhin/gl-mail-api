package entity

import (
	"encoding/json"
	"fmt"
	"strings"
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

// FailMailRequest represents income request body
type FailMailRequest struct {
	Action  string  `json:"action"`
	Payload Payload `json:"payload"`
	Reason  string  `json:"reason"`
}

// Validate - implementation of the InputValidation interface
func (fmr FailMailRequest) Validate() error {
	if strings.Trim(fmr.Action, " ") == "" {
		return fmt.Errorf("action is required")
	}

	if fmr.Payload == nil {
		return fmt.Errorf("payload is required")
	}

	if strings.Trim(fmr.Reason, " ") == "" {
		return fmt.Errorf("reason is required")
	}

	return nil
}

// JSONTime represents time format that should be returned to a client
type JSONTime time.Time

// MarshalJSON implements common interface for changing marshaling
func (t JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format(time.RFC1123) + `"`), nil
}

// Payload is a specific time for json struct of a payload
type Payload map[string]json.RawMessage

// Scan implements a common interface for scanning values from database source to a specific struct
func (p *Payload) Scan(v interface{}) error {
	err := json.Unmarshal(v.([]byte), &p)
	return err
}

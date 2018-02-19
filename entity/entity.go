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
	Action  string          `json:"action"`
	Payload json.RawMessage `json:"payload"`
	Reason  string          `json:"reason"`
}

// Validate - implementation of the InputValidation interface
func (fmr FailMailRequest) Validate() error {
	if strings.Trim(fmr.Action, " ") == "" {
		return fmt.Errorf("action is required")
	}

	var stuff struct{}
	m, err := fmr.Payload.MarshalJSON()
	if err != nil {
		return fmt.Errorf("payload must be a valid json")
	}
	err = json.Unmarshal(m, &stuff)
	if err != nil {
		return fmt.Errorf("payload must be a valid json")
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
type Payload json.RawMessage

//@QUESTION: is it a good way to convert data from storage to a specific struct value

// Scan implements a common interface for scanning values from database source to a specific struct
//func (p Payload) Scan(v interface{}) error {
//	fmt.Println("a", p)
//	err := json.Unmarshal(v.([]byte), &p)
//	return err
//}

func (p Payload) MarshalJSON() ([]byte, error) {
	return []byte(string(p)), nil
}

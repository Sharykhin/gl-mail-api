package entity

import (
	"fmt"
	"strings"
	"time"
)

// InputValidation - an interface for all request structs
type InputValidation interface {
	Validate() error
}

// FailMailRequest represents income request body
type FailMailRequest struct {
	Action  string                 `json:"action"`
	Payload map[string]interface{} `json:"payload"`
	Reason  string                 `json:"reason"`
}

// Validate - implementation of the InputValidation interface
func (mr FailMailRequest) Validate() error {
	if strings.Trim(mr.Action, " ") == "" {
		return fmt.Errorf("action is required")
	}

	if mr.Payload == nil {
		return fmt.Errorf("payload is required")
	}

	if strings.Trim(mr.Reason, " ") == "" {
		return fmt.Errorf("reason is required")
	}

	return nil
}

// Message is a basis entity
type Message struct {
	ID        int64                  `json:"id"`
	Action    string                 `json:"action"`
	Payload   map[string]interface{} `json:"payload"`
	Reason    string                 `json:"reason"`
	CreatedAt time.Time              `json:"created_at"`
	DeletedAt time.Time              `json:"deleted_at"`
}

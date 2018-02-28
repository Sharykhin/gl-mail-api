package entity

import (
	"encoding/json"
)

// FailMail is a basis entity
type FailMail struct {
	ID        int64           `json:"id"`
	Action    string          `json:"action"`
	Payload   json.RawMessage `json:"payload"`
	Reason    string          `json:"reason"`
	CreatedAt string          `json:"created_at"`
	DeletedAt *string         `json:"deleted_at"`
}

package entity

// MessageRequest represent income request body
type MessageRequest struct {
	Action  string                 `json:"action"`
	Payload map[string]interface{} `json:"payload"`
}

// Message is a basis entity
type Message struct {
	ID      int64                  `json:"id"`
	Action  string                 `json:"action"`
	Payload map[string]interface{} `json:"payload"`
}

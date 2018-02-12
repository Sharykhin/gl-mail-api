package entity

type MessageRequest struct {
	Action  string                 `json:action`
	Payload map[string]interface{} `json:payload`
}

type Message struct {
	ID      int64                  `json:"id"`
	Action  string                 `json:action`
	Payload map[string]interface{} `json:payload`
}

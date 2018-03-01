package util

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response struct represents base response format
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   error       `json:"error"`
	Meta    interface{} `json:"meta"`
}

// SendResponse send json response to a client
func SendResponse(r Response, w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		log.Fatalf("could not return response, struct: %v, error: %v", r, err)
	}
}

// TODO: just because we want to get error as type we need to duplicate quite a big part of code. Is it ok?

// MarshalJSON implement MarshalJSON interface method to convert error type to string
func (res Response) MarshalJSON() ([]byte, error) {
	var err interface{}
	if res.Error != nil {
		err = res.Error.Error()
	}
	return json.Marshal(&struct {
		Success bool        `json:"success"`
		Data    interface{} `json:"data"`
		Error   interface{} `json:"error"`
		Meta    interface{} `json:"meta"`
	}{
		Success: res.Success,
		Data:    res.Data,
		Error:   err,
		Meta:    res.Meta,
	})
}

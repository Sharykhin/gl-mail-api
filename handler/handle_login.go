package handler

import (
	"encoding/json"
	"log"
	"github.com/Sharykhin/gl-mail-api/pkg/api"
	"github.com/Sharykhin/gl-mail-api/pkg/jwt"
	"net/http"
	"time"

	"github.com/Sharykhin/gl-mail-api/service"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	//TODO: this check should be moved to a middleware or somewhere else
	if r.Method != "POST" {
		var response = service.Response{
			Success: false,
			Data:    nil,
			Error:   http.StatusText(http.StatusMethodNotAllowed),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	//TODO: this dependency should be passed as an income parameter
	var apiService *api.ApiService
	apiService = api.NewApiService()

	decode := json.NewDecoder(r.Body)
	var cb api.CredentialBody
	err := decode.Decode(&cb)
	if err != nil {
		log.Fatalf("Failed to parse json body. Error: %s", err)
	}
	defer r.Body.Close()

	var cc = new(api.CredentialModel)
	cc, err = apiService.Credential.GetCredentials(cb.ApiKey, cb.SecretKey)
	if err != nil {
		var response = service.Response{
			Success: false,
			Data:   nil,
			Error:  "Credentials are incorrect",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	payload := jwt.NewPayload(map[string]int{
		"id": cc.Id,
	})
	token := jwt.Encode(payload, jwt.SECRET)
	var response = service.Response{}
	response.Data = map[string]interface{}{
		"access_token": token,
		"expires_at":   time.Now().Add(time.Minute * 5).Unix(),
		"type":         "Mail",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
}

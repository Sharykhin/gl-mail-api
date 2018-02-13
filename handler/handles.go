package handler

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/Sharykhin/gl-mail-api/controller"
	"github.com/Sharykhin/gl-mail-api/entity"
	"github.com/Sharykhin/gl-mail-api/util"
)

func pong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("pong"))
	if err != nil {
		log.Fatalf("something really akward went wrong: %v", err)
	}
}

func getFailedMailsList(w http.ResponseWriter, r *http.Request) {
	util.SendResponse(util.Response{Success: true, Data: nil, Error: nil}, w, http.StatusOK)
}

func createFailedMail(w http.ResponseWriter, r *http.Request) {
	var mr entity.MessageRequest

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close() // nolint: errcheck

	err := decoder.Decode(&mr)
	if err != nil {
		log.Printf("could not decode income request to struct: %s, error: %v", mr, err)
		util.SendResponse(util.Response{Success: false, Data: nil, Error: err}, w, http.StatusBadRequest)
		return
	}
	// TODO: this one should be mocked
	err = validate(mr)
	if err != nil {
		util.SendResponse(util.Response{Success: false, Data: nil, Error: err}, w, http.StatusBadRequest)
		return
	}
	// TODO: this one should be mocked
	m, err := controller.Create(r.Context(), mr)
	if err != nil {
		util.SendResponse(util.Response{Success: false, Data: nil, Error: err}, w, http.StatusInternalServerError)
		return
	}

	util.SendResponse(util.Response{Success: true, Data: m, Error: nil}, w, http.StatusCreated)
}

func validate(v entity.InputValidation) error {
	return v.Validate()
}

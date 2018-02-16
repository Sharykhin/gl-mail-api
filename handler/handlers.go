package handler

import (
	"log"
	"net/http"

	"encoding/json"

	"strconv"

	"github.com/Sharykhin/gl-mail-api/controller"
	db "github.com/Sharykhin/gl-mail-api/database"
	"github.com/Sharykhin/gl-mail-api/entity"
	"github.com/Sharykhin/gl-mail-api/util"
)

func pong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte("pong"))
	if err != nil {
		log.Fatalf("something really akward went wrong: %v", err)
	}
}

func getFailedMailsList(w http.ResponseWriter, r *http.Request) {
	limit, err := queryParamInt(r, "limit", 10)
	if err != nil {
		util.SendResponse(util.Response{Success: false, Data: nil, Error: err}, w, http.StatusBadRequest)
		return
	}

	offset, err := queryParamInt(r, "offset", 0)
	if err != nil {
		util.SendResponse(util.Response{Success: false, Data: nil, Error: err}, w, http.StatusBadRequest)
		return
	}

	m, c, err := controller.GetList(r.Context(), db.Storage, limit, offset)

	if err != nil {
		util.SendResponse(util.Response{Success: false, Data: nil, Error: err}, w, http.StatusInternalServerError)
		return
	}

	util.SendResponse(util.Response{Success: true, Data: map[string]interface{}{
		"mails": m,
		"count": c,
	}, Error: nil}, w, http.StatusOK)
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
	// TODO: this one should be mocked in case of unit tests
	err = validate(mr)
	if err != nil {
		util.SendResponse(util.Response{Success: false, Data: nil, Error: err}, w, http.StatusBadRequest)
		return
	}
	// TODO: this one should be mocked in case of unit tests
	m, err := controller.Create(r.Context(), mr, db.Storage)
	if err != nil {
		util.SendResponse(util.Response{Success: false, Data: nil, Error: err}, w, http.StatusInternalServerError)
		return
	}

	util.SendResponse(util.Response{Success: true, Data: m, Error: nil}, w, http.StatusCreated)
}

func validate(v entity.InputValidation) error {
	return v.Validate()
}

func queryParamInt(r *http.Request, key string, defaultValue int) (int, error) {
	v := r.FormValue(key)

	if v == "" {
		return defaultValue, nil
	}

	return strconv.Atoi(v)
}

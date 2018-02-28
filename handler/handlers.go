package handler

import (
	"log"
	"net/http"

	"github.com/Sharykhin/gl-mail-api/controller"
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

	m, c, err := controller.FailMail.GetList(r.Context(), limit, offset)

	if err != nil {
		util.SendResponse(util.Response{Success: false, Data: nil, Error: err}, w, http.StatusInternalServerError)
		return
	}

	util.SendResponse(util.Response{Success: true, Data: m, Error: nil, Meta: map[string]int64{
		"total":  c,
		"count":  int64(len(m)),
		"limit":  limit,
		"offset": offset,
	}}, w, http.StatusOK)
}

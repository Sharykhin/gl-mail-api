package handler

import (
	"log"
	"net/http"

	"strconv"

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

	m, c, err := controller.GetList(r.Context(), limit, offset)

	if err != nil {
		util.SendResponse(util.Response{Success: false, Data: nil, Error: err}, w, http.StatusInternalServerError)
		return
	}

	util.SendResponse(util.Response{Success: true, Data: map[string]interface{}{
		"mails": m,
		"total": c,
		"count": len(m),
	}, Error: nil}, w, http.StatusOK)
}

func queryParamInt(r *http.Request, key string, defaultValue int64) (int64, error) {
	v := r.FormValue(key)

	if v == "" {
		return defaultValue, nil
	}
	n, err := strconv.Atoi(v)
	return int64(n), err
}

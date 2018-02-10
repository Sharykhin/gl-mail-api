package handler

import (
	"log"
	"net/http"

	"github.com/Sharykhin/gl-mail-api/util"
)

func pong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("pong"))
	if err != nil {
		log.Fatalf("something really akward went wrong: %v", err)
	}
}

func failedMailsList(w http.ResponseWriter, r *http.Request) {
	util.SendResponse(util.Response{Success: true, Data: nil, Error: nil}, w, http.StatusOK)
}

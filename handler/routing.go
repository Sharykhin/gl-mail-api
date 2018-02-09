package handler

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Sharykhin/gl-mail-api/middleware"
)

func Handler() http.Handler {
	r := mux.NewRouter()
	r.Handle("/ping", http.HandlerFunc(pong)).Methods("GET")
	r.Handle("/failed-mails", middleware.Chain(http.HandlerFunc(failedMailsList), middleware.JWTAuth)).Methods("GET")
	return r
}

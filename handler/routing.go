package handler

import (
	"net/http"

	"github.com/Sharykhin/gl-mail-api/middleware"
	"github.com/gorilla/mux"
)

// Handler is a main router for this service
func Handler() http.Handler {
	r := mux.NewRouter()
	r.Handle("/ping", http.HandlerFunc(pong)).Methods("GET")
	r.Handle("/failed-mails", middleware.Chain(http.HandlerFunc(getFailedMailsList), middleware.JWTAuth)).Methods("GET")
	return r
}

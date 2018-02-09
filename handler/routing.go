package handler

import (
	"net/http"
	"github.com/gorilla/mux"
)

func Handler() http.Handler {
	r := mux.NewRouter()
	//r.Handle("/login", http.HandlerFunc(HandleLogin))
	r.Handle("/ping", http.HandlerFunc(pong)).Methods("GET")
	r.Handle("/login", http.HandlerFunc(login)).Methods("POST")
	return r
}

package handler

import (
	"net/http"
	"github.com/gorilla/mux"
)

func Handler() http.Handler {
	r := mux.NewRouter()
	//r.Handle("/login", http.HandlerFunc(HandleLogin))
	r.Handle("/ping", http.HandlerFunc(pong))
	return r
}

package routes

import (
	"webapi/methods"

	"github.com/gorilla/mux"
)

func RouteHandler(r *mux.Router) {
	r.HandleFunc("/blog", methods.GetYears).Methods("GET")
	r.HandleFunc("/catalog/{Year}", methods.GetYear).Methods("GET")
	r.HandleFunc("/catalog/{Year}/{Month}", methods.GetYearMonth).Methods("GET")
	r.HandleFunc("/blog/{Id}", methods.GetPost).Methods("GET")
	r.HandleFunc("comment/{id}", methods.GetComment).Methods("GET")
	r.HandleFunc("comment/{id}", methods.PostComment).Methods("POST")
}

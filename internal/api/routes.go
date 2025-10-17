package api

import (
    "net/http"
    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    r := mux.NewRouter()

    // Example route
    r.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Orders API"))
    }).Methods(http.MethodGet)


	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the E-commerce API"))
	}).Methods("GET")

    return r
}




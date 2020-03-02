package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "TEST RESPONSE"}`))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", get).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", router))
}
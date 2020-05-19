package main

import (
	"log"
	"net/http"

	"github.com/353solutions/weather/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{date}", handler.JSONGet).Methods(http.MethodGet)
	r.HandleFunc("/", handler.JSONAdd).Methods(http.MethodPost)

	addr := ":8081"
	log.Printf("server ready on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}

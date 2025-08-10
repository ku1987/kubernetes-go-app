package main

import (
	"kubernetes-go-app/pkg/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/api/items", api.GetItems).Methods("GET")
	r.HandleFunc("/api/items/{id}", api.GetItem).Methods("GET")
	r.HandleFunc("/api/items", api.CreateItem).Methods("POST")
	r.HandleFunc("/api/items/{id}", api.UpdateItem).Methods("PUT")
	r.HandleFunc("/api/items/{id}", api.DeleteItem).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

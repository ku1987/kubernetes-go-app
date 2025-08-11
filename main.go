package main

import (
	"kubernetes-go-app/pkg/api"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Print environment variables
	log.Printf("Environment Variables:")
	log.Printf("APP_ENV: %s", os.Getenv("APP_ENV"))
	log.Printf("LOG_LEVEL: %s", os.Getenv("LOG_LEVEL"))
	log.Printf("MAX_ITEMS: %s", os.Getenv("MAX_ITEMS"))

	r := mux.NewRouter()

	// API Routes
	r.HandleFunc("/api/items", api.GetItems).Methods("GET")
	r.HandleFunc("/api/items/{id}", api.GetItem).Methods("GET")
	r.HandleFunc("/api/items", api.CreateItem).Methods("POST")
	r.HandleFunc("/api/items/{id}", api.UpdateItem).Methods("PUT")
	r.HandleFunc("/api/items/{id}", api.DeleteItem).Methods("DELETE")

	// Health Check Route
	r.HandleFunc("/health", api.HealthCheck).Methods("GET")

	log.Printf("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

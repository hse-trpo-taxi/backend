// Package main provides the entry point for the taxi service backend API server.
// This service manages clients, drivers, and cars for a taxi service.
// It provides RESTful endpoints for CRUD operations on these entities.
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hse-trpo-taxi/backend/config"
	"github.com/hse-trpo-taxi/backend/database"
	"github.com/hse-trpo-taxi/backend/handlers"
)

// main initializes the taxi service backend API server.
// It loads configuration, initializes the database connection,
// sets up HTTP routes, and starts the server on the configured port.
func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	if err := database.InitDB(cfg.DatabaseDSN); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDB()

	// Setup router
	router := mux.NewRouter()

	// Client routes
	router.HandleFunc("/api/clients", handlers.GetClients).Methods("GET")
	router.HandleFunc("/api/clients/{id}", handlers.GetClient).Methods("GET")
	router.HandleFunc("/api/clients", handlers.CreateClient).Methods("POST")
	router.HandleFunc("/api/clients/{id}", handlers.UpdateClient).Methods("PUT")
	router.HandleFunc("/api/clients/{id}", handlers.DeleteClient).Methods("DELETE")

	// Driver routes
	router.HandleFunc("/api/drivers", handlers.GetDrivers).Methods("GET")
	router.HandleFunc("/api/drivers/{id}", handlers.GetDriver).Methods("GET")
	router.HandleFunc("/api/drivers", handlers.CreateDriver).Methods("POST")
	router.HandleFunc("/api/drivers/{id}", handlers.UpdateDriver).Methods("PUT")
	router.HandleFunc("/api/drivers/{id}", handlers.DeleteDriver).Methods("DELETE")

	// Car routes
	router.HandleFunc("/api/cars", handlers.GetCars).Methods("GET")
	router.HandleFunc("/api/cars/{id}", handlers.GetCar).Methods("GET")
	router.HandleFunc("/api/cars", handlers.CreateCar).Methods("POST")
	router.HandleFunc("/api/cars/{id}", handlers.UpdateCar).Methods("PUT")
	router.HandleFunc("/api/cars/{id}", handlers.DeleteCar).Methods("DELETE")

	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// Start server
	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

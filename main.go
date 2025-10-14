// Package main provides the entry point for the taxi service backend API server.
// This service manages clients, drivers, and cars for a taxi service.
// It provides RESTful endpoints for CRUD operations on these entities.
package main

import (
	"github.com/Masterminds/squirrel"
	"github.com/hse-trpo-taxi/backend/database"
	"github.com/hse-trpo-taxi/backend/server"
	"log/slog"
	"os"

	"github.com/hse-trpo-taxi/backend/config"
)

// main initializes the taxi service backend API server.
// It loads configuration, initializes the database connection,
// sets up HTTP routes, and starts the server on the configured port.
func main() {
	// Load configuration
	cfg := config.LoadConfig()
	lgr := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	pgDB, err := database.NewPgPool(cfg)

	if err != nil {
		lgr.Error(err.Error())
	}

	defer pgDB.Close()

	if errDb := database.InitDB(pgDB, lgr); errDb != nil {
		lgr.Error(errDb.Error())
	}

	srv := server.NewServer(cfg, lgr, pgDB, &builder)

	if errSrv := srv.Run(); errSrv != nil {
		lgr.Error(errSrv.Error())
	}
}

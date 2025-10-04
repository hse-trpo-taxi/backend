/*
Package backend provides a RESTful API server for a taxi service management system.

# Overview

This service manages three core entities in a taxi service:
  - Clients: Customers who request taxi rides
  - Drivers: Taxi drivers who provide transportation services
  - Cars: Vehicles associated with drivers

# Architecture

The application follows a layered architecture:

  - main.go: Application entry point and HTTP server setup
  - config/: Configuration management with environment variable support
  - database/: PostgreSQL database connection and table management
  - models/: Data structure definitions for all entities
  - handlers/: HTTP request handlers implementing RESTful API endpoints

# API Endpoints

## Client Management

	GET    /api/clients      - List all clients
	GET    /api/clients/{id} - Get client by ID
	POST   /api/clients      - Create new client
	PUT    /api/clients/{id} - Update client
	DELETE /api/clients/{id} - Delete client

## Driver Management

	GET    /api/drivers      - List all drivers
	GET    /api/drivers/{id} - Get driver by ID
	POST   /api/drivers      - Create new driver
	PUT    /api/drivers/{id} - Update driver
	DELETE /api/drivers/{id} - Delete driver

## Car Management

	GET    /api/cars         - List all cars
	GET    /api/cars/{id}    - Get car by ID
	POST   /api/cars         - Create new car
	PUT    /api/cars/{id}    - Update car
	DELETE /api/cars/{id}    - Delete car

## Health Check

	GET    /health           - Service health status

# Configuration

The service uses environment variables for configuration:

Database Configuration:
  - DATABASE_URL: Complete PostgreSQL connection string
  - DB_HOST: Database host (default: localhost)
  - DB_PORT: Database port (default: 5432)
  - DB_USER: Database username (default: postgres)
  - DB_PASSWORD: Database password (default: postgres)
  - DB_NAME: Database name (default: taxi)
  - DB_SSLMODE: SSL mode (default: disable)

Server Configuration:
  - SERVER_PORT: HTTP server port (default: 8080)

# Database Schema

The service uses PostgreSQL with the following tables:

	clients:
	  - id (SERIAL PRIMARY KEY)
	  - name (VARCHAR(255) NOT NULL)
	  - phone (VARCHAR(50) NOT NULL)
	  - email (VARCHAR(255) NOT NULL)
	  - created_at (TIMESTAMP DEFAULT CURRENT_TIMESTAMP)
	  - updated_at (TIMESTAMP DEFAULT CURRENT_TIMESTAMP)

	drivers:
	  - id (SERIAL PRIMARY KEY)
	  - name (VARCHAR(255) NOT NULL)
	  - phone (VARCHAR(50) NOT NULL)
	  - license_number (VARCHAR(50) NOT NULL)
	  - rating (REAL DEFAULT 0.0)
	  - created_at (TIMESTAMP DEFAULT CURRENT_TIMESTAMP)
	  - updated_at (TIMESTAMP DEFAULT CURRENT_TIMESTAMP)

	cars:
	  - id (SERIAL PRIMARY KEY)
	  - driver_id (INTEGER NOT NULL, FOREIGN KEY to drivers.id)
	  - brand (VARCHAR(100) NOT NULL)
	  - model (VARCHAR(100) NOT NULL)
	  - year (INTEGER NOT NULL)
	  - license_plate (VARCHAR(50) NOT NULL)
	  - color (VARCHAR(50) NOT NULL)
	  - created_at (TIMESTAMP DEFAULT CURRENT_TIMESTAMP)
	  - updated_at (TIMESTAMP DEFAULT CURRENT_TIMESTAMP)

# Usage Example

Starting the server:

	go run main.go

Creating a client:

	curl -X POST http://localhost:8080/api/clients \
	  -H "Content-Type: application/json" \
	  -d '{"name":"John Doe","phone":"+1234567890","email":"john@example.com"}'

Listing all drivers:

	curl http://localhost:8080/api/drivers

# Dependencies

  - github.com/gorilla/mux: HTTP router and URL matcher
  - github.com/lib/pq: PostgreSQL driver for Go

# Error Handling

All endpoints return appropriate HTTP status codes:
  - 200: Success
  - 201: Created
  - 204: No Content (for successful deletions)
  - 400: Bad Request (invalid input)
  - 404: Not Found
  - 500: Internal Server Error (database or server errors)

Error responses include descriptive error messages in the response body.
*/
package main

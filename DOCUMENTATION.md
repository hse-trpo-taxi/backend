# Taxi Service Backend API Documentation

This document explains how to view and use the comprehensive godoc documentation that has been added to this project.

## Viewing Documentation

### 1. Command Line Documentation

View all package documentation:
```bash
go doc -all
```

View specific package documentation:
```bash
go doc config
go doc database
go doc models
go doc handlers
```

View specific type documentation:
```bash
go doc models.Client
go doc models.Driver
go doc models.Car
```

View specific function documentation:
```bash
go doc handlers.GetClients
go doc config.LoadConfig
go doc database.InitDB
```

### 2. Web-based Documentation

Start a local godoc server:
```bash
godoc -http=:6060
```

Then open your browser and visit: `http://localhost:6060/pkg/github.com/hse-trpo-taxi/backend/`

### 3. Online Documentation

If this package is publicly available, you can view the documentation online at:
`https://pkg.go.dev/github.com/hse-trpo-taxi/backend`

## Documentation Structure

The project includes comprehensive documentation for:

- **Package-level documentation**: Overview of the entire taxi service system
- **Type documentation**: Detailed descriptions of all data structures
- **Function documentation**: Complete descriptions of all public functions
- **API documentation**: RESTful endpoint specifications
- **Configuration documentation**: Environment variable descriptions
- **Database schema documentation**: Table structures and relationships

## Key Documentation Features

- **API Endpoints**: Complete REST API specification with HTTP methods and status codes
- **Configuration Guide**: Environment variables and default values
- **Database Schema**: Table structures and relationships
- **Usage Examples**: Sample requests and responses
- **Error Handling**: HTTP status codes and error responses
- **Dependencies**: Required external packages

## Documentation Standards

All documentation follows Go's godoc conventions:
- Package comments start with "Package [name]"
- Function comments start with the function name
- Type comments describe the purpose and usage
- Examples are provided where helpful
- HTTP status codes are documented for API endpoints

## Generating Documentation

To regenerate documentation after code changes:
```bash
go doc -all > documentation.txt
```

The documentation is automatically generated from the source code comments and stays up-to-date with the codebase.
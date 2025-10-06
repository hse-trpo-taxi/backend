package server

import "C"
import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/gorilla/mux"
	"github.com/hse-trpo-taxi/backend/config"
	"github.com/hse-trpo-taxi/backend/handlers"
	"github.com/hse-trpo-taxi/backend/repositories"
	"github.com/hse-trpo-taxi/backend/usecases"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/sync/errgroup"
	"log/slog"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

const (
	GracefulShutdownTimeOut = 5
	ServerTimeOut           = 3
	ServerMaxAge            = 300
)

type Server struct {
	config  *config.Config
	logger  *slog.Logger
	pgDB    *pgxpool.Pool
	builder *squirrel.StatementBuilderType
}

func NewServer(
	config *config.Config,
	logger *slog.Logger,
	pgDB *pgxpool.Pool,
	builder *squirrel.StatementBuilderType,
) *Server {
	return &Server{
		config:  config,
		logger:  logger,
		pgDB:    pgDB,
		builder: builder,
	}
}

func (server *Server) Run() error {
	// Setup router
	router := mux.NewRouter()

	if err := server.PrepareHandlers(router); err != nil {
		return err
	}

	mainCtx, shutdown := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer shutdown()

	httpServer := &http.Server{
		Addr: server.config.ServerPort,
		Handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			ctx := request.Context()
			select {
			case <-ctx.Done():
				time.Sleep(time.Duration(GracefulShutdownTimeOut) * time.Second)
				writer.WriteHeader(http.StatusInternalServerError)
				return
			default:
				router.ServeHTTP(writer, request)
			}
		}),
		ReadHeaderTimeout: ServerTimeOut * time.Second,
		BaseContext: func(_ net.Listener) context.Context {
			return mainCtx
		},
	}

	g, ctx := errgroup.WithContext(mainCtx)
	g.Go(func() error {
		server.logger.With(
			slog.String("port", server.config.ServerPort),
		).Info("Server running on port")

		return httpServer.ListenAndServe()
	})

	g.Go(func() error {
		<-ctx.Done()
		server.logger.Warn("Shutting down server...")

		return httpServer.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		return err
	}

	server.logger.Info("Server shut down")

	return nil
}

func (server *Server) PrepareHandlers(router *mux.Router) error {

	carRepo := repositories.NewCarRepository(server.pgDB, server.builder)

	carUS := usecases.NewCarUseCase(carRepo)

	carHandler := handlers.NewCarHandler(carUS, server.logger)

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
	router.HandleFunc("/api/cars", carHandler.GetCars).Methods("GET")
	router.HandleFunc("/api/cars/{id}", carHandler.GetCarById).Methods("GET")
	router.HandleFunc("/api/cars", handlers.CreateCar).Methods("POST")
	router.HandleFunc("/api/cars/{id}", handlers.UpdateCar).Methods("PUT")
	router.HandleFunc("/api/cars/{id}", carHandler.DeleteCar).Methods("DELETE")

	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	return nil
}

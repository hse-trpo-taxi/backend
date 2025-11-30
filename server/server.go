package server

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/gorilla/mux"
	"github.com/hse-trpo-taxi/backend/config"
	"github.com/hse-trpo-taxi/backend/handlers"
	"github.com/hse-trpo-taxi/backend/repositories"
	"github.com/hse-trpo-taxi/backend/usecases"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/sync/errgroup"
)

const (
	GracefulShutdownTimeOut = 5
	ServerTimeOut           = 3
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
	clientRepo := repositories.NewClientRepository(server.pgDB, server.builder)
	driverRepo := repositories.NewDriverRepository(server.pgDB, server.builder)
	supportRepo := repositories.NewSupportRepository(server.pgDB, server.builder)
	orderRepo := repositories.NewOrderRepository(server.pgDB, server.builder)

	carUS := usecases.NewCarUseCase(carRepo)
	clientUS := usecases.NewClientUseCase(clientRepo)
	driverUS := usecases.NewDriverUseCase(driverRepo, carRepo)
	supportUS := usecases.NewSupportUseCase(supportRepo)
	orderUS := usecases.NewOrderUseCase(orderRepo)
	userUS := usecases.NewUserUseCase()

	carHandler := handlers.NewCarHandler(carUS, server.logger)
	clientHandler := handlers.NewClientHandler(clientUS, server.logger)
	driverHandler := handlers.NewDriverHandler(driverUS, server.logger)
	supportHandler := handlers.NewSupportHandler(supportUS, server.logger)
	orderHandler := handlers.NewOrderHandler(orderUS, server.logger)
	userHandler := handlers.NewUserHandler(userUS, server.logger)

	router.HandleFunc("/api/clients", clientHandler.GetClients).Methods("GET")
	router.HandleFunc("/api/clients/{id}", clientHandler.GetClientById).Methods("GET")
	router.HandleFunc("/api/clients", clientHandler.CreateClient).Methods("POST")
	router.HandleFunc("/api/clients/{id}", clientHandler.UpdateClient).Methods("PUT")
	router.HandleFunc("/api/clients/{id}", clientHandler.DeleteClient).Methods("DELETE")

	// Driver routes
	router.HandleFunc("/api/drivers", driverHandler.GetDrivers).Methods("GET")
	router.HandleFunc("/api/drivers/{id}", driverHandler.GetDriverById).Methods("GET")
	router.HandleFunc("/api/drivers", driverHandler.CreateDriver).Methods("POST")
	router.HandleFunc("/api/drivers/{id}", driverHandler.UpdateDriver).Methods("PUT")
	router.HandleFunc("/api/drivers/{id}", driverHandler.DeleteDriver).Methods("DELETE")
	router.HandleFunc("/api/drivers/meanScore", driverHandler.GetDriversMeanScore).Methods("GET")
	router.HandleFunc("/api/drivers/coords", driverHandler.GetDriversCoords).Methods("GET")
	router.HandleFunc("/api/drivers/stat", driverHandler.GetDriversStat).Methods("POST")
	router.HandleFunc("/api/drivers/{id}/schedule", driverHandler.GetDriverSchedule).Methods("POST")

	// Car routes
	router.HandleFunc("/api/cars", carHandler.GetCars).Methods("GET")
	router.HandleFunc("/api/cars/{id}", carHandler.GetCarById).Methods("GET")
	router.HandleFunc("/api/cars", carHandler.CreateCar).Methods("POST")
	router.HandleFunc("/api/cars/{id}", carHandler.UpdateCar).Methods("PUT")
	router.HandleFunc("/api/cars/{id}", carHandler.DeleteCar).Methods("DELETE")

	// Support routes
	router.HandleFunc("/api/support/requests", supportHandler.GetRecent).Methods("GET")

	// User routes
	router.HandleFunc("/api/user/current", userHandler.GetCurrent).Methods("GET")

	// Order routes
	router.HandleFunc("/api/order/weekStat", orderHandler.GetWeekStat).Methods("GET")
	router.HandleFunc("/api/order/currentStat", orderHandler.GetCurrentStats).Methods("GET")

	// Orders list with pagination
	router.HandleFunc("/api/orders/list", orderHandler.GetOrdersList).Methods("POST")

	// Driver-specific order routes
	router.HandleFunc("/api/drivers/{id}/orders", orderHandler.GetDriverOrders).Methods("POST")
	router.HandleFunc("/api/drivers/{id}/order/current", orderHandler.GetDriverCurrentOrder).Methods("GET")

	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	return nil
}

package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/SamuraiAkira/warehouse-management-service/internal/app/config"
	"github.com/SamuraiAkira/warehouse-management-service/internal/app/repository/postgresql"
	"github.com/SamuraiAkira/warehouse-management-service/internal/app/service"
	"github.com/SamuraiAkira/warehouse-management-service/internal/pkg/logging"

	deliveryhttp "github.com/SamuraiAkira/warehouse-management-service/internal/app/delivery/http"
)

func main() {
	logger := logging.NewLogger()
	defer logger.Sync()

	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Failed to load config", zap.Error(err))
	}

	ctx := context.Background()
	var dbPool *pgxpool.Pool
	for i := 0; i < 5; i++ {
		dbPool, err = pgxpool.New(ctx, buildDBConnString(cfg.Postgres))
		if err == nil {
			break
		}
		logger.Info("Waiting for PostgreSQL...", zap.Int("attempt", i+1))
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		logger.Fatal("Failed to connect to database after retries", zap.Error(err))
	}
	defer dbPool.Close()

	warehouseRepo := postgresql.NewWarehouseRepository(dbPool)
	warehouseService := service.NewWarehouseService(warehouseRepo)
	warehouseHandler := deliveryhttp.NewWarehouseHandler(warehouseService)

	router := mux.NewRouter()

	router.HandleFunc("/api/health", deliveryhttp.HealthCheck).Methods("GET")
	router.HandleFunc("/api/warehouses", warehouseHandler.List).Methods("GET")
	router.HandleFunc("/api/warehouses", warehouseHandler.Create).Methods("POST")

	server := &http.Server{
		Addr:         cfg.HTTP.Host + ":" + cfg.HTTP.Port,
		Handler:      logging.RequestIDMiddleware(logging.LoggingMiddleware(logger, router)),
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
	}

	go func() {
		logger.Info("Starting server", zap.String("address", server.Addr))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Server failed", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), cfg.HTTP.ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown failed", zap.Error(err))
	}

	logger.Info("Server stopped gracefully")
}

func buildDBConnString(cfg config.PostgresConfig) string {
	return "postgres://" + cfg.Username + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.DBName + "?sslmode=" + cfg.SSLMode
}

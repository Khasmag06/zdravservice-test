package app

import (
	"context"
	"github.com/khasmag06/zdravservice-test/config"
	"github.com/khasmag06/zdravservice-test/internal/controller/api"
	productsRepo "github.com/khasmag06/zdravservice-test/internal/repo/product/postgres"
	"github.com/khasmag06/zdravservice-test/internal/service/product"
	"github.com/khasmag06/zdravservice-test/pkg/httpserver"
	"github.com/khasmag06/zdravservice-test/pkg/logger"
	"github.com/khasmag06/zdravservice-test/pkg/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	l, err := logger.New(cfg.Logger.LogFilePath, cfg.Logger.Level)
	if err != nil {
		log.Fatalf("failed to build logger: %s", err)
	}
	defer func() { _ = l.Sync() }()

	ctx := context.Background()
	db, err := postgres.NewDB(ctx, cfg.PG)
	if err != nil {
		l.Fatalf("failed to connect to postgres db: %s", err)
	}
	defer db.Close()

	repo := productsRepo.New(db.Pool)
	service := product.New(repo)

	// HTTP Server
	l.Info("Starting api server...")
	handler := api.NewHandler(service, l)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Errorf("app - Run - httpServer.Notify: %w", err)
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Errorf("app - Run - httpServer.Shutdown: %w", err)
	}
}

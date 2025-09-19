package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/egustafson/fintrax/pkg/config"
	"github.com/egustafson/fintrax/pkg/db"
)

func Start() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// hook signals for shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		slog.Info(fmt.Sprintf("received signal: %s", sig.String()))
		cancel()
	}()

	config, ctx, err := config.InitServerConfig(ctx)
	if err != nil {
		slog.Error("configuration error", "error", err)
		return err
	}
	if err := serverInitLogging(config); err != nil {
		slog.Error("logging initialization failed", "error", err)
		return err
	}
	if err := db.Init(config.DB); err != nil {
		slog.Error("db connection failed", "error", err)
		return err
	}
	defer db.Shutdown()

	router := gin.Default()

	// TODO:  setup API
	// TODO:  setup UI

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: router,
	}

	go func() {
		slog.Info("http server listening", "port", config.Port)
		if err = srv.ListenAndServe(); err != http.ErrServerClosed {
			slog.Warn("http server shutdown", "error", err)
			cancel() // tear everything down
		}
	}()

	<-ctx.Done() // block until the context is canceled
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Warn("http server failed to properly shutdown", "error", err)
		srv.Close() // force closure
		return err
	}
	slog.Info("server shutdown complete")
	return nil
}

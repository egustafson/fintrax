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
	sloggin "github.com/samber/slog-gin"

	"github.com/egustafson/fintrax/api"
	"github.com/egustafson/fintrax/pkg/config"
	"github.com/egustafson/fintrax/pkg/dao"
	"github.com/egustafson/fintrax/pkg/mx"
)

func Start(flags *config.Flags) error {

	// logging is pre-initialized via env vars and default values.
	// see: logging.go

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = hookSignals(ctx) // hook ctrl-c for shutdown

	// Initialize the root managed object
	rootMO := mx.NewBaseMO()
	rootMO.SetState("type-id", "root-mo")
	ctx = context.WithValue(ctx, "root-mo", rootMO)

	config, ctx, err := config.InitServerConfig(ctx, flags)
	if err != nil {
		slog.Error("configuration error", "error", err)
		return err
	}

	daoFactory, err := dao.NewFactory(config.DB)
	if err != nil {
		slog.Error("failed to create dao factory", "error", err)
		return err
	}
	defer daoFactory.Shutdown()
	rootMO.Attach("dao-factory", daoFactory)

	// Run the HTTP server
	defer slog.Info("server shutdown complete")
	return serveHTTP(ctx, daoFactory, config)
}

func hookSignals(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		slog.Info(fmt.Sprintf("received signal: %s", sig.String()))
		cancel()
	}()
	return ctx
}

func serveHTTP(
	ctx context.Context,
	daoFactory dao.Factory,
	config *config.ServerConfig) error {

	ctx, cancel := context.WithCancel(ctx)

	router := gin.New()
	router.Use(sloggin.New(rootLogger))
	router.Use(gin.Recovery())

	api.InitAPI(ctx, router)
	// TODO:  setup UI

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: router,
	}

	go func() {
		slog.Info("http server listening", "port", config.Port)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
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
	return nil
}

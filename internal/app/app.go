package app

import (
	_ "bank/api/v1"
	"bank/internal/config"
	accountcontroller "bank/internal/controllers/account"
	"bank/internal/storages"
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

// @title           Bank
// @version         1.0
// @description     Online Bank API

// @BasePath  /api/v1
func Run(cfg config.Config) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	accountStorage := storages.NewAccountStorageInMemory()
	accountController := accountcontroller.NewAccountController(accountStorage)

	engine := gin.New()
	engine.Use(sloggin.New(logger))
	engine.Use(gin.Recovery())
	engine.GET("api/v1/swagger/*any", ginswagger.WrapHandler(swaggerfiles.Handler))
	accountController.RegisterRoutes(engine)

	srv := &http.Server{
		Addr:    cfg.HTTPServer.IpAddress + ":" + cfg.HTTPServer.Port,
		Handler: engine.Handler(),
	}

	slog.Info("Starting server ...")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server listen: %s\n", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Server Shutdown:", "error", err)
		os.Exit(1)
	}

	select {
	case <-ctx.Done():
		slog.Info("timeout of 5 seconds.")
	}
	slog.Info("Server exiting")
}

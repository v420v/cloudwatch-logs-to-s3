package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/v420v/cloudwatch-logs/internal/config"
	"github.com/v420v/cloudwatch-logs/internal/controller"
	"github.com/v420v/cloudwatch-logs/internal/logger"
	"github.com/v420v/cloudwatch-logs/internal/middleware"
	"github.com/v420v/cloudwatch-logs/internal/router"
	"go.uber.org/zap"
)

func main() {
	l := logger.InitLogger()
	c := controller.NewController(l)
	m := middleware.NewMiddleware(l)
	r := router.NewRouter(m, c)

	cfg := config.LoadConfig()

	defer l.Sync()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	server := &http.Server{
		Addr:    cfg.Port,
		Handler: r,
	}

	go func() {
		l.Info("server_starting", zap.String("port", cfg.Port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Fatal("server_failed_to_start", zap.Error(err))
		}
	}()

	<-quit
	l.Info("shutdown_initiated")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		l.Error("server_forced_to_shutdown", zap.Error(err))
	}

	l.Info("server_stopped")
}

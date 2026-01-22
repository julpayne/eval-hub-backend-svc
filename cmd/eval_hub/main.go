package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.ibm.com/julpayne/eval-hub-backend-svc/cmd/eval_hub/server"
	"github.ibm.com/julpayne/eval-hub-backend-svc/internal/config"
	"github.ibm.com/julpayne/eval-hub-backend-svc/internal/logging"
)

var (
	// Version is set during the compilation
	Version string
	// Build is set during the compilation
	Build string
	// BuildDate is set during the compilation
	BuildDate string
)

func main() {
	// TODO write fatal errors to the error file and close down the server

	// Create logger once for all requests
	logger, err := logging.NewLogger()
	if err != nil {
		// we do this as no point trying to continue
		startUpFailed(nil, err, "Failed to create service logger", logging.FallbackLogger())
	}

	serviceConfig, err := config.LoadConfig(logger, Version, Build, BuildDate)
	if err != nil {
		// we do this as no point trying to continue
		startUpFailed(nil, err, "Failed to create service config", logger)
	}

	srv, err := server.NewServer(logger, serviceConfig)
	if err != nil {
		// we do this as no point trying to continue
		startUpFailed(nil, err, "Failed to create server", logger)
	}

	// Start server in a goroutine
	go func() {
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			// we do this as no point trying to continue
			startUpFailed(nil, err, "Server failed to start", logger)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	// Create a context with timeout for graceful shutdown
	waitForShutdown := 30 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), waitForShutdown)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown", "error", err.Error(), "timeout", waitForShutdown)
		log.Fatal("Server forced to shutdown:", err)
	} else {
		logger.Info("Server shutdown gracefully")
	}
}

func startUpFailed(conf *config.Config, err error, msg string, logger *slog.Logger) {
	termErr := server.SetTerminationMessage(server.GetTerminationFile(conf, logger), fmt.Sprintf("%s: %s", msg, err.Error()), logger)
	if termErr != nil {
		logger.Error("Failed to set termination message", "message", msg, "error", termErr.Error())
		log.Println(termErr.Error())
	}
	log.Fatal(err)
}

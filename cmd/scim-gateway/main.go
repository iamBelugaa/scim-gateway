package main

import (
	"fmt"
	"log"

	"github.com/iamBelugaa/scim-gateway/internal/config"
	"github.com/iamBelugaa/scim-gateway/internal/server"
	"github.com/iamBelugaa/scim-gateway/pkg/logger"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln("application error", err)
	}
}

// run loads configuration, initializes the logger and HTTP server,
// and starts listening for HTTP requests.
func run() error {
	// Load configuration from environment variables.
	conf := config.Load()

	// Initialize structured logger with configuration.
	logger, err := logger.NewWithConfig(conf.Application, conf.Logging)
	if err != nil {
		return fmt.Errorf("failed to construct logger : %w", err)
	}

	// Ensure logs are flushed on exit.
	defer func() {
		if loggerCloseErr := logger.Close(); loggerCloseErr != nil {
			err = fmt.Errorf("failed to flush buffered log entries : %v", err)
		}
	}()

	logger.Infow(fmt.Sprintf("starting %s service", conf.Application.Service))

	// Start serving HTTP requests.
	server := server.NewWithConfig(logger, conf)
	server.ListenAndServe()

	// Wait for shutdown signal or error and gracefully shut down the server.
	return server.Shutdown()
}

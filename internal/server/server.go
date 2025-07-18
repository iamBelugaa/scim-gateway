package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	goahttp "goa.design/goa/v3/http"

	genscimserver "github.com/iamBelugaa/scim-gateway/gen/http/scim/server"
	genscim "github.com/iamBelugaa/scim-gateway/gen/scim"

	"github.com/iamBelugaa/scim-gateway/internal/config"
	"github.com/iamBelugaa/scim-gateway/internal/services/scimsvc"
	"github.com/iamBelugaa/scim-gateway/pkg/logger"
)

// server encapsulates the application configuration,
// logger, HTTP server instance, and error channel.
type server struct {
	cfg         *config.Config // Application configuration
	log         *logger.Logger // Application logger
	httpServer  *http.Server   // Underlying HTTP server
	serverError chan error     // Channel for capturing async server errors
}

func NewWithConfig(logger *logger.Logger, cfg *config.Config) *server {
	// Initialize scim service and endpoints.
	scimsvc := scimsvc.NewService(logger)
	scimEndpoints := genscim.NewEndpoints(scimsvc)

	// Create Goa HTTP multiplexer.
	mux := goahttp.NewMuxer()

	// Setup and mount scim HTTP handlers.
	scimHandlers := genscimserver.New(scimEndpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
	genscimserver.Mount(mux, scimHandlers)

	// Log mounted scim endpoints.
	for _, mount := range scimHandlers.Mounts {
		log.Printf("%q mounted on %s %s", mount.Method, mount.Verb, mount.Pattern)
	}

	return &server{
		cfg:         cfg,
		log:         logger,
		serverError: make(chan error, 1),
		httpServer: &http.Server{
			Handler:      mux,
			IdleTimeout:  cfg.Server.IdleTimeout,
			ReadTimeout:  cfg.Server.ReadTimeout,
			WriteTimeout: cfg.Server.WriteTimeout,
			Addr:         fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		},
	}
}

// ListenAndServe starts the HTTP server.
func (s *server) ListenAndServe() error {
	go func() {
		s.log.Infow("starting scim http server", "address", fmt.Sprintf("%s:%d", s.cfg.Server.Host, s.cfg.Server.Port))
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.serverError <- err
		}
	}()
	return nil
}

// Shutdown listens for termination signals or server errors
// and performs a graceful shutdown of the HTTP server.
func (s *server) Shutdown() error {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	// Handle server startup error.
	case err := <-s.serverError:
		s.log.Infow("received server error", "error", err)
		return fmt.Errorf("server error: %w", err)

	// Handle OS shutdown signal.
	case sig := <-signalCh:
		s.log.Infow("shutting down server signal received", "signal", sig)
		s.log.Infow("initiating graceful shutdown", "service", s.cfg.Application.Service)

		// Create context with timeout for graceful shutdown.
		shutdownCtx, cancel := context.WithTimeout(context.Background(), s.cfg.Server.ShutdownTimeout)
		defer cancel()

		// Attempt graceful shutdown.
		if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}

		s.log.Infow("graceful shutdown completed successfully", "service", s.cfg.Application.Service)
	}

	return nil
}

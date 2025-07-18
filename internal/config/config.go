// Package config defines configuration types used across the service.
package config

import (
	"time"
)

// Environment defines the type for representing different runtime environments.
type Environment string

// Supported environment constants.
var (
	EnvironmentProduction  Environment = "PRODUCTION"
	EnvironmentDevelopment Environment = "DEVELOPMENT"
)

// String returns string representation of the Environment.
func (e Environment) String() string {
	switch e {
	case EnvironmentProduction:
		return "production"
	case EnvironmentDevelopment:
		return "development"
	default:
		return "development"
	}
}

// Logging holds settings for how logging should behave in different environments.
type Logging struct {
	Level       string   `json:"level"`       // Log level: e.g., "debug", "info", "warn", "error".
	OutputPaths []string `json:"outputPaths"` // List of output destinations, e.g., "stderr", "stdout", or file paths.
}

// Application holds application specific metadata.
type Application struct {
	Service     string      `json:"service"`     // Service name identifier.
	Version     string      `json:"version"`     // Application version.
	Environment Environment `json:"environment"` // Runtime environment (production, development, etc.).
}

// Server holds HTTP server configuration parameters.
type Server struct {
	Host            string        `json:"host"`            // Host address to bind the server (e.g., "0.0.0.0").
	Port            uint          `json:"port"`            // Port number on which the server listens.
	ReadTimeout     time.Duration `json:"readTimeout"`     // Maximum duration for reading the entire request.
	WriteTimeout    time.Duration `json:"writeTimeout"`    // Maximum duration before timing out writes of the response.
	IdleTimeout     time.Duration `json:"idleTimeout"`     // Maximum amount of time to wait for the next request.
	ShutdownTimeout time.Duration `json:"shutdownTimeout"` // Grace period for server shutdown.
}

// Config is the top level struct that aggregates all configuration domains.
type Config struct {
	Server      *Server      `json:"server"`      // HTTP server configuration.
	Logging     *Logging     `json:"logging"`     // Logging configuration.
	Application *Application `json:"application"` // Application metadata and environment.
}

// Load gathers configuration values from environment variables,
// applying sensible defaults if variables are not set.
func Load() *Config {
	return &Config{
		Server: &Server{
			Host:            GetEnvString("SERVER_HOST", "0.0.0.0"),
			Port:            uint(GetEnvInt("SERVER_PORT", 8080)),
			ReadTimeout:     GetEnvDuration("SERVER_READ_TIMEOUT", time.Second*15),
			WriteTimeout:    GetEnvDuration("SERVER_WRITE_TIMEOUT", time.Second*15),
			IdleTimeout:     GetEnvDuration("SERVER_IDLE_TIMEOUT", time.Second*30),
			ShutdownTimeout: GetEnvDuration("SERVER_SHUTDOWN_TIMEOUT", time.Second*30),
		},
		Logging: &Logging{
			Level:       GetEnvString("LOG_LEVEL", "info"),
			OutputPaths: GetEnvSlice("LOG_OUTPUT_PATHS", []string{"stderr"}),
		},
		Application: &Application{
			Service:     GetEnvString("SERVICE", "scim-gateway"),
			Version:     GetEnvString("APPLICATION_VERSION", "1.0.0"),
			Environment: ToEnvironment(GetEnvString("ENVIRONMENT", "development")),
		},
	}
}

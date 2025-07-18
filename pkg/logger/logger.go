package logger

import (
	"fmt"
	"os"

	"github.com/iamBelugaa/scim-gateway/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger wraps zap.SugaredLogger.
type Logger struct {
	*zap.SugaredLogger
}

// NewWithConfig initializes a new logger using the provided service name,
// version, environment, and configuration.
func NewWithConfig(appConfig *config.Application, logConfig *config.Logging) (*Logger, error) {
	// Parse the log level from the config string.
	level, err := zapcore.ParseLevel(logConfig.Level)
	if err != nil {
		return nil, fmt.Errorf("failed to parse log level : %w", err)
	}

	// Use production config and encoder as defaults.
	zapConfig := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()

	// Switch to development settings if applicable.
	if appConfig.Environment != config.EnvironmentProduction {
		zapConfig = zap.NewDevelopmentConfig()
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		// Add color for easier reading in development logs.
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// Apply common configurations.
	zapConfig.Encoding = "json"
	zapConfig.DisableCaller = false
	zapConfig.DisableStacktrace = false
	zapConfig.Level = zap.NewAtomicLevelAt(level)
	zapConfig.OutputPaths = append(zapConfig.OutputPaths, logConfig.OutputPaths...)

	// Add service level fields to every log entry.
	zapConfig.InitialFields = map[string]any{
		"pid":         os.Getpid(),
		"service":     appConfig.Service,
		"version":     appConfig.Version,
		"environment": appConfig.Environment,
	}

	// Customize encoder keys and format.
	zapConfig.EncoderConfig = encoderConfig
	zapConfig.EncoderConfig.LevelKey = "level"
	zapConfig.EncoderConfig.CallerKey = "caller"
	zapConfig.EncoderConfig.TimeKey = "timestamp"
	zapConfig.EncoderConfig.MessageKey = "message"
	zapConfig.EncoderConfig.StacktraceKey = "stacktrace"
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	return &Logger{
		SugaredLogger: zap.Must(
			zapConfig.Build(zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel)),
		).Sugar(),
	}, nil
}

// Close ensures that any buffered logs are flushed to the output.
func (l *Logger) Close() error {
	return l.Sync()
}

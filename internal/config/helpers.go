package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// GetEnvString retrieves the value of the environment variable named by `key`.
// If the variable is not set, it returns the provided `fallback` value.
func GetEnvString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}

// GetEnvInt retrieves an environment variable and converts it to an integer.
// If the variable is not set or the conversion fails, it returns the `fallback` value.
func GetEnvInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return intVal
}

// GetEnvDuration retrieves an environment variable and parses it as a time.Duration.
// The string must follow Go's duration format (e.g., "5s", "1h").
// If the variable is not set or the format is invalid, it returns the `fallback`.
func GetEnvDuration(key string, fallback time.Duration) time.Duration {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	durationVal, err := time.ParseDuration(val)
	if err != nil {
		return fallback
	}
	return durationVal
}

// GetEnvSlice retrieves an environment variable and splits it into a slice of strings,
// using commas as the delimiter. If the variable is not set, it returns the `fallback` slice.
func GetEnvSlice(key string, fallback []string) []string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	// Split the environment variable by commas into a string slice.
	parts := strings.Split(val, ",")
	return parts
}

// ToEnvironment converts a string representation of an environment name
// into a predefined Environment type. Useful for normalizing input such as
// "prod", "production", "dev", "development", etc.
//
// Defaults to EnvironmentDevelopment if input is unrecognized.
func ToEnvironment(str string) Environment {
	switch strings.ToLower(str) {
	case "prod", "production":
		return EnvironmentProduction
	case "dev", "develop", "development", "local":
		return EnvironmentDevelopment
	default:
		return EnvironmentDevelopment
	}
}

package logger

import (
	"context"
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger levels
const (
	LevelTrace = "trace"
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"
	LevelFatal = "fatal"
	LevelPanic = "panic"
)

// Config holds logger configuration
type Config struct {
	Level       string
	Environment string
	Service     string
	Output      io.Writer
	Pretty      bool
}

// Logger wraps zerolog.Logger with additional functionality
type Logger struct {
	logger zerolog.Logger
	config Config
}

// New creates a new structured logger
func New(config Config) *Logger {
	// Set default values
	if config.Level == "" {
		config.Level = LevelInfo
	}
	if config.Environment == "" {
		config.Environment = "development"
	}
	if config.Output == nil {
		config.Output = os.Stdout
	}

	// Configure zerolog
	var output io.Writer = config.Output

	// Use pretty logging for development
	if config.Pretty || config.Environment == "development" {
		output = zerolog.ConsoleWriter{
			Out:        config.Output,
			TimeFormat: time.RFC3339,
		}
	}

	// Set global log level
	level := parseLevel(config.Level)
	zerolog.SetGlobalLevel(level)

	// Create logger with common fields
	logger := zerolog.New(output).With().
		Timestamp().
		Str("service", config.Service).
		Str("environment", config.Environment).
		Logger()

	return &Logger{
		logger: logger,
		config: config,
	}
}

// NewFromEnv creates a logger from environment variables
func NewFromEnv(serviceName string) *Logger {
	config := Config{
		Level:       getEnv("LOG_LEVEL", LevelInfo),
		Environment: getEnv("APP_ENV", "development"),
		Service:     serviceName,
		Pretty:      getEnv("LOG_PRETTY", "true") == "true",
	}

	return New(config)
}

// parseLevel converts string level to zerolog level
func parseLevel(level string) zerolog.Level {
	switch strings.ToLower(level) {
	case LevelTrace:
		return zerolog.TraceLevel
	case LevelDebug:
		return zerolog.DebugLevel
	case LevelInfo:
		return zerolog.InfoLevel
	case LevelWarn:
		return zerolog.WarnLevel
	case LevelError:
		return zerolog.ErrorLevel
	case LevelFatal:
		return zerolog.FatalLevel
	case LevelPanic:
		return zerolog.PanicLevel
	default:
		return zerolog.InfoLevel
	}
}

// getEnv gets environment variable with default
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Info logs an info message
func (l *Logger) Info() *zerolog.Event {
	return l.logger.Info()
}

// Debug logs a debug message
func (l *Logger) Debug() *zerolog.Event {
	return l.logger.Debug()
}

// Warn logs a warning message
func (l *Logger) Warn() *zerolog.Event {
	return l.logger.Warn()
}

// Error logs an error message
func (l *Logger) Error() *zerolog.Event {
	return l.logger.Error()
}

// Fatal logs a fatal message and exits
func (l *Logger) Fatal() *zerolog.Event {
	return l.logger.Fatal()
}

// Panic logs a panic message and panics
func (l *Logger) Panic() *zerolog.Event {
	return l.logger.Panic()
}

// Trace logs a trace message
func (l *Logger) Trace() *zerolog.Event {
	return l.logger.Trace()
}

// With creates a child logger with additional fields
func (l *Logger) With() zerolog.Context {
	return l.logger.With()
}

// Logger returns the underlying zerolog.Logger
func (l *Logger) Logger() zerolog.Logger {
	return l.logger
}

// Context methods for request tracing
func (l *Logger) WithContext(ctx context.Context) *Logger {
	return &Logger{
		logger: l.logger.With().Logger(),
		config: l.config,
	}
}

// HTTP request logging methods
func (l *Logger) LogHTTPRequest(method, path, remoteAddr string, statusCode int, duration time.Duration) {
	l.logger.Info().
		Str("type", "http_request").
		Str("method", method).
		Str("path", path).
		Str("remote_addr", remoteAddr).
		Int("status_code", statusCode).
		Dur("duration", duration).
		Msg("HTTP request completed")
}

// Database operation logging
func (l *Logger) LogDBOperation(operation, table string, duration time.Duration, err error) {
	event := l.logger.Info().
		Str("type", "db_operation").
		Str("operation", operation).
		Str("table", table).
		Dur("duration", duration)

	if err != nil {
		event = l.logger.Error().
			Str("type", "db_operation").
			Str("operation", operation).
			Str("table", table).
			Dur("duration", duration).
			Err(err)
	}

	event.Msg("Database operation")
}

// Cache operation logging
func (l *Logger) LogCacheOperation(operation, key string, hit bool, err error) {
	event := l.logger.Debug().
		Str("type", "cache_operation").
		Str("operation", operation).
		Str("key", key).
		Bool("hit", hit)

	if err != nil {
		event = l.logger.Warn().
			Str("type", "cache_operation").
			Str("operation", operation).
			Str("key", key).
			Bool("hit", hit).
			Err(err)
	}

	event.Msg("Cache operation")
}

// Authentication logging
func (l *Logger) LogAuth(userID, email, action string, success bool, reason string) {
	event := l.logger.Info().
		Str("type", "auth").
		Str("user_id", userID).
		Str("email", email).
		Str("action", action).
		Bool("success", success)

	if !success && reason != "" {
		event = event.Str("reason", reason)
	}

	if !success {
		event = l.logger.Warn().
			Str("type", "auth").
			Str("user_id", userID).
			Str("email", email).
			Str("action", action).
			Bool("success", success).
			Str("reason", reason)
	}

	event.Msg("Authentication event")
}

// Security event logging
func (l *Logger) LogSecurityEvent(eventType, description string, severity string, metadata map[string]interface{}) {
	event := l.logger.Warn().
		Str("type", "security").
		Str("event_type", eventType).
		Str("description", description).
		Str("severity", severity)

	for key, value := range metadata {
		event = event.Interface(key, value)
	}

	event.Msg("Security event")
}

// Business logic logging
func (l *Logger) LogBusinessEvent(eventType, description string, userID string, metadata map[string]interface{}) {
	event := l.logger.Info().
		Str("type", "business").
		Str("event_type", eventType).
		Str("description", description)

	if userID != "" {
		event = event.Str("user_id", userID)
	}

	for key, value := range metadata {
		event = event.Interface(key, value)
	}

	event.Msg("Business event")
}

// Global logger instance
var globalLogger *Logger

// InitGlobalLogger initializes the global logger
func InitGlobalLogger(serviceName string) {
	globalLogger = NewFromEnv(serviceName)
}

// Global convenience methods
func Info() *zerolog.Event {
	if globalLogger == nil {
		return log.Info()
	}
	return globalLogger.Info()
}

func Debug() *zerolog.Event {
	if globalLogger == nil {
		return log.Debug()
	}
	return globalLogger.Debug()
}

func Warn() *zerolog.Event {
	if globalLogger == nil {
		return log.Warn()
	}
	return globalLogger.Warn()
}

func Error() *zerolog.Event {
	if globalLogger == nil {
		return log.Error()
	}
	return globalLogger.Error()
}

func Fatal() *zerolog.Event {
	if globalLogger == nil {
		return log.Fatal()
	}
	return globalLogger.Fatal()
}

func GetGlobalLogger() *Logger {
	return globalLogger
}

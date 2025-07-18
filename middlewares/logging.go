package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-common/logger"
)

// LoggingConfig holds logging middleware configuration
type LoggingConfig struct {
	Logger       *logger.Logger
	SkipPaths    []string
	LogRequests  bool
	LogResponses bool
}

// LoggingMiddleware creates a logging middleware for Echo
func LoggingMiddleware(config LoggingConfig) echo.MiddlewareFunc {
	if config.Logger == nil {
		config.Logger = logger.GetGlobalLogger()
	}

	return echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Skip logging for certain paths
			path := c.Request().URL.Path
			for _, skipPath := range config.SkipPaths {
				if path == skipPath {
					return next(c)
				}
			}

			start := time.Now()

			// Execute the handler
			err := next(c)

			// Calculate duration
			duration := time.Since(start)

			// Get response status
			status := c.Response().Status

			// Log the request
			if config.LogRequests {
				config.Logger.LogHTTPRequest(
					c.Request().Method,
					path,
					c.RealIP(),
					status,
					duration,
				)
			}

			// Log errors if any
			if err != nil {
				config.Logger.Error().
					Err(err).
					Str("method", c.Request().Method).
					Str("path", path).
					Str("remote_addr", c.RealIP()).
					Int("status", status).
					Dur("duration", duration).
					Msg("Request error")
			}

			return err
		}
	})
}

// SimpleLoggingMiddleware creates a simple logging middleware with default config
func SimpleLoggingMiddleware() echo.MiddlewareFunc {
	return LoggingMiddleware(LoggingConfig{
		LogRequests:  true,
		LogResponses: false,
		SkipPaths:    []string{"/health", "/live", "/ready"},
	})
}

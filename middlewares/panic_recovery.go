package middlewares

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/labstack/echo/v4"
)

// PanicRecoveryMiddleware recovers from panics and returns a 500 error
func PanicRecoveryMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					// Get stack trace
					buf := make([]byte, 4096)
					n := runtime.Stack(buf, false)
					stackTrace := string(buf[:n])

					// Log the panic (when structured logging is implemented)
					// TODO: Replace with structured logging
					fmt.Printf("PANIC RECOVERED: %v\nStack trace:\n%s\n", r, stackTrace)

					// Return 500 error
					if !c.Response().Committed {
						c.JSON(http.StatusInternalServerError, map[string]string{
							"error": "Internal server error",
						})
					}
				}
			}()

			return next(c)
		}
	}
}

// PanicRecoveryWithCustomHandler allows custom panic handling
func PanicRecoveryWithCustomHandler(handler func(c echo.Context, err interface{}, stack []byte) error) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					// Get stack trace
					buf := make([]byte, 4096)
					n := runtime.Stack(buf, false)
					stackTrace := buf[:n]

					// Call custom handler
					if err := handler(c, r, stackTrace); err != nil {
						// If custom handler fails, fallback to default response
						if !c.Response().Committed {
							c.JSON(http.StatusInternalServerError, map[string]string{
								"error": "Internal server error",
							})
						}
					}
				}
			}()

			return next(c)
		}
	}
}

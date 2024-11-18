package utils

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RouteHealth(e *echo.Echo) {
	// Ruta de salud para Consul
	e.GET("/health", func(c echo.Context) error {
		return SendResponse(c, http.StatusOK, true, "Health", nil)
	})
}

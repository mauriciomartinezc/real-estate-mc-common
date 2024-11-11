package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-common/handlers"
	"github.com/mauriciomartinezc/real-estate-mc-common/repositories"
	"github.com/mauriciomartinezc/real-estate-mc-common/services"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	g := e.Group("api")
	country(g, db)
	state(g, db)
	city(g, db)
}

func country(g *echo.Group, db *gorm.DB) {
	repo := repositories.NewCountryRepository(db)
	service := services.NewCountryService(repo)
	handler := handlers.NewCountryHandler(service)

	g.GET("/countries", handler.Countries)
}

func state(g *echo.Group, db *gorm.DB) {
	repo := repositories.NewStateRepository(db)
	service := services.NewStateService(repo)
	handler := handlers.NewStateHandler(service)

	g.GET("/states/:countryUuid", handler.GetCountryStates)
}

func city(g *echo.Group, db *gorm.DB) {
	repo := repositories.NewCityRepository(db)
	service := services.NewCityService(repo)
	handler := handlers.NewCityHandler(service)

	g.GET("/cities/:stateUuid", handler.GetStateCities)
}

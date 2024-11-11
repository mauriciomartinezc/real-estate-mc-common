package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-common/cache"
	"github.com/mauriciomartinezc/real-estate-mc-common/handlers"
	"github.com/mauriciomartinezc/real-estate-mc-common/repositories"
	"github.com/mauriciomartinezc/real-estate-mc-common/services"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB, cache cache.Cache) {
	g := e.Group("api")
	country(g, db, cache)
	state(g, db, cache)
	city(g, db, cache)
}

func country(g *echo.Group, db *gorm.DB, cache cache.Cache) {
	repo := repositories.NewCountryRepository(db, cache)
	service := services.NewCountryService(repo)
	handler := handlers.NewCountryHandler(service)

	g.GET("/countries", handler.Countries)
}

func state(g *echo.Group, db *gorm.DB, cache cache.Cache) {
	repo := repositories.NewStateRepository(db, cache)
	service := services.NewStateService(repo)
	handler := handlers.NewStateHandler(service)

	g.GET("/states/:countryUuid", handler.GetCountryStates)
}

func city(g *echo.Group, db *gorm.DB, cache cache.Cache) {
	repo := repositories.NewCityRepository(db, cache)
	service := services.NewCityService(repo)
	handler := handlers.NewCityHandler(service)

	g.GET("/cities/:stateUuid", handler.GetStateCities)
}

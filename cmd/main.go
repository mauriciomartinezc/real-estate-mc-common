package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-common/config"
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"github.com/mauriciomartinezc/real-estate-mc-common/handler"
	"github.com/mauriciomartinezc/real-estate-mc-common/middleware"
	"github.com/mauriciomartinezc/real-estate-mc-common/repository"
	"github.com/mauriciomartinezc/real-estate-mc-common/seeds/cities"
	"github.com/mauriciomartinezc/real-estate-mc-common/seeds/countries"
	"github.com/mauriciomartinezc/real-estate-mc-common/seeds/currencies"
	"github.com/mauriciomartinezc/real-estate-mc-common/seeds/states"
	"github.com/mauriciomartinezc/real-estate-mc-common/service"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("application failed: %v", err)
	}
}

func run() error {
	if err := config.LoadEnv(); err != nil {
		return fmt.Errorf("failed to load environment: %w", err)
	}

	if err := config.ValidateEnvironments(); err != nil {
		return fmt.Errorf("invalid environment configuration: %w", err)
	}

	dsn, err := config.GetDSN()
	if err != nil {
		return fmt.Errorf("failed to get DSN: %w", err)
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Error,
			Colorful: true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.AutoMigrate(&domain.Currency{}, &domain.Country{}, &domain.State{}, &domain.City{}, &domain.Neighborhood{}); err != nil {
		return fmt.Errorf("failed to auto migrate models: %w", err)
	}

	// Seeds
	currencies.CreateCurrencySeeds(db)
	countries.CreateCountrySeeds(db)
	states.CreateStateSeeds(db)
	cities.CreateCitySeeds(db)

	// Repositories and Services
	countryRepo := repository.NewCountryRepository(db)
	stateRepo := repository.NewStateRepository(db)
	cityRepo := repository.NewCityRepository(db)

	countryService := service.NewCountryService(countryRepo)
	stateService := service.NewStateService(stateRepo)
	cityService := service.NewCityService(cityRepo)

	e := echo.New()
	e.Use(middleware.LanguageHandler())

	api := e.Group("/api")
	handler.NewCityHandler(api, cityService)
	handler.NewCountryHandler(api, countryService)
	handler.NewStateHandler(api, stateService)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return e.Start(":8080")
}

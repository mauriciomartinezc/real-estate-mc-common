package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-common/config"
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"github.com/mauriciomartinezc/real-estate-mc-common/middlewares"
	"github.com/mauriciomartinezc/real-estate-mc-common/routes"
	"github.com/mauriciomartinezc/real-estate-mc-common/seeds/cities"
	"github.com/mauriciomartinezc/real-estate-mc-common/seeds/countries"
	"github.com/mauriciomartinezc/real-estate-mc-common/seeds/currencies"
	"github.com/mauriciomartinezc/real-estate-mc-common/seeds/states"
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

	e := echo.New()
	e.Use(middlewares.LanguageHandler())
	routes.SetupRoutes(e, db)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return e.Start(":" + os.Getenv("SERVER_PORT"))
}

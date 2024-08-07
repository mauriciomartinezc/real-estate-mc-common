package main

import (
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
	err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	err = config.ValidateEnvironments()
	if err != nil {
		log.Fatal(err)
	}

	dsn := config.GetDSN()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Error,
			Colorful: true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&domain.Currency{}, &domain.Country{}, &domain.State{}, &domain.City{}, &domain.Neighborhood{})
	if err != nil {
		log.Fatalf("failed to auto migrate models: %v", err)
	}

	currencies.CreateCurrencySeeds(db)
	countries.CreateCountrySeeds(db)
	states.CreateStateSeeds(db)
	cities.CreateCitySeeds(db)

	countryRepo := repository.NewCountryRepository(db)
	stateRepo := repository.NewStateRepository(db)
	cityRepo := repository.NewCityRepository(db)

	countryService := service.NewCountryService(countryRepo)
	stateService := service.NewStateService(stateRepo)
	cityService := service.NewCityService(cityRepo)

	e := echo.New()

	e.Use(middleware.LanguageHandler())

	api := e.Group("/api")

	handler.NewCountryHandler(api, countryService)
	handler.NewStateHandler(api, stateService)
	handler.NewCityHandler(api, cityService)

	// swagger documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	log.Fatal(e.Start(":" + os.Getenv("SERVER_PORT")))
}

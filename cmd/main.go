package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-common/config"
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"github.com/mauriciomartinezc/real-estate-mc-common/middleware"
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
	err := config.LoadEnv()
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

	e := echo.New()

	e.Use(middleware.LanguageHandler())

	//api := e.Group("/api")

	// swagger documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	log.Fatal(e.Start(":" + os.Getenv("SERVER_PORT")))
}

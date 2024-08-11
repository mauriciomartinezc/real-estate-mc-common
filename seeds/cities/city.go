package cities

import (
	"encoding/json"
	"fmt"
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

func CreateCitySeeds(db *gorm.DB) {
	fmt.Println("CreateCitySeeds starting...")
	var count int64
	db.Model(&domain.City{}).Count(&count)
	if count == 0 {
		var cities domain.Cities
		data, err := getAllCitiesJson()
		if err != nil {
			log.Printf("failed to get cities JSON: %v", err)
			return
		}

		if err := json.Unmarshal(data, &cities); err != nil {
			log.Printf("failed to unmarshal JSON: %v", err)
			return
		}

		for _, city := range cities {
			if err := db.Create(&city).Error; err != nil {
				log.Printf("failed to create city %s: %v", city.Name, err)
			}
		}
	}
	fmt.Println("CreateCitySeeds completed successfully.")
}

func getAllCitiesJson() ([]byte, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error getting current working directory: %w", err)
	}
	pathFile := filepath.Join(cwd, "../seeds/cities/data.json")
	return os.ReadFile(pathFile)
}

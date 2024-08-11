package countries

import (
	"encoding/json"
	"fmt"
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

func CreateCountrySeeds(db *gorm.DB) {
	fmt.Println("CreateCountrySeeds starting...")
	var count int64
	db.Model(&domain.Country{}).Count(&count)
	if count == 0 {
		var countries domain.Countries
		data, err := getAllCountriesJson()
		if err != nil {
			log.Printf("failed to get countries JSON: %v", err)
			return
		}

		if err := json.Unmarshal(data, &countries); err != nil {
			log.Printf("failed to unmarshal JSON: %v", err)
			return
		}

		for _, country := range countries {
			if err := db.Create(&country).Error; err != nil {
				log.Printf("failed to create country %s: %v", country.Name, err)
			}
		}
	}
	fmt.Println("CreateCountrySeeds completed successfully.")
}

func getAllCountriesJson() ([]byte, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error getting current working directory: %w", err)
	}
	pathFile := filepath.Join(cwd, "../seeds/countries/data.json")
	return os.ReadFile(pathFile)
}

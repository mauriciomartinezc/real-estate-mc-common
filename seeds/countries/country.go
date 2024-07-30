package countries

import (
	"encoding/json"
	"fmt"
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"gorm.io/gorm"
	"io/ioutil"
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
		data := getAllCountriesJson()

		if err := json.Unmarshal(data, &countries); err != nil {
			log.Fatalf("failed to unmarshal JSON: %v", err)
		}

		for _, country := range countries {
			if err := db.Create(&country).Error; err != nil {
				log.Printf("failed to create countries %s: %v", country.Name, err)
			}
		}
	}
	fmt.Println("CreateCountrySeeds completed success.")
}

func getAllCountriesJson() []byte {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	pathFile := filepath.Join(cwd, "../seeds/countries/data.json")
	file, err := os.Open(pathFile)
	if err != nil {
		log.Fatalf("failed to open JSON file: %v", err)
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read JSON file: %v", err)
	}

	return byteValue
}

package cities

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

func CreateCitySeeds(db *gorm.DB) {
	fmt.Println("CreateCitySeeds starting...")
	var count int64
	db.Model(&domain.City{}).Count(&count)
	if count == 0 {
		var cities domain.Cities
		data := getAllCitiesJson()

		if err := json.Unmarshal(data, &cities); err != nil {
			log.Fatalf("failed to unmarshal JSON: %v", err)
		}

		for _, city := range cities {
			if err := db.Create(&city).Error; err != nil {
				log.Printf("failed to create city %s: %v", city.Name, err)
			}
		}
	}
	fmt.Println("CreateCitySeeds completed success.")
}

func getAllCitiesJson() []byte {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	pathFile := filepath.Join(cwd, "../seeds/cities/data.json")
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
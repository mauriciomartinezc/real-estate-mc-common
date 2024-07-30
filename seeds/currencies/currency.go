package currencies

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

func CreateCurrencySeeds(db *gorm.DB) {
	fmt.Println("CreateCurrencySeeds starting...")
	var count int64
	db.Model(&domain.Currency{}).Count(&count)
	if count == 0 {
		var currencies domain.Currencies
		data := getAllCurrenciesJson()

		if err := json.Unmarshal(data, &currencies); err != nil {
			log.Fatalf("failed to unmarshal JSON: %v", err)
		}

		for _, currency := range currencies {
			if err := db.Create(&currency).Error; err != nil {
				log.Printf("failed to create countries %s: %v", currency.Name, err)
			}
		}
	}
	fmt.Println("CreateCurrencySeeds completed success.")
}

func getAllCurrenciesJson() []byte {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	pathFile := filepath.Join(cwd, "../seeds/currencies/data.json")
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

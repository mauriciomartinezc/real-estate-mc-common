package currencies

import (
	"encoding/json"
	"fmt"
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"gorm.io/gorm"
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
		data, err := getAllCurrenciesJson()
		if err != nil {
			log.Printf("failed to get currencies JSON: %v", err)
			return
		}

		if err := json.Unmarshal(data, &currencies); err != nil {
			log.Printf("failed to unmarshal JSON: %v", err)
			return
		}

		for _, currency := range currencies {
			if err := db.Create(&currency).Error; err != nil {
				log.Printf("failed to create currency %s: %v", currency.Name, err)
			}
		}
	}
	fmt.Println("CreateCurrencySeeds completed successfully.")
}

func getAllCurrenciesJson() ([]byte, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error getting current working directory: %w", err)
	}
	pathFile := filepath.Join(cwd, "../seeds/currencies/data.json")
	return os.ReadFile(pathFile)
}

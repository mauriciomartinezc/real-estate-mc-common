package states

import (
	"encoding/json"
	"fmt"
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

func CreateStateSeeds(db *gorm.DB) {
	fmt.Println("CreateStateSeeds starting...")
	var count int64
	db.Model(&domain.State{}).Count(&count)
	if count == 0 {
		var states domain.States
		data, err := getAllStatesJson()
		if err != nil {
			log.Printf("failed to get states JSON: %v", err)
			return
		}

		if err := json.Unmarshal(data, &states); err != nil {
			log.Printf("failed to unmarshal JSON: %v", err)
			return
		}

		for _, state := range states {
			if err := db.Create(&state).Error; err != nil {
				log.Printf("failed to create state %s: %v", state.Name, err)
			}
		}
	}
	fmt.Println("CreateStateSeeds completed successfully.")
}

func getAllStatesJson() ([]byte, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error getting current working directory: %w", err)
	}
	pathFile := filepath.Join(cwd, "../seeds/states/data.json")
	return os.ReadFile(pathFile)
}

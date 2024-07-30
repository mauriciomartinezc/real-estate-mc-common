package states

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

func CreateStateSeeds(db *gorm.DB) {
	fmt.Println("CreateStateSeeds starting...")
	var count int64
	db.Model(&domain.State{}).Count(&count)
	if count == 0 {
		var states domain.States
		data := getAllEstatesJson()

		if err := json.Unmarshal(data, &states); err != nil {
			log.Fatalf("failed to unmarshal JSON: %v", err)
		}

		for _, estate := range states {
			if err := db.Create(&estate).Error; err != nil {
				log.Printf("failed to create estate %s: %v", estate.Name, err)
			}
		}
	}
	fmt.Println("CreateStateSeeds completed success.")
}

func getAllEstatesJson() []byte {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	pathFile := filepath.Join(cwd, "../seeds/states/data.json")
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

package repository

import (
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
	"testing"
)

func SetupTestDB() (*gorm.DB, func()) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	_ = db.AutoMigrate(&domain.City{}, &domain.Neighborhood{})

	cleanup := func() {
		sqlDB, err := db.DB()
		if err == nil {
			_ = sqlDB.Close()
		}
	}

	return db, cleanup
}

func TestCountryRepository_GetAll(t *testing.T) {
	db, cleanup := SetupTestDB()
	defer cleanup()

	repo := NewCountryRepository(db)

	// Seed some data
	db.Create(&domain.Country{Name: "Test Country", Active: true})

	t.Run("Get All Active Countries", func(t *testing.T) {
		countries, err := repo.GetAll()
		assert.NoError(t, err)
		assert.NotEmpty(t, countries)
		assert.Equal(t, "Test Country", (*countries)[0].Name)
	})

	t.Run("No Active Countries", func(t *testing.T) {
		db.Delete(&domain.Country{}, "name = ?", "Test Country")
		countries, err := repo.GetAll()
		assert.NoError(t, err)
		assert.Empty(t, countries)
	})
}

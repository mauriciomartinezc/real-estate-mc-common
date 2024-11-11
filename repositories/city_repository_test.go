package repositories

import (
	"github.com/google/uuid"
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCityRepository_GetStateCities(t *testing.T) {
	db, cleanup := SetupTestDB()
	defer cleanup()

	repo := NewCityRepository(db)

	// Seed some data
	stateUUID := "123e4567-e89b-12d3-a456-426614174000"
	db.Create(&domain.City{Name: "Test City", StateId: uuid.MustParse(stateUUID)})

	t.Run("Valid UUID", func(t *testing.T) {
		cities, err := repo.GetStateCities(stateUUID)
		assert.NoError(t, err)
		assert.NotEmpty(t, cities)
		assert.Equal(t, "Test City", (*cities)[0].Name)
	})

	t.Run("Invalid UUID", func(t *testing.T) {
		invalidUUID := "invalid-uuid"
		cities, err := repo.GetStateCities(invalidUUID)
		assert.Error(t, err)
		assert.Nil(t, cities)
	})
}

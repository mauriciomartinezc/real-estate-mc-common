package repository

import (
	"github.com/google/uuid"
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStateRepository_GetCountryStates(t *testing.T) {
	db, cleanup := SetupTestDB()
	defer cleanup()

	repo := NewStateRepository(db)

	// Seed some data
	countryUUID := "123e4567-e89b-12d3-a456-426614174000"
	db.Create(&domain.State{Name: "Test State", CountryId: uuid.MustParse(countryUUID)})

	t.Run("Valid UUID", func(t *testing.T) {
		states, err := repo.GetCountryStates(countryUUID)
		assert.NoError(t, err)
		assert.NotEmpty(t, states)
		assert.Equal(t, "Test State", (*states)[0].Name)
	})

	t.Run("Invalid UUID", func(t *testing.T) {
		invalidUUID := "invalid-uuid"
		cities, err := repo.GetCountryStates(invalidUUID)
		assert.Error(t, err)
		assert.Nil(t, cities)
	})
}

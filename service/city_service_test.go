package service

import (
	"errors"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"testing"

	"github.com/google/uuid"
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCityRepository struct {
	mock.Mock
}

func (m *MockCityRepository) GetStateCities(stateUuid string) (*domain.Cities, error) {
	args := m.Called(stateUuid)
	return args.Get(0).(*domain.Cities), args.Error(1)
}

func TestGetStateCities(t *testing.T) {
	mockRepo := new(MockCityRepository)
	cityService := NewCityService(mockRepo)

	t.Run("Valid UUID", func(t *testing.T) {
		stateUuid := uuid.New().String()
		expectedCities := &domain.Cities{}

		mockRepo.On("GetStateCities", stateUuid).Return(expectedCities, nil)

		cities, err := cityService.GetStateCities(stateUuid)

		assert.NoError(t, err)
		assert.Equal(t, expectedCities, cities)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid UUID", func(t *testing.T) {
		invalidUuid := "invalid-uuid"

		cities, err := cityService.GetStateCities(invalidUuid)

		assert.Error(t, err)
		assert.Nil(t, cities)
		assert.EqualError(t, err, locales.InvalidUuid)
	})

	t.Run("Repository Error", func(t *testing.T) {
		stateUuid := uuid.New().String()
		expectedCities := &domain.Cities{}
		expectedError := errors.New("repository error")

		mockRepo.On("GetStateCities", stateUuid).Return(expectedCities, expectedError)

		cities, err := cityService.GetStateCities(stateUuid)

		assert.Error(t, err)
		assert.Nil(t, cities)
		assert.EqualError(t, err, expectedError.Error())

		mockRepo.AssertExpectations(t)
	})
}

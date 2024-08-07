package service

import (
	"github.com/google/uuid"
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
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

	stateUuid := uuid.New().String()
	expectedCities := &domain.Cities{}

	mockRepo.On("GetStateCities", stateUuid).Return(expectedCities, nil)

	cities, err := cityService.GetStateCities(stateUuid)

	assert.NoError(t, err)
	assert.Equal(t, expectedCities, cities)

	mockRepo.AssertExpectations(t)
}

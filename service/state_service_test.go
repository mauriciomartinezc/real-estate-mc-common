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

type MockStateRepository struct {
	mock.Mock
}

func (m *MockStateRepository) GetCountryStates(countryUuid string) (*domain.States, error) {
	args := m.Called(countryUuid)
	return args.Get(0).(*domain.States), args.Error(1)
}

func TestGetCountryStates(t *testing.T) {
	mockRepo := new(MockStateRepository)
	stateService := NewStateService(mockRepo)

	t.Run("Valid UUID", func(t *testing.T) {
		countryUuid := uuid.New().String()
		expectedStates := &domain.States{}

		mockRepo.On("GetCountryStates", countryUuid).Return(expectedStates, nil)

		states, err := stateService.GetCountryStates(countryUuid)

		assert.NoError(t, err)
		assert.Equal(t, expectedStates, states)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid UUID", func(t *testing.T) {
		invalidUuid := "invalid-uuid"

		states, err := stateService.GetCountryStates(invalidUuid)

		assert.Error(t, err)
		assert.Nil(t, states)
		assert.EqualError(t, err, locales.InvalidUuid)
	})

	t.Run("Repository Error", func(t *testing.T) {
		countryUuid := uuid.New().String()
		expectedStates := &domain.States{}
		expectedError := errors.New("repository error")

		mockRepo.On("GetCountryStates", countryUuid).Return(expectedStates, expectedError)

		states, err := stateService.GetCountryStates(countryUuid)

		assert.Error(t, err)
		assert.Nil(t, states)
		assert.EqualError(t, err, expectedError.Error())

		mockRepo.AssertExpectations(t)
	})
}

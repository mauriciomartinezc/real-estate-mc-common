package service

import (
	"github.com/google/uuid"
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
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

	stateUuid := uuid.New().String()
	expectedStates := &domain.States{}

	mockRepo.On("GetCountryStates", stateUuid).Return(expectedStates, nil)

	states, err := stateService.GetCountryStates(stateUuid)

	assert.NoError(t, err)
	assert.Equal(t, expectedStates, states)

	mockRepo.AssertExpectations(t)
}

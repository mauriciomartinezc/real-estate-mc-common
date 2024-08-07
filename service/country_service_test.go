package service

import (
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockCountryRepository struct {
	mock.Mock
}

func (m *MockCountryRepository) GetAll() (*domain.Countries, error) {
	args := m.Called()
	return args.Get(0).(*domain.Countries), args.Error(1)
}

func TestGetCountries(t *testing.T) {
	mockRepo := new(MockCountryRepository)
	countryService := NewCountryService(mockRepo)

	expectedCountries := &domain.Countries{}

	mockRepo.On("GetAll").Return(expectedCountries, nil)

	countries, err := countryService.Countries()

	assert.NoError(t, err)
	assert.Equal(t, expectedCountries, countries)

	mockRepo.AssertExpectations(t)
}

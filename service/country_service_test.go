package service

import (
	"errors"
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// Mock del repositorio
type MockCountryRepository struct {
	mock.Mock
}

func (m *MockCountryRepository) GetAll() (*domain.Countries, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Countries), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestCountryService_GetCountries_Success(t *testing.T) {
	mockRepo := new(MockCountryRepository)
	countryService := NewCountryService(mockRepo)

	expectedCountries := &domain.Countries{}

	mockRepo.On("GetAll").Return(expectedCountries, nil)

	countries, err := countryService.Countries()

	assert.NoError(t, err)
	assert.Equal(t, expectedCountries, countries)

	mockRepo.AssertExpectations(t)
}

func TestCountryService_GetCountries_Error(t *testing.T) {
	mockRepo := new(MockCountryRepository)
	countryService := NewCountryService(mockRepo)

	var nilCountries *domain.Countries
	mockRepo.On("GetAll").Return(nilCountries, errors.New("database error"))

	countries, err := countryService.Countries()

	assert.Error(t, err)
	assert.Nil(t, countries)

	mockRepo.AssertExpectations(t)
}

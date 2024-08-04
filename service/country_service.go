package service

import (
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"github.com/mauriciomartinezc/real-estate-mc-common/repository"
)

type CountryService interface {
	Countries() (*domain.Countries, error)
}

type countryService struct {
	countryRepository repository.CountryRepository
}

func NewCountryService(countryRepository repository.CountryRepository) CountryService {
	return &countryService{
		countryRepository: countryRepository,
	}
}

func (s *countryService) Countries() (countries *domain.Countries, err error) {
	countries, err = s.countryRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return countries, nil
}

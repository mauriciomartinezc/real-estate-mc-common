package services

import (
	"errors"
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/repositories"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
)

type CityService interface {
	GetStateCities(stateUuid string) (*domain.Cities, error)
}

type cityService struct {
	cityRepository repositories.CityRepository
}

func NewCityService(cityRepository repositories.CityRepository) CityService {
	return &cityService{cityRepository: cityRepository}
}

func (s *cityService) GetStateCities(stateUuid string) (*domain.Cities, error) {
	if !utils.IsValidUUID(stateUuid) {
		return nil, errors.New(locales.InvalidUuid)
	}
	cities, err := s.cityRepository.GetStateCities(stateUuid)
	if err != nil {
		return nil, err
	}
	return cities, nil
}

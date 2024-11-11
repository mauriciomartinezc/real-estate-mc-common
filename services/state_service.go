package services

import (
	"errors"
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/repositories"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
)

type StateService interface {
	GetCountryStates(countryUuid string) (*domain.States, error)
}

type stateService struct {
	stateRepository repositories.StateRepository
}

func NewStateService(stateRepository repositories.StateRepository) StateService {
	return &stateService{stateRepository: stateRepository}
}

func (s *stateService) GetCountryStates(countryUuid string) (*domain.States, error) {
	if !utils.IsValidUUID(countryUuid) {
		return nil, errors.New(locales.InvalidUuid)
	}
	states, err := s.stateRepository.GetCountryStates(countryUuid)
	if err != nil {
		return nil, err
	}
	return states, nil
}

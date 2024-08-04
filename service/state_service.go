package service

import (
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"github.com/mauriciomartinezc/real-estate-mc-common/repository"
)

type StateService interface {
	GetCountryStates(countryUuid string) (*domain.States, error)
}

type stateService struct {
	stateRepository repository.StateRepository
}

func NewStateService(stateRepository repository.StateRepository) StateService {
	return &stateService{stateRepository: stateRepository}
}

func (s *stateService) GetCountryStates(countryUuid string) (*domain.States, error) {
	states, err := s.stateRepository.GetCountryStates(countryUuid)
	if err != nil {
		return nil, err
	}
	return states, nil
}

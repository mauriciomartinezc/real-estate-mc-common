package repository

import (
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"gorm.io/gorm"
)

type StateRepository interface {
	GetCountryStates(countryUuid string) (*domain.States, error)
}

type stateRepository struct {
	db *gorm.DB
}

func NewStateRepository(db *gorm.DB) StateRepository {
	return &stateRepository{db: db}
}

func (r *stateRepository) GetCountryStates(countryUuid string) (*domain.States, error) {
	states := &domain.States{}
	err := r.db.Where("country_id = ?", countryUuid).Order("name").Find(states).Error
	if err != nil {
		return nil, err
	}
	return states, nil
}

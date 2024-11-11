package repositories

import (
	"errors"
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
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
	if !utils.IsValidUUID(countryUuid) {
		return nil, errors.New(locales.InvalidUuid)
	}
	states := &domain.States{}
	err := r.db.Where("country_id = ?", countryUuid).Order("name").Find(states).Error
	if err != nil {
		return nil, err
	}
	return states, nil
}

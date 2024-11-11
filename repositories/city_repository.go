package repositories

import (
	"errors"
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
	"gorm.io/gorm"
)

type CityRepository interface {
	GetStateCities(stateUuid string) (*domain.Cities, error)
}

type cityRepository struct {
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) CityRepository {
	return &cityRepository{db: db}
}

func (r *cityRepository) GetStateCities(stateUuid string) (*domain.Cities, error) {
	if !utils.IsValidUUID(stateUuid) {
		return nil, errors.New(locales.InvalidUuid)
	}
	cities := &domain.Cities{}
	err := r.db.Preload("Neighborhoods").Where("state_id = ?", stateUuid).Order("name").Find(cities).Error
	if err != nil {
		return nil, err
	}
	return cities, nil
}

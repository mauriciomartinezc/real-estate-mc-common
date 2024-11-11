package repositories

import (
	"errors"
	"fmt"
	"github.com/mauriciomartinezc/real-estate-mc-common/cache"
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
	"gorm.io/gorm"
	"time"
)

type StateRepository interface {
	GetCountryStates(countryUuid string) (*domain.States, error)
}

type stateRepository struct {
	db    *gorm.DB
	cache cache.Cache
}

func NewStateRepository(db *gorm.DB, cache cache.Cache) StateRepository {
	return &stateRepository{db: db, cache: cache}
}

func (r *stateRepository) GetCountryStates(countryUuid string) (*domain.States, error) {
	if !utils.IsValidUUID(countryUuid) {
		return nil, errors.New(locales.InvalidUuid)
	}
	cacheKey := fmt.Sprintf("states_country:%s", countryUuid)
	var states domain.States

	if err := r.cache.Get(cacheKey, &states); err != nil {
		if len(states) > 0 {
			return &states, nil
		}
	}

	err := r.db.Where("country_id = ?", countryUuid).Order("name").Find(&states).Error

	if err != nil {
		return nil, err
	}

	if len(states) > 0 {
		if err := r.cache.Set(cacheKey, states, 5*time.Minute); err != nil {
			return nil, fmt.Errorf("error saving to cache: %v", err)
		}
	}

	return &states, nil
}

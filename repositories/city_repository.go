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

type CityRepository interface {
	GetStateCities(stateUuid string) (*domain.Cities, error)
}

type cityRepository struct {
	db    *gorm.DB
	cache cache.Cache
}

func NewCityRepository(db *gorm.DB, cache cache.Cache) CityRepository {
	return &cityRepository{db: db, cache: cache}
}

func (r *cityRepository) GetStateCities(stateUuid string) (*domain.Cities, error) {
	if !utils.IsValidUUID(stateUuid) {
		return nil, errors.New(locales.InvalidUuid)
	}

	cacheKey := fmt.Sprintf("cities_state:%s", stateUuid)
	var cities domain.Cities

	if err := r.cache.Get(cacheKey, &cities); err != nil {
		if len(cities) > 0 {
			return &cities, nil
		}
	}

	err := r.db.Preload("Neighborhoods").Where("state_id = ?", stateUuid).Order("name").Find(&cities).Error
	if err != nil {
		return nil, err
	}

	if len(cities) > 0 {
		if err := r.cache.Set(cacheKey, cities, 5*time.Minute); err != nil {
			return nil, fmt.Errorf("error saving to cache: %v", err)
		}
	}

	return &cities, nil
}

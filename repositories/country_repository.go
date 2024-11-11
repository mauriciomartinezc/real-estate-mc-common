package repositories

import (
	"fmt"
	"github.com/mauriciomartinezc/real-estate-mc-common/cache"
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"gorm.io/gorm"
	"time"
)

type CountryRepository interface {
	GetAll() (*domain.Countries, error)
}

type countryRepository struct {
	db    *gorm.DB
	cache cache.Cache
}

func NewCountryRepository(db *gorm.DB, cache cache.Cache) CountryRepository {
	return &countryRepository{db: db, cache: cache}
}

func (r *countryRepository) GetAll() (*domain.Countries, error) {
	cacheKey := "countries:all"
	var countries domain.Countries

	if err := r.cache.Get(cacheKey, &countries); err != nil {
		if len(countries) > 0 {
			return &countries, nil
		}
	}

	result := r.db.Preload("Currency").Where("active = ?", true).Find(&countries)

	if result.Error != nil {
		return nil, result.Error
	}

	if len(countries) > 0 {
		if err := r.cache.Set(cacheKey, countries, 5*time.Minute); err != nil {
			return nil, fmt.Errorf("error saving to cache: %v", err)
		}
	}

	return &countries, nil
}

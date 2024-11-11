package repositories

import (
	"github.com/mauriciomartinezc/real-estate-mc-common/domain"
	"gorm.io/gorm"
)

type CountryRepository interface {
	GetAll() (*domain.Countries, error)
}

type countryRepository struct {
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) CountryRepository {
	return &countryRepository{db: db}
}

func (r *countryRepository) GetAll() (*domain.Countries, error) {
	var countries domain.Countries
	result := r.db.Preload("Currency").Where("active = ?", true).Find(&countries)
	if result.Error != nil {
		return nil, result.Error
	}
	return &countries, nil
}

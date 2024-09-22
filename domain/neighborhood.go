package domain

import (
	"github.com/google/uuid"
)

type Neighborhood struct {
	ID     uuid.UUID  `json:"id,omitempty" gorm:"type:uuid;primaryKey"`
	Name   string     `json:"name,omitempty" gorm:"not null"`
	CityId *uuid.UUID `json:"city_id,omitempty" gorm:"type:uuid;not null"`
	City   *City      `json:"service_test,omitempty" gorm:"foreignKey:CityId"`
}

type Neighborhoods []Neighborhood

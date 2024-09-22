package domain

import (
	"github.com/google/uuid"
)

type State struct {
	ID        uuid.UUID  `json:"id,omitempty" gorm:"type:uuid;primaryKey"`
	Name      string     `json:"name,omitempty" gorm:"not null"`
	Code      string     `json:"state_code,omitempty" gorm:"not null"`
	Latitude  string     `json:"latitude,omitempty" gorm:"not null"`
	Longitude string     `json:"longitude,omitempty" gorm:"not null"`
	CountryId *uuid.UUID `json:"country_id,omitempty" gorm:"type:uuid;not null"`
	Country   *Country   `json:"country,omitempty" gorm:"foreignKey:CountryId"`
	Cities    *Cities    `json:"cities,omitempty" gorm:"references:ID"`
}

type StateJson struct {
	*States
}

type States []State
type StatesJson []StateJson

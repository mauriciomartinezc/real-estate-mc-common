package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type State struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Code      string    `json:"state_code" gorm:"not null"`
	Latitude  string    `json:"latitude" gorm:"not null"`
	Longitude string    `json:"longitude" gorm:"not null"`
	CountryId uuid.UUID `json:"country_id" gorm:"type:uuid;not null"`
	Country   Country   `gorm:"foreignKey:CountryId"`
	Cities    Cities    `gorm:"references:ID"`
}

type States []State

func (e *State) BeforeCreate(ctx *gorm.DB) (err error) {
	e.ID = uuid.New()
	return
}

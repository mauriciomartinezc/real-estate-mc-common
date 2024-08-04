package domain

import (
	"github.com/google/uuid"
)

type City struct {
	ID            uuid.UUID     `json:"id,omitempty" gorm:"type:uuid;primaryKey"`
	Name          string        `json:"name,omitempty" gorm:"not null"`
	StateId       uuid.UUID     `json:"state_id,omitempty" gorm:"type:uuid;not null"`
	State         State         `json:"state,omitempty" gorm:"foreignKey:StateId"`
	Neighborhoods Neighborhoods `json:"neighborhoods" gorm:"references:ID"`
}

type Cities []City

package domain

import (
	"github.com/google/uuid"
)

type Currency struct {
	ID        uuid.UUID  `json:"id,omitempty" gorm:"type:uuid;primaryKey"`
	Name      string     `json:"name,omitempty" gorm:"not null"`
	Code      string     `json:"code,omitempty" gorm:"not null"`
	Symbol    string     `json:"symbol,omitempty" gorm:"not null"`
	Countries *Countries `json:"countries,omitempty" gorm:"references:ID"`
}

type Currencies []Currency

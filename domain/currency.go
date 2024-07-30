package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Currency struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Code      string    `json:"code" gorm:"not null"`
	Symbol    string    `json:"symbol" gorm:"not null"`
	Countries Countries `gorm:"references:ID"`
}

type Currencies []Currency

func (c *Currency) BeforeCreate(ctx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}

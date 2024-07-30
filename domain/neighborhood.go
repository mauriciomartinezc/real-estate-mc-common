package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type Neighborhood struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name   string    `json:"name" gorm:"not null"`
	CityId uuid.UUID `gorm:"type:uuid;not null"`
	City   City      `gorm:"foreignKey:CityId"`
}

type Neighborhoods []Neighborhood

func (n *Neighborhood) BeforeCreate(ctx *gorm.DB) (err error) {
	n.ID = uuid.New()
	n.Name = strings.TrimSpace(n.Name)
	return
}

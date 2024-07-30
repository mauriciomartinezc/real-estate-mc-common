package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type City struct {
	ID            uuid.UUID     `gorm:"type:uuid;primaryKey"`
	Name          string        `json:"name" gorm:"not null"`
	StateId       uuid.UUID     `json:"state_id" gorm:"type:uuid;not null"`
	State         State         `gorm:"foreignKey:StateId"`
	Neighborhoods Neighborhoods `gorm:"references:ID"`
}

type Cities []City

func (c *City) BeforeCreate(ctx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}

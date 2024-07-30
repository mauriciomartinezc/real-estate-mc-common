package domain

import (
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Country struct {
	ID           uuid.UUID       `gorm:"type:uuid;primaryKey"`
	Name         string          `json:"name" gorm:"not null"`
	Iso3         string          `json:"iso3" gorm:"not null"`
	Iso2         string          `json:"iso2" gorm:"not null"`
	NumericCode  string          `json:"numeric_code" gorm:"not null"`
	Capital      string          `json:"capital" gorm:"not null"`
	Tld          string          `json:"tld" gorm:"not null"`
	Native       string          `json:"native" gorm:"not null"`
	Region       string          `json:"region" gorm:"not null"`
	Nationality  string          `json:"nationality" gorm:"not null"`
	Timezones    json.RawMessage `json:"timezones" gorm:"type:jsonb"`
	Translations json.RawMessage `json:"translations" gorm:"type:jsonb"`
	Latitude     string          `json:"latitude" gorm:"not null"`
	Longitude    string          `json:"longitude" gorm:"not null"`
	Emoji        string          `json:"emoji" gorm:"not null"`
	EmojiU       string          `json:"emojiU" gorm:"not null"`
	Active       bool            `json:"active" gorm:"default:false"`
	CurrencyId   uuid.UUID       `json:"currency_id" gorm:"not null"`
	Currency     Currency        `gorm:"foreignKey:CurrencyId"`
	States       States          `gorm:"references:ID"`
}

type Countries []Country

func (c *Country) BeforeCreate(ctx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}

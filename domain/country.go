package domain

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Country struct {
	ID           uuid.UUID       `json:"id,omitempty" gorm:"type:uuid;primaryKey"`
	Name         string          `json:"name,omitempty" gorm:"not null"`
	Iso3         string          `json:"iso3,omitempty" gorm:"not null"`
	Iso2         string          `json:"iso2,omitempty" gorm:"not null"`
	NumericCode  string          `json:"numeric_code,omitempty" gorm:"not null"`
	Capital      string          `json:"capital,omitempty" gorm:"not null"`
	Tld          string          `json:"tld,omitempty" gorm:"not null"`
	Native       string          `json:"native,omitempty" gorm:"not null"`
	Region       string          `json:"region,omitempty" gorm:"not null"`
	Nationality  string          `json:"nationality,omitempty" gorm:"not null"`
	Timezones    json.RawMessage `json:"timezones,omitempty" gorm:"type:jsonb"`
	Translations json.RawMessage `json:"translations,omitempty" gorm:"type:jsonb"`
	Latitude     string          `json:"latitude,omitempty" gorm:"not null"`
	Longitude    string          `json:"longitude,omitempty" gorm:"not null"`
	Emoji        string          `json:"emoji,omitempty" gorm:"not null"`
	EmojiU       string          `json:"emojiU,omitempty" gorm:"not null"`
	Active       bool            `json:"active,omitempty" gorm:"default:false"`
	CurrencyId   *uuid.UUID      `json:"currency_id,omitempty" gorm:"type:uuid;not null"`
	Currency     Currency        `json:"currency,omitempty" gorm:"foreignKey:CurrencyId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	States       States          `json:"states,omitempty" gorm:"references:ID"`
}

type Countries []Country

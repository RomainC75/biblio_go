package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Title struct {
	gorm.Model
	ID           uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4()"`
	LanguageCode LanguageCode `form:"language_code" `
	TitleName    string       `gorm:"varchar:191"`
	BookRef      uuid.UUID
}

type LanguageCode int

const (
	English = iota
	French
	Deutsch
	Japanese
)
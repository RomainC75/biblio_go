package models

import (
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	// -> ID, CreatedAt, UpdatedAt, DeletedAt
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	ISBN      string    `gorm:"varchar:191;unique"`
	GenreCode GenreCode `form:"genre_code"`
	Title     []Title   `gorm:"many2many:book_title;"`
	Authors   []Author  `gorm:"many2many:book_author;"`
}

type Title struct {
	gorm.Model
	ID           uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4()"`
	LanguageCode LanguageCode `form:"language_code" `
	TitleName    string       `gorm:"varchar:191"`
	Books        []*Book      `gorm:"many2many:book_title;"`
}

type LanguageCode int

const (
	English = iota
	French
	Deutsch
	Japanese
)

type GenreCode int

const (
	ScienceFiction = iota
	Fantastic
	Love
)

package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	// -> ID, CreatedAt, UpdatedAt, DeletedAt
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title string `gorm:"foreignKey:BookRefer"`
	ISBN      string    `gorm:"varchar:191;"`
	Authors   []Author  `gorm:"many2many:book_author;"`
	LanguageCode LanguageCode
	GenreCode GenreCode `form:"genre_code"`
	UserRefer	uuid.UUID
}

type GenreCode int

const (
	ScienceFiction = iota
	Fantastic
	Love
)

type LanguageCode int

const (
	English = iota
	French
	Deutsch
	Japanese
)
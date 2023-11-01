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
	Authors   []Author  `gorm:"many2many:book_author;"`
	// Title     []Title   `gorm:"many2many:book_title;"`
	Title []Title `gorm:"foreignKey:BookRefer"`
}

type GenreCode int

const (
	ScienceFiction = iota
	Fantastic
	Love
)

package models

import (
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	// -> ID, CreatedAt, UpdatedAt, DeletedAt
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title   string    `gorm:"varchar:191"`
	ISBN    string    `gorm:"varchar:191;unique"`
	Genre   Genre     `form:"genre" binding:"oneof=S.F fantastic love"`
	Authors []Author  `gorm:"many2many:book_author;"`
}

type Genre string

const (
	ScienceFiction Genre = "S.F"
	Fantastic      Genre = "fantastic"
	Love           Genre = "love"
)

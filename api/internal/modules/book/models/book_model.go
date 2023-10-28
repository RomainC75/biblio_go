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
	ISNB    string    `gorm:"varchar:191"`
	Authors []Author  `gorm:"many2many:book_author;"`
}

package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// -> ID, CreatedAt, UpdatedAt, DeletedAt
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Email    string    `gorm:"varchar:191"`
	Password string    `gorm:"varchar:191"`
	Firstname string	`gorm:"varchar:191"`
  	Lastname string `gorm:"varchar:191"`
}

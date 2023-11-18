package models

import (
	"github.com/google/uuid"
	bookModels "gitub.com/RomainC75/biblio/internal/modules/book/models"
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
	Books	[]*bookModels.Book	`gorm:"foreignKey:UserRefer"`
}

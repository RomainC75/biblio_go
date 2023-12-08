package repositories

import (
	"gitub.com/RomainC75/biblio/data/database"
	Models "gitub.com/RomainC75/biblio/data/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepo() *UserRepository {
	return &UserRepository{
		DB: database.Connection(),
	}
}

func (UserRepository *UserRepository) Create(user Models.User) Models.User {
	var newUser Models.User

	UserRepository.DB.Create(&user).Scan(&newUser)

	return newUser
}

func (UserRepository *UserRepository) FindByEmail(email string) Models.User {
	var foundUser Models.User

	UserRepository.DB.First(&foundUser, "email = ?", email)

	return foundUser
}

func (UserRepository *UserRepository) FindById(id int) Models.User {
	var foundUser Models.User

	UserRepository.DB.First(&foundUser, "id = ?", id)

	return foundUser
}

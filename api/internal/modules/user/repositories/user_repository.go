package repositories

import (
	userModel "gitub.com/RomainC75/biblio/internal/modules/user/models"
	"gitub.com/RomainC75/biblio/pkg/database"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func New() *UserRepository {
	return &UserRepository{
		DB: database.Connection(),
	}
}

func (UserRepository *UserRepository) Create(user userModel.User) userModel.User {
	var newUser userModel.User

	UserRepository.DB.Create(&user).Scan(&newUser)

	return newUser
}

func (UserRepository *UserRepository) FindByEmail(email string) userModel.User {
	var foundUser userModel.User

	UserRepository.DB.First(&foundUser, "email = ?", email)

	return foundUser
}

func (UserRepository *UserRepository) FindById(id int) userModel.User {
	var foundUser userModel.User

	UserRepository.DB.First(&foundUser, "id = ?", id)

	return foundUser
}

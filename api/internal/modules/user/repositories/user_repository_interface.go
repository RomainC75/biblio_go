package repositories

import (
	userModel "gitub.com/RomainC75/biblio/internal/modules/user/models"
)

type UserRepositoryInterface interface {
	Create(user userModel.User) userModel.User
	FindByEmail(email string) userModel.User
	FindById(id int) userModel.User
}

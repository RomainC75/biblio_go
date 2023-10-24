package repositories

import (
	userModel "gitub.com/RomainC75/biblio/internal/modules/book/models"
)

type BookRepositoryInterface interface {
	Create(book userModel.Book) userModel.Book
	FindByEmail(email string) userModel.Book
	FindById(id int) userModel.Book
}

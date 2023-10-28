package repositories

import (
	bookModel "gitub.com/RomainC75/biblio/internal/modules/book/models"
)

type BookRepositoryInterface interface {
	Create(book bookModel.Book) bookModel.Book
	FirstOrCreateAuthors(newAuthors []string) []bookModel.Author
	FindByEmail(email string) bookModel.Book
	FindById(id int) bookModel.Book
}

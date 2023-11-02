package repositories

import (
	bookModel "gitub.com/RomainC75/biblio/internal/modules/book/models"
)

type BookRepositoryInterface interface {
	Create(book bookModel.Book) bookModel.Book
	FirstOrCreateAuthors(newAuthors []string) []bookModel.Author
	FindByISBN(isbn string) (bookModel.Book, error)
	FindById(id string) (bookModel.Book, error)
	FindByUserID(userId string) []bookModel.Book
	DeleteBookById(userId string) (bookModel.Book, error)
	UpdateBookById(userId bookModel.Book) (bookModel.Book, error)
}

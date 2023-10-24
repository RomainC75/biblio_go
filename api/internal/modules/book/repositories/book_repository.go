package repositories

import (
	bookModel "gitub.com/RomainC75/biblio/internal/modules/book/models"
	"gitub.com/RomainC75/biblio/pkg/database"

	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func New() *BookRepository {
	return &BookRepository{
		DB: database.Connection(),
	}
}

func (BookRepository *BookRepository) Create(book bookModel.Book) bookModel.Book {
	var newBook bookModel.Book

	BookRepository.DB.Create(&book).Scan(&newBook)

	return newBook
}

func (BookRepository *BookRepository) FindByEmail(email string) bookModel.Book {
	var foundUser bookModel.Book

	BookRepository.DB.First(&foundUser, "email = ?", email)

	return foundUser
}

func (BookRepository *BookRepository) FindById(id int) bookModel.Book {
	var foundUser bookModel.Book

	BookRepository.DB.First(&foundUser, "id = ?", id)

	return foundUser
}

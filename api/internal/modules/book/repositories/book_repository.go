package repositories

import (
	"errors"
	"fmt"

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
	// BookRepository.DB.Model(&book).Association("Authors")
	// BookRepository.DB.Save(&book)
	return newBook
}

func (BookRepository *BookRepository) FirstOrCreateAuthors(newAuthors []string) []bookModel.Author {
	var authors []bookModel.Author
	for _, newAuthor := range newAuthors {
		var author bookModel.Author

		BookRepository.DB.FirstOrCreate(&author, &bookModel.Author{Name: newAuthor})
		authors = append(authors, author)
	}
	return authors
}

func (BookRepository *BookRepository) FindByISBN(isbn string) (bookModel.Book, error) {
	var foundBook bookModel.Book
	result := BookRepository.DB.First(&foundBook, "ISBN = ?", isbn)
	if result.RowsAffected == 0 {
		return foundBook, errors.New("not found")
	}
	return foundBook, nil
}

func (BookRepository *BookRepository) FindById(id int) bookModel.Book {
	var foundBook bookModel.Book
	BookRepository.DB.First(&foundBook, "id = ?", id)
	return foundBook
}

func (BookRepository *BookRepository) FindByUserID(userId string) []bookModel.Book {
	var foundBooks []bookModel.Book
	fmt.Printf("==> userID : ", userId)
	result := BookRepository.DB.Where("user_refer = ?", userId).Find(&foundBooks)
	if result.Error != nil {
        fmt.Printf("Error: %s", result.Error)
    }
	// utils.PrettyDisplay(foundBooks)
	fmt.Printf("RESULT : ", result.RowsAffected)
	return foundBooks
}


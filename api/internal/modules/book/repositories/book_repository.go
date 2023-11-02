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

func (BookRepository *BookRepository) FindById(id string) (bookModel.Book, error) {
	var foundBook bookModel.Book
	result := BookRepository.DB.First(&foundBook, "id = ?", id)
	if result.RowsAffected == 0 {
		return foundBook, errors.New("book not found")
	}
	return foundBook, nil
}

func (BookRepository *BookRepository) FindByUserID(userId string) []bookModel.Book {
	var foundBooks []bookModel.Book
	result := BookRepository.DB.Where("user_refer = ?", userId).Find(&foundBooks)
	if result.Error != nil {
        fmt.Printf("Error: %s", result.Error)
    }
	// utils.PrettyDisplay(foundBooks)
	return foundBooks
}

func (BookRepository *BookRepository) DeleteBookById(bookId string) (bookModel.Book, error){
	var deletedBook bookModel.Book

	if err := BookRepository.DB.First(&deletedBook, "id = ?", bookId).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return bookModel.Book{}, errors.New("Book not found")
        }
    }

	if err := BookRepository.DB.Delete(&deletedBook).Error; err != nil {
        return bookModel.Book{}, errors.New("problem on delete")
    }
	return deletedBook, nil
}

func (BookRepository *BookRepository) UpdateBookById(book bookModel.Book) (bookModel.Book, error){
	var newBook bookModel.Book

	if err := BookRepository.DB.First(&newBook, "id = ?", book.ID).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return bookModel.Book{}, errors.New("Book not found")
        }
    }

	if err := BookRepository.DB.Model(&newBook).Updates(book).Error; err != nil {
        return bookModel.Book{}, errors.New("problem on update")
    }
	return newBook, nil
}

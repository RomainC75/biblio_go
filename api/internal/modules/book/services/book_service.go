package services

import (
	"errors"

	"github.com/google/uuid"
	"gitub.com/RomainC75/biblio/internal/modules/apis/openlibrary/responses"
	BookModel "gitub.com/RomainC75/biblio/internal/modules/book/models"
	BookRepository "gitub.com/RomainC75/biblio/internal/modules/book/repositories"
	BookRequest "gitub.com/RomainC75/biblio/internal/modules/book/requests"
	"gitub.com/RomainC75/biblio/pkg/utils"
)

type BookService struct {
	bookRepository BookRepository.BookRepositoryInterface
}

func New() *BookService {
	return &BookService{
		bookRepository: BookRepository.New(),
	}
}

func (bookService *BookService) CreateFromSearchResponse(book responses.SearchResponse) (BookModel.Book, error) {
	_, err := bookService.bookRepository.FindByISBN(book.Q)
	if err == nil {
		return BookModel.Book{}, errors.New("isbn already in DB")
	}

	createdAuthors := bookService.bookRepository.FirstOrCreateAuthors(book.Docs[0].Authors)

	newBook := BookModel.Book{
		Authors: createdAuthors,
		Title:   book.Docs[0].Title,
		ISBN:    book.Q,
	}
	result := bookService.bookRepository.Create(newBook)
	result.Authors = createdAuthors

	return result, nil
}

func (bookService *BookService) CreateBook(userId uuid.UUID, book BookRequest.CreateBookRequest) (BookModel.Book, error) {
	createdAuthors := bookService.bookRepository.FirstOrCreateAuthors(book.Authors)
	utils.PrettyDisplay(book)
	newBook := BookModel.Book{
		Authors: createdAuthors,
		Title:   book.Title,
		ISBN:    book.ISBN,
		GenreCode:   book.GenreCode,
		LanguageCode: book.LanguageCode,
		UserRefer: userId,
	}
	result := bookService.bookRepository.Create(newBook)
	
	result.Authors=createdAuthors
	// handle errors
	return result, nil
}

func (bookService *BookService) FindBooksByUserId(userId string) []BookModel.Book {
	foundBooks := bookService.bookRepository.FindByUserID(userId)
	return foundBooks
}

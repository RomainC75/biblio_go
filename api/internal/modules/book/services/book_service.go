package services

import (
	"errors"
	"fmt"

	"gitub.com/RomainC75/biblio/internal/modules/apis/openlibrary/responses"
	BookModel "gitub.com/RomainC75/biblio/internal/modules/book/models"
	BookRepository "gitub.com/RomainC75/biblio/internal/modules/book/repositories"
)

type BookService struct {
	bookRepository BookRepository.BookRepositoryInterface
}

func New() *BookService {
	return &BookService{
		bookRepository: BookRepository.New(),
	}
}

func (bookService *BookService) Create(book responses.SearchResponse) (BookModel.Book, error) {

	_, err := bookService.bookRepository.FindByISBN(book.Q)
	if err == nil {
		fmt.Printf("EEEEEERRRROOOOOOORRRRRR!!!!!!!!!!!!!!!!")
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

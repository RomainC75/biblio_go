package services

import (
	"errors"
	"fmt"

	"gitub.com/RomainC75/biblio/internal/modules/apis/openlibrary/responses"
	BookModel "gitub.com/RomainC75/biblio/internal/modules/book/models"
	BookRepository "gitub.com/RomainC75/biblio/internal/modules/book/repositories"
	BookRequest "gitub.com/RomainC75/biblio/internal/modules/book/requests"
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
		fmt.Printf("EEEEEERRRROOOOOOORRRRRR!!!!!!!!!!!!!!!!")
		return BookModel.Book{}, errors.New("isbn already in DB")
	}

	createdAuthors := bookService.bookRepository.FirstOrCreateAuthors(book.Docs[0].Authors)

	// create Title !! 
	BookModel.Title{
		LanguageCode: 1,
		TitleName: book.Docs[0].Title,
	}//

	newBook := BookModel.Book{
		Authors: createdAuthors,
		// Title:   book.Docs[0].Title,
		Title:   []BookModel.Title{
			// created Title
		}

		ISBN:    book.Q,
	}
	result := bookService.bookRepository.Create(newBook)
	result.Authors = createdAuthors

	return result, nil
}

func (bookService *BookService) CreateBook(book BookRequest.CreateBookRequest) (BookModel.Book, error) {
	createdAuthors := bookService.bookRepository.FirstOrCreateAuthors(book.Authors)

	newBook := BookModel.Book{
		Authors: createdAuthors,
		Title:   book.Title,
		ISBN:    book.ISBN,
		Genre:   book.Genre,
	}
	result := bookService.bookRepository.Create(newBook)
	// handle errors
	return result, nil
}

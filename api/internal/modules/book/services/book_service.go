package services

import (
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

	foundBook, err := bookService.bookRepository.FindByISBN(book.Q)
	if err == nil {
		return BookModel.Book{}, err
	}
	fmt.Printf("--> foundBook : ", foundBook)
	fmt.Printf("--> foundBook.ID :  : ", foundBook.ID)

	// bookModell := responses.ToBookModel(book)
	fmt.Println("-->INSIDE SERVICE------")
	fmt.Println("-->INSIDE SERVICE : ", book)
	fmt.Println("-->INSIDE SERVICE : ", book.Docs[0].Authors[0])
	// create/get authors

	createdAuthors := bookService.bookRepository.FirstOrCreateAuthors(book.Docs[0].Authors)
	fmt.Println("--> createdAuthor : ", createdAuthors)

	newBook := BookModel.Book{
		Authors: createdAuthors,
		Title:   book.Docs[0].Title,
		ISBN:    book.Q,
	}
	result := bookService.bookRepository.Create(newBook)
	result.Authors = createdAuthors
	// newBook := bookService.bookRepository.Create(bookModell)
	// fmt.Print("---->", newBook)
	// if newBook.ID == 0 {
	// 	return response, errors.New("error creating the user")
	// }

	fmt.Println("RESULT, ", result)

	// return BookResponse.ToBook(result), nil
	return result, nil
}

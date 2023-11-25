package services

import (
	ApisHandler "gitub.com/RomainC75/biblio/internal/modules/apis/handler"
	"gitub.com/RomainC75/biblio/internal/modules/apis/third-party/openlibrary/responses"
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

func (bookService *BookService) CreateFromSearchResponse(book responses.SearchResponseData) (BookModel.Book, error) {
// 	_, err := bookService.bookRepository.FindByISBN(book.Q)
// 	if err == nil {
// 		return BookModel.Book{}, errors.New("isbn already in DB")
// 	}

// 	createdAuthors := bookService.bookRepository.FirstOrCreateAuthors(book.Docs[0].Authors)

// 	newBook := BookModel.Book{
// 		Authors: createdAuthors,
// 		Title:   book.Docs[0].Title,
// 		ISBN:    book.Q,
// 	}
// 	result := bookService.bookRepository.Create(newBook)
// 	result.Authors = createdAuthors

// 	return result, nil
	return BookModel.Book{}, nil
}

func (bookService *BookService) CreateNewBook(bookInfos ApisHandler.SearchInApisResponse) (BookModel.Book, error) {
	
	newBook, _ := bookService.bookRepository.CreateBook(bookInfos)

	return newBook, nil
}

func (bookService *BookService) FindBooksByUserId(userId string) []BookModel.Book {
// 	foundBooks := bookService.bookRepository.FindByUserID(userId)
// 	return foundBooks
// }

// func (bookService *BookService) DeleteBook(userId string, bookId string) (BookModel.Book, error) {
// 	foundBook, err := bookService.bookRepository.FindById(bookId)
// 	if err != nil {
// 		return BookModel.Book{}, err
// 	} else if foundBook.UserRefer.String() != userId {
// 		return BookModel.Book{}, errors.New(fmt.Sprintf("unauthorized to delete the book : ", bookId))
// 	}
// 	deletedBook, err := bookService.bookRepository.DeleteBookById(bookId)
// 	if err != nil {
// 		return BookModel.Book{}, err
// 	}
// 	return deletedBook, nil
return []BookModel.Book{}
}

func (bookService *BookService) UpdateBook(userId string, book BookModel.Book) (BookModel.Book, error) {
// 	foundBook, err := bookService.bookRepository.FindById(book.ID.String())
// 	if err != nil {
// 		return BookModel.Book{}, err
// 	} else if foundBook.UserRefer.String() != userId {
// 		return BookModel.Book{}, errors.New(fmt.Sprintf("unauthorized to delete the book : ", book.ID.String()))
// 	}
// 	updatedBook, err := bookService.bookRepository.UpdateBookById(book)
// 	if err != nil {
// 		return BookModel.Book{}, err
// 	}
// 	return updatedBook, nil
return BookModel.Book{}, nil
}

func (bookService *BookService) UpdateAuthorsSrv(authors []BookModel.Author) ([]BookModel.Author, error) {
// 	updatedAuthors, _ := bookService.bookRepository.UpdateAuthors(authors)
// 	return updatedAuthors, nil
return []BookModel.Author{}, nil
}

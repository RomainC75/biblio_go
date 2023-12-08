package services

import (
	Repository "gitub.com/RomainC75/biblio/api/repositories"
	Model "gitub.com/RomainC75/biblio/data/models"
	"gitub.com/RomainC75/biblio/utils/third-party-apis/openlibrary/responses"
	TPApisServices "gitub.com/RomainC75/biblio/utils/third-party-apis/services"
)

type BookService struct {
	bookRepository Repository.BookRepositoryInterface
}

func NewBookSrv() *BookService {
	return &BookService{
		bookRepository: Repository.NewBookRepo(),
	}
}

func (bookService *BookService) CreateFromSearchResponse(book responses.SearchResponseData) (Model.Book, error) {
// 	_, err := bookService.bookRepository.FindByISBN(book.Q)
// 	if err == nil {
// 		return Model.Book{}, errors.New("isbn already in DB")
// 	}

// 	createdAuthors := bookService.bookRepository.FirstOrCreateAuthors(book.Docs[0].Authors)

// 	newBook := Model.Book{
// 		Authors: createdAuthors,
// 		Title:   book.Docs[0].Title,
// 		ISBN:    book.Q,
// 	}
// 	result := bookService.bookRepository.Create(newBook)
// 	result.Authors = createdAuthors

// 	return result, nil
	return Model.Book{}, nil
}

func (bookService *BookService) CreateNewBook(bookInfos TPApisServices.SearchInApisResponse) (Model.Book, error) {
	
	newBook, _ := bookService.bookRepository.CreateBook(bookInfos)

	return newBook, nil
}

func (bookService *BookService) FindBookByIsbnSrv(isbn string) (Model.Book, error) {
	newBook, err := bookService.bookRepository.FindBookByISBN(isbn)
	return newBook, err
}

func (bookService *BookService) FindBooksByUserId(userId string) []Model.Book {
// 	foundBooks := bookService.bookRepository.FindByUserID(userId)
// 	return foundBooks
// }

// func (bookService *BookService) DeleteBook(userId string, bookId string) (Model.Book, error) {
// 	foundBook, err := bookService.bookRepository.FindById(bookId)
// 	if err != nil {
// 		return Model.Book{}, err
// 	} else if foundBook.UserRefer.String() != userId {
// 		return Model.Book{}, errors.New(fmt.Sprintf("unauthorized to delete the book : ", bookId))
// 	}
// 	deletedBook, err := bookService.bookRepository.DeleteBookById(bookId)
// 	if err != nil {
// 		return Model.Book{}, err
// 	}
// 	return deletedBook, nil
return []Model.Book{}
}

func (bookService *BookService) UpdateBook(userId string, book Model.Book) (Model.Book, error) {
// 	foundBook, err := bookService.bookRepository.FindById(book.ID.String())
// 	if err != nil {
// 		return Model.Book{}, err
// 	} else if foundBook.UserRefer.String() != userId {
// 		return Model.Book{}, errors.New(fmt.Sprintf("unauthorized to delete the book : ", book.ID.String()))
// 	}
// 	updatedBook, err := bookService.bookRepository.UpdateBookById(book)
// 	if err != nil {
// 		return Model.Book{}, err
// 	}
// 	return updatedBook, nil
return Model.Book{}, nil
}

func (bookService *BookService) UpdateAuthorsSrv(authors []Model.Author) ([]Model.Author, error) {
// 	updatedAuthors, _ := bookService.bookRepository.UpdateAuthors(authors)
// 	return updatedAuthors, nil
return []Model.Author{}, nil
}

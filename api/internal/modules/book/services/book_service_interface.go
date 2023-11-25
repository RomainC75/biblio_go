package services

import (
	"gitub.com/RomainC75/biblio/internal/modules/apis/third-party/openlibrary/responses"

	// BookResponse "gitub.com/RomainC75/biblio/internal/modules/book/responses"
	BookModel "gitub.com/RomainC75/biblio/internal/modules/book/models"

	ApisHandler "gitub.com/RomainC75/biblio/internal/modules/apis/handler"
)

type BookServiceInterface interface {
	CreateFromSearchResponse(book responses.SearchResponseData) (BookModel.Book, error)
	// CreateBook(uuid.UUID, BookRequest.CreateBookRequest) (BookModel.Book, error)
	CreateNewBook(bookInfos ApisHandler.SearchInApisResponse) (BookModel.Book, error)
	FindBooksByUserId(userId string) ([]BookModel.Book)
	// DeleteBook(userId string, bookId string) (BookModel.Book, error)
	UpdateBook(userId string, book BookModel.Book ) (BookModel.Book, error)
	UpdateAuthorsSrv(authors []BookModel.Author ) ([]BookModel.Author, error)
	// CheckIfUserExists(email string) bool
	// HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error)
}

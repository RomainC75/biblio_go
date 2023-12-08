package services

import (
	"gitub.com/RomainC75/biblio/utils/third-party-apis/openlibrary/responses"

	Models "gitub.com/RomainC75/biblio/data/models"

	TPApisServices "gitub.com/RomainC75/biblio/utils/third-party-apis/services"
)

type BookServiceInterface interface {
	CreateFromSearchResponse(book responses.SearchResponseData) (Models.Book, error)
	// CreateBook(uuid.UUID, BookRequest.CreateBookRequest) (Models.Book, error)
	CreateNewBook(bookInfos TPApisServices.SearchInApisResponse) (Models.Book, error)
	FindBooksByUserId(userId string) ([]Models.Book)
	FindBookByIsbnSrv(isbn string) (Models.Book, error)
	// DeleteBook(userId string, bookId string) (Models.Book, error)
	UpdateBook(userId string, book Models.Book ) (Models.Book, error)
	UpdateAuthorsSrv(authors []Models.Author ) ([]Models.Author, error)
	// CheckIfUserExists(email string) bool
	// HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error)
}

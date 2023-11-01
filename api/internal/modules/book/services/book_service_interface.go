package services

import (
	"github.com/google/uuid"
	"gitub.com/RomainC75/biblio/internal/modules/apis/openlibrary/responses"

	// BookResponse "gitub.com/RomainC75/biblio/internal/modules/book/responses"
	BookModel "gitub.com/RomainC75/biblio/internal/modules/book/models"
	BookRequest "gitub.com/RomainC75/biblio/internal/modules/book/requests"
)

type BookServiceInterface interface {
	CreateFromSearchResponse(book responses.SearchResponse) (BookModel.Book, error)
	CreateBook(uuid.UUID, BookRequest.CreateBookRequest) (BookModel.Book, error)
	FindBooksByUserId(userId string) ([]BookModel.Book)
	// CheckIfUserExists(email string) bool
	// HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error)
}

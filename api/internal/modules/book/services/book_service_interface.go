package services

import (
	"gitub.com/RomainC75/biblio/internal/modules/apis/openlibrary/responses"
	// BookResponse "gitub.com/RomainC75/biblio/internal/modules/book/responses"
	BookModel "gitub.com/RomainC75/biblio/internal/modules/book/models"
	BookRequest "gitub.com/RomainC75/biblio/internal/modules/book/requests"
)

type BookServiceInterface interface {
	CreateFromSearchResponse(book responses.SearchResponse) (BookModel.Book, error)
	CreateBook(book BookRequest.CreateBookRequest) (BookModel.Book, error)
	// CheckIfUserExists(email string) bool
	// HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error)
}

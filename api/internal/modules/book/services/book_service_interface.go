package services

import (
	"gitub.com/RomainC75/biblio/internal/modules/apis/openlibrary/responses"
	// BookResponse "gitub.com/RomainC75/biblio/internal/modules/book/responses"
	BookModel "gitub.com/RomainC75/biblio/internal/modules/book/models"
)

type BookServiceInterface interface {
	Create(book responses.SearchResponse) (BookModel.Book, error)
	// CheckIfUserExists(email string) bool
	// HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error)
}

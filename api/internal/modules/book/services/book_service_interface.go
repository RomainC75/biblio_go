package services

import (
	BookResponse "gitub.com/RomainC75/biblio/internal/modules/book/responses"
	"gitub.com/RomainC75/biblio/internal/modules/user/requests/auth"
)

type BookServiceInterface interface {
	Create(request auth.RegisterRequest) (BookResponse.Book, error)
	// CheckIfUserExists(email string) bool
	// HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error)
}

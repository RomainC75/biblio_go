package services

import (
	"gitub.com/RomainC75/biblio/internal/modules/user/requests/auth"
	UserResponse "gitub.com/RomainC75/biblio/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (UserResponse.User, error)
	CheckIfUserExists(email string) bool
	HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error)
}

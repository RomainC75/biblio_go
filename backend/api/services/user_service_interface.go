package services

import (
	Dto "gitub.com/RomainC75/biblio/api/dto"
	Responses "gitub.com/RomainC75/biblio/api/dto/responses"
)

type UserServiceInterface interface {
	Create(request Dto.RegisterRequest) (Responses.User, error)
	CheckIfUserExists(email string) bool
	HandleUserLogin(request Dto.LoginRequest) (Responses.User, error)
}

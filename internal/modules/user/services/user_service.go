package services

import (
	"errors"
	"fmt"

	userModel "gitub.com/RomainC75/biblio/internal/modules/user/models"
	UserRepository "gitub.com/RomainC75/biblio/internal/modules/user/repositories"
	"gitub.com/RomainC75/biblio/internal/modules/user/requests/auth"
	UserResponse "gitub.com/RomainC75/biblio/internal/modules/user/responses"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository UserRepository.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepository: UserRepository.New(),
	}
}

func (userService *UserService) Create(request auth.RegisterRequest) (UserResponse.User, error) {
	var response UserResponse.User
	var user userModel.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		return response, errors.New("error ashing the password")
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hashedPassword)

	newUser := userService.userRepository.Create(user)
	fmt.Print("---->", newUser)
	if newUser.ID == 0 {
		return response, errors.New("error creating the user")
	}
	return UserResponse.ToUser(newUser), nil
}

func (userService *UserService) CheckIfUserExists(email string) bool {
	user := userService.userRepository.FindByEmail(email)
	return user.ID != 0
}

func (userService *UserService) HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error) {
	var response UserResponse.User
	existUser := userService.userRepository.FindByEmail(request.Email)

	if existUser.ID == 0 {
		return response, errors.New("Invalid Credentials !")
	}

	err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(request.Password))
	if err != nil {
		return response, errors.New("invalid credentials !")
	}

	return UserResponse.ToUser(existUser), nil

}

package services

import (
	"errors"
	"fmt"

	Dto "gitub.com/RomainC75/biblio/api/dto"
	Responses "gitub.com/RomainC75/biblio/api/dto/responses"
	Repositories "gitub.com/RomainC75/biblio/api/repositories"
	Models "gitub.com/RomainC75/biblio/data/models"
	"gitub.com/RomainC75/biblio/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository Repositories.UserRepositoryInterface
}

func NewUserSrv() *UserService {
	return &UserService{
		userRepository: Repositories.NewUserRepo(),
	}
}

func (userService *UserService) Create(request Dto.RegisterRequest) (Responses.User, error) {
	var response Responses.User
	var user Models.User
	fmt.Println("---------> create a new User !")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		return response, errors.New("error ashing the password")
	}

	user.Email = request.Email
	user.Password = string(hashedPassword)

	newUser := userService.userRepository.Create(user)
	fmt.Print("---->", newUser)
	if newUser.ID == utils.InitUUID() {
		return response, errors.New("error creating the user")
	}
	return Responses.ToUser(newUser), nil
}

func (userService *UserService) CheckIfUserExists(email string) bool {
	user := userService.userRepository.FindByEmail(email)
	return user.ID != utils.InitUUID()
}

func (userService *UserService) HandleUserLogin(request Dto.LoginRequest) (Responses.User, error) {
	var response Responses.User
	existUser := userService.userRepository.FindByEmail(request.Email)

	if existUser.ID == utils.InitUUID() {
		return response, errors.New("Invalid Credentials !")
	}

	err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(request.Password))
	if err != nil {
		return response, errors.New("invalid credentials !")
	}

	return Responses.ToUser(existUser), nil

}

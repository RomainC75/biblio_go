package services

import (
	"fmt"

	bookModel "gitub.com/RomainC75/biblio/internal/modules/book/models"
	BookRepository "gitub.com/RomainC75/biblio/internal/modules/book/repositories"
	BookResponse "gitub.com/RomainC75/biblio/internal/modules/book/responses"
	"gitub.com/RomainC75/biblio/internal/modules/user/requests/auth"
)

type BookService struct {
	bookRepository BookRepository.BookRepositoryInterface
}

func New() *BookService {
	return &BookService{
		bookRepository: BookRepository.New(),
	}
}

func (bookService *BookService) Create(request auth.RegisterRequest) (BookResponse.Book, error) {
	// var response BookResponse.Book
	var book bookModel.Book
	fmt.Println("---------> create a new User !")

	newBook := bookService.bookRepository.Create(book)
	fmt.Print("---->", newBook)
	// if newBook.ID == 0 {
	// 	return response, errors.New("error creating the user")
	// }
	return BookResponse.ToBook(newBook), nil

	// return BookResponse.ToUser(newUser), nil
}

// func (userService *UserService) CheckIfUserExists(email string) bool {
// 	user := userService.userRepository.FindByEmail(email)
// 	return user.ID != 0
// }

// func (userService *UserService) HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error) {
// 	var response UserResponse.User
// 	existUser := userService.userRepository.FindByEmail(request.Email)

// 	if existUser.ID == 0 {
// 		return response, errors.New("Invalid Credentials !")
// 	}

// 	err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(request.Password))
// 	if err != nil {
// 		return response, errors.New("invalid credentials !")
// 	}

// 	return UserResponse.ToUser(existUser), nil

// }

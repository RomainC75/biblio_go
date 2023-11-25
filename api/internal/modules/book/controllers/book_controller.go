package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	ApisHandler "gitub.com/RomainC75/biblio/internal/modules/apis/handler"
	BookService "gitub.com/RomainC75/biblio/internal/modules/book/services"
	UserService "gitub.com/RomainC75/biblio/internal/modules/user/services"
	"gitub.com/RomainC75/biblio/pkg/utils"
)

type Controller struct {
	userService UserService.UserServiceInterface
	bookService BookService.BookServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: UserService.New(),
		bookService: BookService.New(),
	}
}

func (controller *Controller) SearchBook(c *gin.Context) {
	isbn := c.Query("isbn")

	foundBook, err := controller.bookService.FindBookByIsbnSrv(isbn)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"foundBook": foundBook})	
		return 
	}

	compilated, err := ApisHandler.SearchInApis(isbn)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	}

	newBook, _ := controller.bookService.CreateNewBook(compilated)
	
	fmt.Println("==========================> OUT FROM DB ")
	utils.PrettyDisplay(newBook)

	c.JSON(http.StatusOK, gin.H{"search": newBook})
}


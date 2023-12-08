package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	Handlers "gitub.com/RomainC75/biblio/api/handlers"
	Services "gitub.com/RomainC75/biblio/api/services"
	"gitub.com/RomainC75/biblio/utils"
)

type BookController struct {
	userService Services.UserServiceInterface
	bookService Services.BookServiceInterface
}

func NewBookCtrl() *BookController {
	return &BookController{
		userService: Services.NewUserSrv(),
		bookService: Services.NewBookSrv(),
	}
}

func (controller *BookController) SearchBook(c *gin.Context) {
	isbn := c.Query("isbn")

	foundBook, err := controller.bookService.FindBookByIsbnSrv(isbn)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"foundBook": foundBook})	
		return 
	}

	compilated, err := Handlers.SearchInApis(isbn)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	}

	newBook, _ := controller.bookService.CreateNewBook(compilated)
	
	fmt.Println("==========================> OUT FROM DB ")
	utils.PrettyDisplay(newBook)

	c.JSON(http.StatusOK, gin.H{"search": newBook})
}


package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	Responses "gitub.com/RomainC75/biblio/api/dto/responses"
	Services "gitub.com/RomainC75/biblio/api/services"
	"gitub.com/RomainC75/biblio/utils"
	FilterHelper "gitub.com/RomainC75/biblio/utils/filter"
	TPApisServices "gitub.com/RomainC75/biblio/utils/third-party-apis/services"
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
	rawIsbn := c.Query("isbn")
	isbn, err := FilterHelper.FilterISBN(rawIsbn)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})	
		return
	}

	foundBook, err := controller.bookService.FindBookByIsbnSrv(isbn)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"book": Responses.ToBookResponse(foundBook)})	
		return 
	}

	compilated, err := TPApisServices.SearchInApis(isbn)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("==> Compilated : ")
	utils.PrettyDisplay(compilated)

	newBook, _ := controller.bookService.CreateNewBook(compilated)
	
	fmt.Println("==========================> OUT FROM DB ")
	utils.PrettyDisplay(newBook)


	c.JSON(http.StatusOK, gin.H{"book": Responses.ToBookResponse(newBook)})
}


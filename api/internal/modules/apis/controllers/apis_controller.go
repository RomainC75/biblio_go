package controllers

import (
	"fmt"
	"net/http"

	// UserService "gitub.com/RomainC75/biblio/internal/modules/user/services"
	"github.com/gin-gonic/gin"
	OpenLibraryService "gitub.com/RomainC75/biblio/internal/modules/apis/openlibrary/services"
	BookRequest "gitub.com/RomainC75/biblio/internal/modules/book/requests"
	BookService "gitub.com/RomainC75/biblio/internal/modules/book/services"
	Utils "gitub.com/RomainC75/biblio/pkg/utils"
)

type Controller struct {
	bookService BookService.BookServiceInterface
}

func New() *Controller {
	return &Controller{
		bookService: BookService.New(),
	}
}

func (controller *Controller) Search(c *gin.Context) {

	foundBook, err := OpenLibraryService.SearchByReq("0552778079")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	book, err := controller.bookService.CreateFromSearchResponse(foundBook)
	fmt.Println("inside controller ", book, err)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": book})
}

func (controller *Controller) CreateNewBook(c *gin.Context) {
	var request BookRequest.CreateBookRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	Utils.PrettyDisplay(request)

	book, _ := controller.bookService.CreateBook(request)

	c.JSON(http.StatusBadGateway, gin.H{"message": book})

}

func (controller *Controller) GetBooks(c *gin.Context) {

}

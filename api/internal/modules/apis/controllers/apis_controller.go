package controllers

import (
	"fmt"
	"net/http"

	// UserService "gitub.com/RomainC75/biblio/internal/modules/user/services"
	"github.com/gin-gonic/gin"
	OpenLibraryService "gitub.com/RomainC75/biblio/internal/modules/apis/openlibrary/services"
	BookService "gitub.com/RomainC75/biblio/internal/modules/book/services"
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
	book, err := controller.bookService.Create(foundBook)
	fmt.Println("inside controller ", book, err)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": book})
}

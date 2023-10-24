package controllers

import (
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

	OpenLibraryService.SearchByReq("0860519600")

	c.JSON(http.StatusOK, gin.H{"message": "Search Route"})

}

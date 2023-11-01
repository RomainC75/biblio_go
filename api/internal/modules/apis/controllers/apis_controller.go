package controllers

import (
	"fmt"
	"net/http"

	// UserService "gitub.com/RomainC75/biblio/internal/modules/user/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func (controller *Controller) CreateNewBook(c *gin.Context) {
	var request BookRequest.CreateBookRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	Utils.PrettyDisplay(request)

	userId, _ := c.Get("user_id")
	userUuid, _ := uuid.Parse(userId.(string))

	book, _ := controller.bookService.CreateBook(userUuid, request)

	c.JSON(http.StatusBadGateway, gin.H{"message": book})
}

func (controller *Controller) GetBooks(c *gin.Context) {
	userId, _ := c.Get("user_id")
	if userIdStr, ok := userId.(string); ok {
		foundBooks := controller.bookService.FindBooksByUserId(userIdStr)
		c.JSON(http.StatusBadRequest, gin.H{"found books": foundBooks})
		return 
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "error converting the user id"})
		return 
}

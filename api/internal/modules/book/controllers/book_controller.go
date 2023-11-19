package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	ApisHandler "gitub.com/RomainC75/biblio/internal/modules/apis/handler"
	UserService "gitub.com/RomainC75/biblio/internal/modules/user/services"
)

type Controller struct {
	userService UserService.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: UserService.New(),
		
	}
}


func (controller *Controller) SearchBook(c *gin.Context) {
	isbn := c.Query("isbn")
	fmt.Println("ISBN : ", isbn)

	ApisHandler.SearchInApis(isbn)

	c.JSON(http.StatusOK, gin.H{"search": "book"})
}


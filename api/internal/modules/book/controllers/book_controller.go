package controllers

import (
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

	compilated, err := ApisHandler.SearchInApis(isbn)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"search": compilated})
}


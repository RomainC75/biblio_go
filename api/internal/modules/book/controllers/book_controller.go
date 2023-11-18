package controllers

import (
	"fmt"
	"net/http"

	UserService "gitub.com/RomainC75/biblio/internal/modules/user/services"

	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusOK, gin.H{"search": "book"})
}


package controllers

import (
	"net/http"

	// UserService "gitub.com/RomainC75/biblio/internal/modules/user/services"
	// OpenLibraryService "gitub.com/RomainC75/biblio/internal/modules/apis/openlibrary/services"
	"github.com/gin-gonic/gin"
	"gitub.com/RomainC75/biblio/internal/modules/apis/openlibrary/services"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{
		// openLibraryService: OpenLibraryService.New(),
	}
}

func (controller *Controller) Search(c *gin.Context) {
	// services.Search("")
	services.SearchByReq("0860519600")

	c.JSON(http.StatusOK, gin.H{"message": "Search Route"})
	// html.Render(c, http.StatusOK, "modules/user/html/login", gin.H{
	// 	"title": "Login",
	// })
}

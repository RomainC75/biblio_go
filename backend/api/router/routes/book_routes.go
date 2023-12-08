package routes

import (
	bookCtrl "gitub.com/RomainC75/biblio/internal/modules/book/controllers"

	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {

	userController := bookCtrl.New()
	guestGroup := router.Group("/books")
	{
		guestGroup.GET("/search", userController.SearchBook)
	}

	// authGroup := router.Group("/")
	// authGroup.Use()
	// {
		// router.GET("/test", middlewares.IsAuth(), userController.HandleWhoAmI)
	// }
}

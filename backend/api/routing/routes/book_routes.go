package routes

import (
	Ctrl "gitub.com/RomainC75/biblio/api/controllers"

	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {

	userController := Ctrl.NewBookCtrl()
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

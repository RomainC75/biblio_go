package routes

import (
	bookCtrl "gitub.com/RomainC75/biblio/internal/modules/book/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	userController := bookCtrl.New()
	guestGroup := router.Group("/books")
	// guestGroup.Use(middlewares.IsAuth())
	{
		// guestGroup.POST("/register", userController.HandleRegister)
		// guestGroup.POST("/login", userController.HandleLogin)
		guestGroup.GET("/search", userController.SearchBook)
	}

	// authGroup := router.Group("/")
	// authGroup.Use()
	// {
		// router.GET("/test", middlewares.IsAuth(), userController.HandleWhoAmI)
	// }
}

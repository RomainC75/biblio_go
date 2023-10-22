package routes

import (
	apiCtrl "gitub.com/RomainC75/biblio/internal/modules/apis/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	// userController := userCtrl.New()

	// guestGroup := router.Group("/")
	// guestGroup.Use(middlewares.IsAuth())
	// {
	// 	guestGroup.GET("/register", userController.Register)
	// 	guestGroup.POST("/register", userController.HandleRegister)

	// 	guestGroup.GET("/login", userController.Login)
	// 	guestGroup.POST("/login", userController.HandleLogin)
	// }

	apiController := apiCtrl.New()
	router.GET("/search", apiController.Search)

}

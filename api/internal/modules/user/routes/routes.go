package routes

import (
	middlewares "gitub.com/RomainC75/biblio/internal/middleware"
	userCtrl "gitub.com/RomainC75/biblio/internal/modules/user/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	userController := userCtrl.New()

	guestGroup := router.Group("/")
	guestGroup.Use(middlewares.IsAuth())
	{
		guestGroup.POST("/register", userController.HandleRegister)

		guestGroup.POST("/login", userController.HandleLogin)
	}

	authGroup := router.Group("/")
	authGroup.Use(middlewares.IsAuth())
	{
		router.GET("/test", userController.HandleTest)
	}

}

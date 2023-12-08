package routes

import (
	Ctrl "gitub.com/RomainC75/biblio/api/controllers"
	middlewares "gitub.com/RomainC75/biblio/api/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {

	userController := Ctrl.NewAuthCtrl()
	guestGroup := router.Group("/auth")
	// guestGroup.Use(middlewares.IsAuth())
	{
		guestGroup.POST("/register", userController.HandleRegister)
		guestGroup.POST("/login", userController.HandleLogin)
	}

	authGroup := router.Group("/")
	authGroup.Use()
	{
		router.GET("/test", middlewares.IsAuth(), userController.HandleWhoAmI)
	}
}

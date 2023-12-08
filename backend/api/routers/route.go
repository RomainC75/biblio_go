package routes

import (
	Routes "gitub.com/RomainC75/biblio/api/routers/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	Routes.UserRoutes(router)
	Routes.BookRoutes(router)
}

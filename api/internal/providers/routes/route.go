package routes

import (
	bookRoutes "gitub.com/RomainC75/biblio/internal/modules/book/routes"
	userRoutes "gitub.com/RomainC75/biblio/internal/modules/user/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	userRoutes.Routes(router)
	bookRoutes.Routes(router)
}

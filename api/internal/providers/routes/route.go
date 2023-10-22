package routes

import (
	// articlesRoutes "gitub.com/RomainC75/biblio/internal/modules/article/routes"
	// homeRoutes "gitub.com/RomainC75/biblio/internal/modules/home/routes"
	userRoutes "gitub.com/RomainC75/biblio/internal/modules/user/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// homeRoutes.Routes(router)
	// articlesRoutes.Routes(router)
	userRoutes.Routes(router)

}

package routes

import (
	apiRoutes "gitub.com/RomainC75/biblio/internal/modules/apis/routes"
	userRoutes "gitub.com/RomainC75/biblio/internal/modules/user/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// homeRoutes.Routes(router)
	// articlesRoutes.Routes(router)
	userRoutes.Routes(router)
	apiRoutes.Routes(router)

}

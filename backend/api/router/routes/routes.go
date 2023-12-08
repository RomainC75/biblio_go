package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	UserRoutes(router)
	BookRoutes(router)
}

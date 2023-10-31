package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gitub.com/RomainC75/biblio/pkg/jwt"
	"gitub.com/RomainC75/biblio/pkg/utils"
)

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("'''''''''IS AUTH''''''''''")

		auth_header, ok := c.Request.Header["Authorization"]
		if !ok || !strings.HasPrefix(auth_header[0], "Bearer") {
			c.JSON(http.StatusBadRequest, gin.H{"message": "token missing"})
			c.Abort()
			return
		}
		token := strings.Split(auth_header[0], " ")[1]
		fmt.Println("1")
		claim, err := jwt.GetClaimsFromToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauhorized"})
			c.Abort()
			return
		}
		fmt.Println("--------> token claim : ")
		utils.PrettyDisplay(claim)

		c.Set("email", "mlksjdf")
		fmt.Print(">>>>>>>>>>>>>>>>>IsAuth middleware<<<<<<<<<<<<<<<<<<<")
		c.Next()
	}
}

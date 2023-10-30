package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func IsAuth() gin.HandlerFunc {

	// var userRepo = UserRepository.New()

	return func(c *gin.Context) {
		// userID, _ := strconv.Atoi(authID)

		// user := userRepo.FindById(userID)

		// if user.ID == 0 {
		// 	c.Redirect(http.StatusFound, "/login")
		// 	return
		// }

		fmt.Print(">>>>>>>>>>>>>>>>>IsAuth middleware<<<<<<<<<<<<<<<<<<<")
		c.Next()
	}
}

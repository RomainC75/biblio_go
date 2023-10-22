package middlewares

import (
	"github.com/gin-gonic/gin"
)

func IsGuest() gin.HandlerFunc {

	// var userRepo = UserRepository.New()

	return func(c *gin.Context) {
		// authID := sessions.Get(c, "auth")
		// userID, _ := strconv.Atoi(authID)

		// user := userRepo.FindById(userID)

		// if user.ID == 0 {
		// 	c.Redirect(http.StatusFound, "/")
		// 	return
		// }

		c.Next()
	}
}

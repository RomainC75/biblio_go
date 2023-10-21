package helpers

import (
	UserResponse "gitub.com/RomainC75/biblio/internal/modules/user/responses"

	// "gitub.com/RomainC75/biblio/pkg/sessions"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) UserResponse.User {
	var response UserResponse.User

	// authID := sessions.Get(c, "auth")
	// userID, _ := strconv.Atoi(authID)

	// var userRepo = UserRepository.New()
	// user := userRepo.FindById(userID)

	// if user.ID == 0 {
	// 	return response
	// }

	// return UserResponse.ToUser(user)
	return response
}

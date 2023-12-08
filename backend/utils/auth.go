package utils

import (
	Responses "gitub.com/RomainC75/biblio/api/dto/responses"

	// "gitub.com/RomainC75/biblio/pkg/sessions"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) Responses.User {
	var response Responses.User

	// authID := sessions.Get(c, "auth")
	// userID, _ := strconv.Atoi(authID)

	// var userRepo = UserRepository.New()
	// user := userRepo.FindById(userID)

	// if user.ID == 0 {
	// 	return response
	// }

	// return Responses.ToUser(user)
	return response
}

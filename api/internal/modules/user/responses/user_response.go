package responses

import (
	"fmt"

	"github.com/google/uuid"
	userModel "gitub.com/RomainC75/biblio/internal/modules/user/models"
)

type User struct {
	ID    uuid.UUID
	Image string
	Name  string
	Email string
}

type Users struct {
	Data []User
}

func ToUser(user userModel.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: fmt.Sprintf("https://ui-avatars.com/api/?name=%s", user.Name),
	}
}

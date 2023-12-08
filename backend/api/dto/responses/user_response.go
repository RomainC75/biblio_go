package responses

import (
	"fmt"

	"github.com/google/uuid"
	Models "gitub.com/RomainC75/biblio/data/models"
)

type User struct {
	ID    uuid.UUID
	Image string
	Firstname  string
	Lastname  string
	Email string
}

type Users struct {
	Data []User
}

func ToUser(user Models.User) User {
	return User{
		ID:    user.ID,
		Firstname:  user.Firstname,
		Lastname:  user.Lastname,
		Email: user.Email,
		Image: fmt.Sprintf("https://ui-avatars.com/api/?name=%s", user.Firstname),
	}
}

package services

import (
	OpenlibraryResponse "gitub.com/RomainC75/biblio/internal/modules/apis/openlibrary/responses"

	"gitub.com/RomainC75/biblio/internal/modules/user/requests/auth"
)

type UserServiceInterface interface {
	Search(request auth.RegisterRequest) (OpenlibraryResponse.Search, error)
}

package jwt

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	Responses "gitub.com/RomainC75/biblio/internal/modules/user/responses"
	"gitub.com/RomainC75/biblio/pkg/configu"
)

type Claims struct {
	*jwt.RegisteredClaims
	UserInfo interface{}
}

type BasicUserInfo struct {
	ID    uint
	Email string
}

func Generate(user Responses.User) (string, error) {
	configs := configu.Get()

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	exp := time.Now().Add(time.Hour * 24)

	token.Claims = &Claims{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			Subject:   strconv.FormatUint(uint64(user.ID), 10),
		},
		user,
	}
	val, err := token.SignedString([]byte(configs.Jwt.Secret))

	if err != nil {
		return "", err
	}
	return val, nil
}

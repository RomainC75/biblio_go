package jwt

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	Responses "gitub.com/RomainC75/biblio/internal/modules/user/responses"
)

type Claims struct {
	*jwt.RegisteredClaims
	UserInfo interface{}
}

type BasicUserInfo struct {
	ID    uint
	Email string
}

var secret = []byte("can-you-keep-a-secret?")

func Generate(user Responses.User) (string, error) {

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	exp := time.Now().Add(time.Hour * 24)

	token.Claims = &Claims{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			Subject:   strconv.FormatUint(uint64(user.ID), 10),
		},
		user,
	}
	val, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}
	return val, nil
}

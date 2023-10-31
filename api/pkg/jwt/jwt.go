package jwt

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	Responses "gitub.com/RomainC75/biblio/internal/modules/user/responses"
	"gitub.com/RomainC75/biblio/pkg/configu"
	"gitub.com/RomainC75/biblio/pkg/utils"
)

type Claims struct {
	*jwt.RegisteredClaims
	ID    uint
	Email string
}

func Generate(user Responses.User) (string, error) {
	secret := configu.Get().Jwt.Secret

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	exp := time.Now().Add(time.Hour * 24)

	token.Claims = &Claims{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			Subject:   strconv.FormatUint(uint64(user.ID), 10),
		},
		user.ID,
		user.Email,
	}
	val, err := token.SignedString([]byte(secret))

	if err != nil {
		return "error trying to set the token", err
	}
	return val, nil
}

func GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	secret := configu.Get().Jwt.Secret
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("ZERROR : ", err.Error())
		return nil, err
	}
	utils.PrettyDisplay(token)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("claims : ", claims)
		return claims, nil
	}
	return nil, err
}

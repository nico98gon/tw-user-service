package jwt

import (
	"errors"
	"strings"
	"user-service/internal/domain"
	"user-service/internal/infrastructure/db"

	"github.com/golang-jwt/jwt/v5"
)

var Email string
var IDUsuario string

func ProcessToken(token string, JWTSign string) (*domain.Claim, bool, string, error) {
	myKey := []byte(JWTSign)
	var claims domain.Claim

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, "", errors.New("formato de token inválido")
	}
	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		_, found, _ := db.UserAlreadyExists(claims.Email)
		if found {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return &claims, true, IDUsuario, nil
	}
	if !tkn.Valid {
		return &claims, false, "", errors.New("token inválido")
	}

	Email = claims.Email
	IDUsuario = claims.ID.Hex()

	return &claims, false, "", nil
}
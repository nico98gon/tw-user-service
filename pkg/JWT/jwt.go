package jwt

import (
	"context"
	"errors"
	"time"
	"user-service/internal/domain"
	"user-service/internal/domain/users"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(ctx context.Context, u users.User) (string, error) {
	jwtSign, ok := ctx.Value(domain.Key("jwtSign")).(string)
	if !ok {
		return "", errors.New("no se pudo encontrar jwt_sign")
	}
	myKey := []byte(jwtSign)

	payload := jwt.MapClaims{
		"_id":	u.ID.Hex(),
		"email":	u.Email,
		"name":	u.Name,
		"last_name":	u.LastName,
		"birthdate":	u.Birthdate,
		"bio":	u.Bio,
		"web_site":	u.WebSite,
		"location":	u.Location,
		"created_at":	u.CreatedAt,
		"exp":	time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
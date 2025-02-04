package handlers

import (
	"context"
	"fmt"
	"os"
	"strings"
	"user-service/internal/domain"

	jwt "user-service/pkg/JWT"

	"github.com/aws/aws-lambda-go/events"
)

func checkAuth(ctx context.Context, request events.APIGatewayProxyRequest) (isOk bool, statusCode int, msg string, claim *domain.Claim) {
	path := ctx.Value(domain.Key("path")).(string)
	fmt.Println("path:", path)
	validPaths := []string{"register", "login", "get-profile", "get-avatar", "get-banner"}
	for _, validPath := range validPaths {
		if strings.Contains(path, validPath) {
			return true, 200, "OK", &domain.Claim{}
		}
	}

	var token string

	if os.Getenv("APP_ENV") == "local" {
		for key, value := range request.Headers {
			if strings.ToLower(key) == "authorization" {
				token = value
				break
			}
		}
	} else {
		token = request.Headers["Authorization"]
		fmt.Println("Token en lambda: ", token)
	}

	if len(token) == 0 {
		fmt.Println("path:", path)
		fmt.Println("Token no encontrado en el encabezado de la solicitud")
		return false, 401, "Unauthorized: Token requerido", &domain.Claim{}
	}
	if !strings.HasPrefix(token, "Bearer") {
    fmt.Println("Error: No empieza con 'Bearer'", token)
    return false, 401, "Unauthorized: Formato de token incorrecto", &domain.Claim{}
	}

	claim, isOk, msg, err := jwt.ProcessToken(token, ctx.Value(domain.Key("jwtSign")).(string))
	if !isOk {
		if err != nil {
			fmt.Println("Error en el token: ", err)
			return false, 401, "Unauthorized: Token inválido", &domain.Claim{}
		} else {
			fmt.Println("Error en el token: ", msg)
			return false, 401, msg, &domain.Claim{}
		}
	}

	fmt.Println("Token OK - ID del usuario en el token:", claim.ID.Hex())

	fmt.Println("Token OK")
	return true, 200, "OK", claim
}
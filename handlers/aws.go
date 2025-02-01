package handlers

import (
	"context"
	"fmt"
	"user-service/internal/domain"
	"user-service/internal/infrastructure/routers"
	jwt "user-service/pkg/JWT"

	"github.com/aws/aws-lambda-go/events"
)

func AwsHandler(ctx context.Context, request events.APIGatewayProxyRequest) domain.RespAPI {
	// Envía a CloudWatch Logs
	fmt.Println("Procesando"+ctx.Value(domain.Key("path")).(string) + " > " + ctx.Value(domain.Key("method")).(string))

	var r domain.RespAPI
	r.Status = 400

	isOk, statusCode, msg, _ := checkAuth(ctx, request)
	if !isOk {
		r.Status = statusCode
		r.Message = msg
		return r
	}

	switch ctx.Value(domain.Key("method")).(string) {
	case "GET":
		switch ctx.Value(domain.Key("path")).(string) {
		
		}
		//
	case "POST":
		switch ctx.Value(domain.Key("path")).(string) {
		case "register":
			return routers.Register(ctx)
		
		}
		//
	case "PUT":
		switch ctx.Value(domain.Key("path")).(string) {
		
		}
		//
	case "DELETE":
		switch ctx.Value(domain.Key("path")).(string) {
		
		}
		//
	}

	r.Message = "Method Invalid"
	return r
}

func checkAuth(ctx context.Context, request events.APIGatewayProxyRequest) (isOk bool, statusCode int, msg string, claim *domain.Claim) {
	path := ctx.Value(domain.Key("path")).(string)
	if path == "/register" || path == "/login" || path == "/get-avatar" || path == "/get-banner" {
		return true, 200, "OK", &domain.Claim{}
	}

	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Unauthorized: Token requerido", &domain.Claim{}
	}

	claim, isOk, msg, err := jwt.ProcessToken(token, ctx.Value(domain.Key("JWTSign")).(string))
	if !isOk {
		if err != nil {
			fmt.Println("Error en el token: ", err)
			return false, 401, "Unauthorized: Token inválido", &domain.Claim{}
		} else {
			fmt.Println("Error en el token: ", msg)
			return false, 401, msg, &domain.Claim{}
		}
	}

	fmt.Println("Token OK")
	return true, 200, "OK", claim
}
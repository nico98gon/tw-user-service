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
	fmt.Println("Procesando:", ctx.Value(domain.Key("path")).(string), ">", ctx.Value(domain.Key("method")).(string))

	var r domain.RespAPI
	r.Status = 400

	isOk, statusCode, msg, _ := checkAuth(ctx, request)
	if !isOk {
		fmt.Println("Falló la autenticación:", msg)
		r.Status = statusCode
		r.Message = msg
		return r
	}

	fmt.Println("Autenticación exitosa")

	switch ctx.Value(domain.Key("method")).(string) {
	case "GET":
		fmt.Println("Método GET detectado")
	case "POST":
		fmt.Println("Método POST detectado")
		switch ctx.Value(domain.Key("path")).(string) {
		case "register":
			fmt.Println("Procesando registro de usuario...")
			r = routers.Register(ctx)
			fmt.Println("Registro finalizado:", r.Message)
			return r
		case "login":
			fmt.Println("Procesando inicio de sesión...")
			r = routers.Login(ctx)
			fmt.Println("Inicio de sesión finalizado:", r.Message)
			return r
		}
	}

	fmt.Println("Método inválido detectado")
	r.Message = "Method Invalid"
	return r
}

func checkAuth(ctx context.Context, request events.APIGatewayProxyRequest) (isOk bool, statusCode int, msg string, claim *domain.Claim) {
	path := ctx.Value(domain.Key("path")).(string)
	if path == "register" || path == "login" || path == "get-avatar" || path == "get-banner" {
		return true, 200, "OK", &domain.Claim{}
	}

	token := request.Headers["Authorization"]
	if len(token) == 0 {
			fmt.Println("path:", path)
			fmt.Println("Token no encontrado en el encabezado de la solicitud")
			return false, 401, "Unauthorized: Token requerido", &domain.Claim{}
	}

	fmt.Println("Token recibido:", token)

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
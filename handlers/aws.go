package handlers

import (
	"context"
	"fmt"
	"user-service/internal/domain"
	"user-service/internal/infrastructure/routers"

	"github.com/aws/aws-lambda-go/events"
)

func AwsHandler(ctx context.Context, request events.APIGatewayProxyRequest) domain.RespAPI {
	fmt.Println("Procesando:", ctx.Value(domain.Key("path")).(string), ">", ctx.Value(domain.Key("method")).(string))

	var r domain.RespAPI
	r.Status = 400

	isOk, statusCode, msg, claim := checkAuth(ctx, request)
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
		switch ctx.Value(domain.Key("path")).(string) {
		case "get-profile":
			fmt.Println("Procesando perfil de usuario...")
			r = routers.Profile(request, claim)
			fmt.Println("Perfil de usuario finalizado:", r.Message)
			return r
		
		case "get-avatar":
			fmt.Println("Procesando avatar de usuario...")
			r = routers.GetImage(ctx, "A", request)
			fmt.Println("Avatar de usuario finalizado:", r.Message)
			return r

		case "get-banner":
			fmt.Println("Procesando banner de usuario...")
			r = routers.GetImage(ctx, "B", request)
			fmt.Println("Banner de usuario finalizado:", r.Message)
			return r

		case "get-relation":
			fmt.Println("Procesando relaciones de usuario...")
			r = routers.GetRelation(request, claim)
			fmt.Println("Relaciones de usuario finalizado:", r.Message)
			return r

		}

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

		case "upload-avatar":
			fmt.Println("Procesando carga de avatar...")
			r = routers.UploadImage(ctx, "A", request, claim)
			fmt.Println("Carga de avatar finalizada:", r.Message)
			return r

		case "upload-banner":
			fmt.Println("Procesando carga de banner...")
			r = routers.UploadImage(ctx, "B", request, claim)
			fmt.Println("Carga de banner finalizada:", r.Message)
			return r

		case "new-relation":
			fmt.Println("Procesando registro de relación...")
			r = routers.RegisterRelation(ctx, request, claim)
			fmt.Println("Registro de relación finalizado:", r.Message)
			return r

		}

	case "PUT":
		fmt.Println("Método PUT detectado")
		switch ctx.Value(domain.Key("path")).(string) {
		case "update-profile":
			fmt.Println("Procesando actualización de perfil de usuario...")
			r = routers.UpdateProfile(ctx, claim)
			fmt.Println("Actualización de perfil de usuario finalizada:", r.Message)
			return r
		}

		case "DELETE":
			fmt.Println("Método DELETE detectado")
			switch ctx.Value(domain.Key("path")).(string) {
			case "delete-relation":
				fmt.Println("Procesando eliminación de relación...")
				r = routers.DeleteRelation(request, claim)
				return r
			}
	}

	fmt.Println("Método inválido detectado")
	r.Message = "Method Invalid"
	return r
}
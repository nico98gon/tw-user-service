package services

import (
	"context"
	"fmt"
	"os"
	"strings"
	"user-service/handlers"
	"user-service/internal/domain"
	"user-service/internal/infrastructure/db"
	"user-service/internal/utils"
	"user-service/pkg/aws"
	secretmanager "user-service/pkg/secret-manager"

	"github.com/aws/aws-lambda-go/events"
)

func LambdaExec(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	aws.StartAWS()

	if !validateParams() {
		return utils.HandleError(400, "Error en las variables de entorno. Deben incluir 'SECRET_NAME', 'BUCKET_NAME' y 'URL_PREFIX'"), nil
	}

	SecretModels, err := secretmanager.GetSecret(os.Getenv("SECRET_NAME"))
	if err != nil {
		return utils.HandleError(400, "Error al obtener secret: "+err.Error()), nil
	}

	path := strings.Replace(request.PathParameters["twitteruala"], os.Getenv("URL_PREFIX"), "", -1)
	path = strings.TrimPrefix(path, "/")

	if path == "" {
		path = strings.TrimPrefix(request.Path, "/development/")
		path = strings.TrimPrefix(path, "/")
	}

	if (request.HTTPMethod == "POST" || request.HTTPMethod == "PUT" || request.HTTPMethod == "PATCH") && request.Body == "" {
		return utils.HandleError(400, "El cuerpo de la solicitud no puede estar vacío"), nil
	}

	if SecretModels.JWTSign == "" {
		return nil, fmt.Errorf("JWTSign no encontrado en SecretModels")
	}

	aws.Ctx = context.WithValue(aws.Ctx, domain.Key("path"), path)
	aws.Ctx = context.WithValue(aws.Ctx, domain.Key("method"), request.HTTPMethod)
	aws.Ctx = context.WithValue(aws.Ctx, domain.Key("body"), request.Body)
	aws.Ctx = context.WithValue(aws.Ctx, domain.Key("user"), SecretModels.Username)
	aws.Ctx = context.WithValue(aws.Ctx, domain.Key("password"), SecretModels.Password)
	aws.Ctx = context.WithValue(aws.Ctx, domain.Key("host"), SecretModels.Host)
	aws.Ctx = context.WithValue(aws.Ctx, domain.Key("isSrv"), SecretModels.IsServ)
	aws.Ctx = context.WithValue(aws.Ctx, domain.Key("database"), SecretModels.Database)
	aws.Ctx = context.WithValue(aws.Ctx, domain.Key("jwtSign"), SecretModels.JWTSign)
	aws.Ctx = context.WithValue(aws.Ctx, domain.Key("bucket_name"), os.Getenv("BUCKET_NAME"))

	err = db.ConnectMongo(aws.Ctx)
	if err != nil {
		return utils.HandleError(500, "Error al conectar a la base de datos: "+err.Error()), nil
	}

	respAPI := handlers.AwsHandler(aws.Ctx, request)
	if respAPI.CustomResp == nil {
		return utils.FormatResponse(respAPI.Status, respAPI.Message, respAPI.Data, respAPI.Meta), nil
	} else {
		return respAPI.CustomResp, nil
	}
}

func validateParams() bool {
	_, isParam := os.LookupEnv("SECRET_NAME")
	if !isParam {
		return false
	}

	_, isParam = os.LookupEnv("BUCKET_NAME")
	if !isParam {
		return false
	}

	_, isParam = os.LookupEnv("URL_PREFIX")
	if !isParam {
		return false
	}

	return true
}
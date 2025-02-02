package services

import (
	"context"
	"os"
	"strings"
	"user-service/handlers"
	"user-service/internal/domain"
	"user-service/internal/infrastructure/db"
	"user-service/pkg/aws"
	secretmanager "user-service/pkg/secret-manager"

	"github.com/aws/aws-lambda-go/events"
)

func LambdaExec(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

	aws.StartAWS()   

	if !validateParams() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body: "Error en las variables de entorno. Deben incluir 'SECRET_NAME', 'BUCKET_NAME' y 'URL_PREFIX'",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	SecretModels, err := secretmanager.GetSecret(os.Getenv("SECRET_NAME"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body: "Error al obtener secret"+err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	path := strings.Replace(request.PathParameters["twitteruala"], os.Getenv("URL_PREFIX"), "", -1)
	path = strings.TrimPrefix(path, "/")

	if path == "" {
		path = strings.TrimPrefix(request.Path, "/development/")
		path = strings.TrimPrefix(path, "/")
	}

	if request.Body == "" {
    res := &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "El cuerpo de la solicitud no puede estar vacío",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
    }
    return res, nil
	}

	aws.Ctx = context.WithValue(aws.Ctx, domain.Key("path"), path)
	aws.Ctx = context.WithValue(aws.Ctx, domain.Key("method"), request.HTTPMethod)
	aws.Ctx = context.WithValue(aws.Ctx, domain.Key("body"), request.Body)
	aws.Ctx = context.WithValue(aws.Ctx, domain.Key("user"), SecretModels.Username)
	aws.Ctx = context.WithValue(aws.Ctx, domain.Key("password"), SecretModels.Password)
	aws.Ctx = context.WithValue(aws.Ctx, domain.Key("host"), SecretModels.Host)
	aws.Ctx = context.WithValue(aws.Ctx, domain.Key("database"), SecretModels.Database)
	aws.Ctx = context.WithValue(aws.Ctx, domain.Key("jwtSign"), SecretModels.JWTSign)
	aws.Ctx = context.WithValue(aws.Ctx, domain.Key("bucket_name"), os.Getenv("BUCKET_NAME"))

	err = db.ConnectMongo(aws.Ctx)
	if err != nil {
		res := &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body: "Error al conectar a la base de datos" + err.Error(),
			Headers: map[string]string {
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	respAPI := handlers.AwsHandler(aws.Ctx, request)
	if respAPI.CustomResp == nil {
		res := &events.APIGatewayProxyResponse{
			StatusCode: respAPI.Status,
			Body: respAPI.Message,
			Headers: map[string]string {
				"Content-Type": "application/json",
			},
		}
		return res, nil
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
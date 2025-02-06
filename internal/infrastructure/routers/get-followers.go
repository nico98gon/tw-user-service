package routers

import (
	"user-service/internal/domain"
	"user-service/internal/infrastructure/db"

	"github.com/aws/aws-lambda-go/events"
)

func GetFollowers(request events.APIGatewayProxyRequest) domain.RespAPI {
	var r domain.RespAPI
	r.Status = 400
	
	UserID := request.QueryStringParameters["id"]
	if len(UserID) == 0 || UserID == "000000000000000000000000" {
		r.Message = "ID de usuario inválido"
		return r
	}

	users, err := db.SearchFollowers(UserID)
	if err != nil {
			r.Message = "Error al obtener seguidores: " + err.Error()
			return r
	}

	r.Status = 200
	r.Message = "Lista de seguidores"
	r.Data = users
	return r
}
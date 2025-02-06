package routers

import (
	"user-service/internal/domain"
	"user-service/internal/infrastructure/db"

	"github.com/aws/aws-lambda-go/events"
)

func GetFollowing(request events.APIGatewayProxyRequest) domain.RespAPI {
	var r domain.RespAPI
	r.Status = 400
	
	UserID := request.QueryStringParameters["id"]
	if len(UserID) == 0 || UserID == "000000000000000000000000" {
		r.Message = "ID de usuario inválido"
		return r
	}

	users, err := db.SearchFollowing(UserID)
	if err != nil {
		r.Message = "Error al obtener seguidos: " + err.Error()
		return r
	}

	r.Status = 200
	r.Message = "Lista de seguidos"
	r.Data = users
	return r
}
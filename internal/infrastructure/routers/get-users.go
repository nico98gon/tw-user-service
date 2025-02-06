package routers

import (
	"user-service/internal/domain"
	"user-service/internal/infrastructure/db"

	"github.com/aws/aws-lambda-go/events"
)

func GetUsers(request events.APIGatewayProxyRequest, claim *domain.Claim) domain.RespAPI {
	var r domain.RespAPI
	r.Status = 400

	cursor := request.QueryStringParameters["cursor"]
	typeUser := request.QueryStringParameters["type"]
	search := request.QueryStringParameters["search"]
	IDUser := claim.ID.Hex()

	users, nextCursor, status := db.SearchAllUsers(IDUser, cursor, search, typeUser)
	if !status {
		r.Message = "Error al buscar los usuarios"
		return r
	}

	r.Status = 200
	r.Message = "Usuarios encontrados"
	r.Data = users
	r.Meta = map[string]interface{}{"nextCursor": nextCursor}

	return r
}
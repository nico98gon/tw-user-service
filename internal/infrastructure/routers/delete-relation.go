package routers

import (
	"fmt"
	"user-service/internal/domain"
	"user-service/internal/domain/relations"
	"user-service/internal/infrastructure/db"

	"github.com/aws/aws-lambda-go/events"
)

func DeleteRelation(request events.APIGatewayProxyRequest, claim *domain.Claim) domain.RespAPI {
	var r domain.RespAPI
	r.Status = 400

	IDRel := request.QueryStringParameters["id"]
	if len(IDRel) == 0 || IDRel == "000000000000000000000000" {
		r.Message = "ID de usuarioImageRelation es requerido: " + IDRel
		return r
	}

	UserID := claim.ID.Hex()
	if len(UserID) == 0 || UserID == "000000000000000000000000" {
		r.Message = "ID de usuario no válido en el token: " + UserID
		fmt.Println(r.Message)
		return r
	}

	var rel relations.Relation
	rel.UserID = UserID
	rel.UserIDRel = IDRel

	status, err := db.DeleteRelationFromDB(rel)
	if err != nil {
		r.Message = "Error al eliminar la relación: " + err.Error()
		return r
	}

	if !status {
		r.Message = "Error al eliminar la relación, status error"
		return r
	}

	r.Status = 200
	r.Message = "Se elimino la relación correctamente"

	return r
}
package routers

import (
	"fmt"
	"user-service/internal/domain"
	"user-service/internal/domain/relations"
	"user-service/internal/infrastructure/db"

	"github.com/aws/aws-lambda-go/events"
)

func GetRelation(request events.APIGatewayProxyRequest, claim *domain.Claim) domain.RespAPI {
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

	var resp domain.RespGetRel

	isRelation := db.SearchRelation(rel)

	if isRelation {
		resp.Status = true
	} else {
		resp.Status = false
	}

	r.Status = 200
	r.Message = "Relacion encontrada"
	r.Data = resp

	return r
}
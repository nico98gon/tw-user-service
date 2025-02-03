package routers

import (
	"encoding/json"
	"fmt"
	"user-service/internal/domain"
	"user-service/internal/infrastructure/db"

	"github.com/aws/aws-lambda-go/events"
)

func Profile(request events.APIGatewayProxyRequest, claim *domain.Claim) domain.RespAPI {
	var r domain.RespAPI
	r.Status = 400

	fmt.Println(" > Perfil de usuario")

	ID := claim.ID.Hex()
	if len(ID) == 0 {
		r.Message = "ID es requerido"
		return r
	}

	fmt.Println("Buscando perfil con ID:", ID)
	profile, err := db.SearchProfile(ID)
	if err != nil {
		r.Message = "Error al buscar el perfil: " + err.Error()
		return r
	}

	respJson, err := json.Marshal(profile)
	if err != nil {
		r.Status = 500
		r.Message = "Error al serializar el perfil: " + err.Error()
		return r
	}

	r.Status = 200
	r.CustomResp = &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(respJson),
	}
	return r
}

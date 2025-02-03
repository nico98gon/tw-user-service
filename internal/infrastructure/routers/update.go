package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"user-service/internal/domain"
	"user-service/internal/domain/users"
	"user-service/internal/infrastructure/db"
)

func UpdateProfile(ctx context.Context, claim *domain.Claim) domain.RespAPI {
	var r domain.RespAPI
	r.Status = 400

	var u users.User

	body, ok := ctx.Value(domain.Key("body")).(string)
	if !ok {
		r.Message = "Error: No se pudo obtener el body de la solicitud"
		fmt.Println(r.Message)
		return r
	}

	err := json.Unmarshal([]byte(body), &u)
	if err != nil {
		r.Message = "Error al parsear el body: " + err.Error()
		fmt.Println(r.Message, u)
		return r
	}

	if err := users.UpdateValidations(u); err != nil {
		r.Message = err.Error()
		fmt.Println("Error en las validaciones:", r.Message, u)
		return r
	}

	ID := claim.ID.Hex()
	if len(ID) == 0 || ID == "000000000000000000000000" {
		r.Message = "Error: ID de usuario no válido en el token"
		fmt.Println(r.Message)
		return r
	}

	status, err := db.UpdateRegister(u, ID)
	if err != nil {
		r.Message = "Error al intentar actualizar el registro de usuario: " + err.Error()
		fmt.Println(r.Message, claim, u)
		return r
	}

	if !status {
		r.Message = "Error con el estado de la actualización del registro"
		fmt.Println(r.Message)
		return r
	}

	r.Status = 200
	r.Message = "Perfil actualizado correctamente"
	r.Data = map[string]interface{}{
		"userID": ID,
	}

	fmt.Println(r.Message)
	return r
}

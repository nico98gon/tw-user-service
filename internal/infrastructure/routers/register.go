package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"user-service/internal/domain"
	"user-service/internal/domain/users"
	"user-service/internal/infrastructure/db"
)

func Register(ctx context.Context) domain.RespAPI {
	var u users.User
	var r domain.RespAPI
	r.Status = 400

	fmt.Println(" > Registro de usuario")

	body, ok := ctx.Value(domain.Key("body")).(string)
	if !ok {
		r.Message = "Error: No se pudo obtener el body de la solicitud"
		fmt.Println(r.Message)
		return r
	}

	err := json.Unmarshal([]byte(body), &u)
	if err != nil {
		r.Message = "Error al parsear el body: " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	if err := users.RegisterValidations(u); err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	_, found, _ := db.UserAlreadyExists(u.Email)
	if found {
		r.Message = "El usuario ya existe registrado con ese email"
		fmt.Println(r.Message)
		return r
	}

	u.CreatedAt = time.Now()

	_, status, err := db.InsertRegister(u)
	if err != nil {
		r.Message = "Error al intentar registrar el usuario: " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	if !status {
		r.Message = "Error al intentar registrar el usuario, estado false"
		fmt.Println(r.Message)
		return r
	}

	r.Status = 200
	r.Message = "El usuario se ha registrado correctamente"
	fmt.Println(r.Message)

	return r
}
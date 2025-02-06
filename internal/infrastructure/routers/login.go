package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"user-service/internal/domain"
	"user-service/internal/domain/users"
	"user-service/internal/infrastructure/db"
	jwt "user-service/pkg/JWT"
)

func Login(ctx context.Context) domain.RespAPI {
	var u users.User
	var r domain.RespAPI
	r.Status = 400

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

	if err := users.LoginValidations(u); err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	userData, found := db.TryLogin(u.Email, u.Password)
	if !found {
		r.Message = "Usuario y/o contraseña incorrectos"
		fmt.Println(r.Message)
		return r
	}

	jwtKey, err := jwt.GenerateToken(ctx, userData)
	if err != nil {
		r.Message = "Error al intentar generar el token: " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	r.Status = 200
	r.Message = "Login exitoso"
	r.Data = map[string]interface{}{
		"token": jwtKey,
	}

	return r
}
